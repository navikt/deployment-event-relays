package vera

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"

	"github.com/navikt/deployment-event-relays/pkg/deployment"
)

// Payload represents the JSON payload supported by the Vera API. All fields are required
type Payload struct {
	Environment      string `json:"environment"`
	Application      string `json:"application"`
	Version          string `json:"version"`
	Deployer         string `json:"deployedBy"`
	Environmentclass string `json:"environmentclass"`
}

// BuildVeraEvent collects data from a deployment event and creates a valid payload for POSTing to the vera api.
func BuildVeraEvent(event *deployment.Event) Payload {
	var deployer string

	if len(event.GetDeployer().GetName()) > 0 {
		deployer = fmt.Sprintf("%s (%s)", event.GetSource().String(), event.GetDeployer().GetName())
	} else if len(event.GetDeployer().GetIdent()) > 0 {
		deployer = fmt.Sprintf("%s (%s)", event.GetSource().String(), event.GetDeployer().GetIdent())
	} else {
		deployer = event.GetSource().String()
	}

	return Payload{
		Environment:      getEnvironment(event),
		Application:      event.GetApplication(),
		Version:          event.GetVersion(),
		Deployer:         deployer,
		Environmentclass: event.GetEnvironment().String(),
	}
}

func getEnvironment(event *deployment.Event) string {
	if event.GetSkyaEnvironment() != "" {
		return event.GetSkyaEnvironment()
	}
	return fmt.Sprintf("%s (%s)", event.GetCluster(), event.GetNamespace())
}

// Marshal VeraPayload struct to JSON
func (payload Payload) Marshal() ([]byte, error) {
	var marshaledPayload []byte
	var err error

	marshaledPayload, err = json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return marshaledPayload, nil
}
