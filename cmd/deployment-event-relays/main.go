package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/nais/liberator/pkg/conftools"
	"github.com/nais/liberator/pkg/tlsutil"
	"github.com/navikt/deployment-event-relays/pkg/config"
	"github.com/navikt/deployment-event-relays/pkg/deployment"
	"github.com/navikt/deployment-event-relays/pkg/influx"
	"github.com/navikt/deployment-event-relays/pkg/kafka/consumer"
	"github.com/navikt/deployment-event-relays/pkg/logging"
	"github.com/navikt/deployment-event-relays/pkg/nora"
	"github.com/navikt/deployment-event-relays/pkg/vera"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type Processor interface {
	Process(event *deployment.Event) (retry bool, err error)
}

func kafkaConfig(cfg *config.Config, subsystem string, callback consumer.Callback) (*consumer.Config, error) {
	tlsConfig, err := tlsutil.TLSConfigFromFiles(
		cfg.Kafka.TLS.CertificatePath,
		cfg.Kafka.TLS.PrivateKeyPath,
		cfg.Kafka.TLS.CAPath,
	)
	if err != nil {
		return nil, err
	}
	return &consumer.Config{
		Brokers:           cfg.Kafka.Brokers,
		Callback:          callback,
		GroupID:           cfg.Kafka.GroupIDPrefix + "/" + subsystem,
		MaxProcessingTime: time.Second * 3,
		Logger:            log.StandardLogger(),
		RetryInterval:     time.Second * 30,
		TlsConfig:         tlsConfig,
		Topic:             cfg.Kafka.Topic,
	}, nil
}

func main() {
	err := run()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.DefaultConfig()
	config.BindFlags(cfg)

	conftools.Initialize("DER")
	err := conftools.Load(cfg)
	if err != nil {
		return err
	}

	pflag.Parse()

	err = logging.Apply(log.StandardLogger(), cfg.Log.Verbosity, cfg.Log.Format)
	if err != nil {
		return err
	}

	log.Infof("deployment-event-relays starting up")
	log.Infof("--- configuration ---")

	disallowedKeys := []string{
		"influxdb.password",
	}
	for _, configLine := range conftools.Format(disallowedKeys) {
		log.Info(configLine)
	}

	log.Infof("--- end configuration ---")

	sarama.Logger = log.StandardLogger()

	subsystems := make(map[string]Processor)

	if len(cfg.InfluxDB.URL) > 0 {
		subsystems["influxdb"] = &influx.Relay{
			URL:      cfg.InfluxDB.URL,
			Username: cfg.InfluxDB.Username,
			Password: cfg.InfluxDB.Password,
		}
	}

	if len(cfg.Nora.URL) > 0 {
		subsystems["nora"] = &nora.Relay{
			URL: cfg.InfluxDB.URL,
		}
	}

	if len(cfg.Vera.URL) > 0 {
		subsystems["vera"] = &vera.Relay{
			URL: cfg.Vera.URL,
		}
	}

	if len(subsystems)==0 {
		return fmt.Errorf("no subsystems enabled")
	}

	for key, relayer := range subsystems {
		callback := func(message *sarama.ConsumerMessage, logger *log.Entry) (retry bool, err error) {
			event := &deployment.Event{}
			err = proto.Unmarshal(message.Value, event)
			if err != nil {
				return false, fmt.Errorf("%s: unmarshal incoming message: %w", key, err)
			}
			log.Tracef("subsystem '%s': incoming message: %#v", key, event)
			return relayer.Process(event)
		}
		kafkacfg, err := kafkaConfig(cfg, key, callback)
		if err != nil {
			return fmt.Errorf("set up subsystem '%s': %w", key, err)
		}
		_, err = consumer.New(*kafkacfg)
		if err != nil {
			return fmt.Errorf("initialize Kafka for subsystem '%s': %w", key, err)
		}
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	sig := <-sigs

	log.Infof("Received %s, shutting down.", sig)

	return nil
}
