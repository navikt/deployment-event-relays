// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// PlatformType denotes what kind of platform are used for deploying.
type PlatformType int32

const (
	PlatformType_jboss PlatformType = 0
	PlatformType_was   PlatformType = 1
	PlatformType_bpm   PlatformType = 2
	PlatformType_nais  PlatformType = 3
)

var PlatformType_name = map[int32]string{
	0: "jboss",
	1: "was",
	2: "bpm",
	3: "nais",
}

var PlatformType_value = map[string]int32{
	"jboss": 0,
	"was":   1,
	"bpm":   2,
	"nais":  3,
}

func (x PlatformType) String() string {
	return proto.EnumName(PlatformType_name, int32(x))
}

func (PlatformType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

// Only enumerated systems can report deployment status.
type System int32

const (
	System_aura       System = 0
	System_naisd      System = 1
	System_naiserator System = 2
)

var System_name = map[int32]string{
	0: "aura",
	1: "naisd",
	2: "naiserator",
}

var System_value = map[string]int32{
	"aura":       0,
	"naisd":      1,
	"naiserator": 2,
}

func (x System) String() string {
	return proto.EnumName(System_name, int32(x))
}

func (System) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

// RolloutStatus denotes whether a deployment has been initialized,
// rolled out successfully, or if the status is altogether unknown.
type RolloutStatus int32

const (
	RolloutStatus_unknown     RolloutStatus = 0
	RolloutStatus_initialized RolloutStatus = 1
	RolloutStatus_complete    RolloutStatus = 2
)

var RolloutStatus_name = map[int32]string{
	0: "unknown",
	1: "initialized",
	2: "complete",
}

var RolloutStatus_value = map[string]int32{
	"unknown":     0,
	"initialized": 1,
	"complete":    2,
}

func (x RolloutStatus) String() string {
	return proto.EnumName(RolloutStatus_name, int32(x))
}

func (RolloutStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

// Environment separates between production and development environments.
type Environment int32

const (
	Environment_production  Environment = 0
	Environment_development Environment = 1
)

var Environment_name = map[int32]string{
	0: "production",
	1: "development",
}

var Environment_value = map[string]int32{
	"production":  0,
	"development": 1,
}

func (x Environment) String() string {
	return proto.EnumName(Environment_name, int32(x))
}

func (Environment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}

// A platform represents a place where applications and systems are deployed.
// Since platforms come in different versions and flavors, a variant can also be specified.
type Platform struct {
	Type                 PlatformType `protobuf:"varint,1,opt,name=type,proto3,enum=deployment.PlatformType" json:"type,omitempty"`
	Variant              string       `protobuf:"bytes,2,opt,name=variant,proto3" json:"variant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Platform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform.Unmarshal(m, b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
}
func (m *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(m, src)
}
func (m *Platform) XXX_Size() int {
	return xxx_messageInfo_Platform.Size(m)
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func (m *Platform) GetType() PlatformType {
	if m != nil {
		return m.Type
	}
	return PlatformType_jboss
}

func (m *Platform) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

// Actor is a human being or a service account.
type Actor struct {
	Ident                string   `protobuf:"bytes,1,opt,name=ident,proto3" json:"ident,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Actor) Reset()         { *m = Actor{} }
func (m *Actor) String() string { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()    {}
func (*Actor) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *Actor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Actor.Unmarshal(m, b)
}
func (m *Actor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Actor.Marshal(b, m, deterministic)
}
func (m *Actor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Actor.Merge(m, src)
}
func (m *Actor) XXX_Size() int {
	return xxx_messageInfo_Actor.Size(m)
}
func (m *Actor) XXX_DiscardUnknown() {
	xxx_messageInfo_Actor.DiscardUnknown(m)
}

var xxx_messageInfo_Actor proto.InternalMessageInfo

func (m *Actor) GetIdent() string {
	if m != nil {
		return m.Ident
	}
	return ""
}

func (m *Actor) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Actor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ContainerImage is a reference to a image that can be deployed as a container,
// typically a Docker container inside a Kubernetes pod.
type ContainerImage struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerImage) Reset()         { *m = ContainerImage{} }
func (m *ContainerImage) String() string { return proto.CompactTextString(m) }
func (*ContainerImage) ProtoMessage()    {}
func (*ContainerImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

func (m *ContainerImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerImage.Unmarshal(m, b)
}
func (m *ContainerImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerImage.Marshal(b, m, deterministic)
}
func (m *ContainerImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerImage.Merge(m, src)
}
func (m *ContainerImage) XXX_Size() int {
	return xxx_messageInfo_ContainerImage.Size(m)
}
func (m *ContainerImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerImage.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerImage proto.InternalMessageInfo

func (m *ContainerImage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerImage) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ContainerImage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// Event represents a deployment that has been made on any of NAV's systems.
type Event struct {
	// CorrelationID can be used to correlate events across different systems.
	CorrelationID string `protobuf:"bytes,1,opt,name=correlationID,proto3" json:"correlationID,omitempty"`
	// Platform represents the technical platform on which the deployment was made.
	Platform *Platform `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	// Source tells which system that reported the deployment.
	Source System `protobuf:"varint,3,opt,name=source,proto3,enum=deployment.System" json:"source,omitempty"`
	// Deployer is a reference to a human being that started the deployment.
	Deployer *Actor `protobuf:"bytes,4,opt,name=deployer,proto3" json:"deployer,omitempty"`
	// Team is an organizational structure within NAV and refers to a group of people.
	Team string `protobuf:"bytes,5,opt,name=team,proto3" json:"team,omitempty"`
	// RolloutStatus shows the deployment status.
	RolloutStatus RolloutStatus `protobuf:"varint,6,opt,name=rolloutStatus,proto3,enum=deployment.RolloutStatus" json:"rolloutStatus,omitempty"`
	// Environment can be production or development.
	Environment Environment `protobuf:"varint,7,opt,name=environment,proto3,enum=deployment.Environment" json:"environment,omitempty"`
	// The SKYA platform divides between production, development, staging, and test.
	// Furthermore, these environments are divided into smaller segments denoted with
	// a number, such as q0, t6, u11.
	SkyaEnvironment string `protobuf:"bytes,8,opt,name=skyaEnvironment,proto3" json:"skyaEnvironment,omitempty"`
	// Namespace represents the Kubernetes namespace this deployment was made into.
	Namespace string `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Cluster is the name of the Kubernetes cluster that was deployed to.
	Cluster string `protobuf:"bytes,10,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Application is the name of the deployed application.
	Application string `protobuf:"bytes,11,opt,name=application,proto3" json:"application,omitempty"`
	// Version is the version of the deployed application.
	Version string `protobuf:"bytes,12,opt,name=version,proto3" json:"version,omitempty"`
	// Image refers to the container source, usually a Docker image.
	Image *ContainerImage `protobuf:"bytes,13,opt,name=image,proto3" json:"image,omitempty"`
	// Timestamp is the generation time of the deployment event.
	Timestamp            int64    `protobuf:"varint,14,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetCorrelationID() string {
	if m != nil {
		return m.CorrelationID
	}
	return ""
}

func (m *Event) GetPlatform() *Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

func (m *Event) GetSource() System {
	if m != nil {
		return m.Source
	}
	return System_aura
}

func (m *Event) GetDeployer() *Actor {
	if m != nil {
		return m.Deployer
	}
	return nil
}

func (m *Event) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *Event) GetRolloutStatus() RolloutStatus {
	if m != nil {
		return m.RolloutStatus
	}
	return RolloutStatus_unknown
}

func (m *Event) GetEnvironment() Environment {
	if m != nil {
		return m.Environment
	}
	return Environment_production
}

func (m *Event) GetSkyaEnvironment() string {
	if m != nil {
		return m.SkyaEnvironment
	}
	return ""
}

func (m *Event) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Event) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Event) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *Event) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Event) GetImage() *ContainerImage {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("deployment.PlatformType", PlatformType_name, PlatformType_value)
	proto.RegisterEnum("deployment.System", System_name, System_value)
	proto.RegisterEnum("deployment.RolloutStatus", RolloutStatus_name, RolloutStatus_value)
	proto.RegisterEnum("deployment.Environment", Environment_name, Environment_value)
	proto.RegisterType((*Platform)(nil), "deployment.Platform")
	proto.RegisterType((*Actor)(nil), "deployment.Actor")
	proto.RegisterType((*ContainerImage)(nil), "deployment.ContainerImage")
	proto.RegisterType((*Event)(nil), "deployment.Event")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xbb, 0xf9, 0xce, 0x6c, 0x93, 0x2e, 0xa3, 0x4a, 0x98, 0x8a, 0x43, 0x14, 0x71, 0x88,
	0x22, 0x1a, 0x55, 0x45, 0x1c, 0x10, 0x07, 0xc4, 0x47, 0x85, 0xca, 0x09, 0x6d, 0x79, 0x01, 0x77,
	0x33, 0xb4, 0xa6, 0xbb, 0xb6, 0x65, 0x7b, 0x53, 0x85, 0x27, 0xe3, 0xf1, 0x90, 0xbd, 0xdb, 0xc4,
	0x41, 0xdc, 0x66, 0xfe, 0xf3, 0x9b, 0xc9, 0xce, 0xf8, 0xaf, 0x40, 0x4a, 0x1b, 0x92, 0x6e, 0xa5,
	0x8d, 0x72, 0x0a, 0x61, 0x4d, 0xba, 0x54, 0xdb, 0x8a, 0xa4, 0x9b, 0xe7, 0x30, 0xfa, 0x5e, 0x72,
	0xf7, 0x53, 0x99, 0x0a, 0x5f, 0x43, 0xcf, 0x6d, 0x35, 0xb1, 0x64, 0x96, 0x2c, 0xa6, 0x97, 0x6c,
	0xb5, 0xc7, 0x56, 0x4f, 0xcc, 0x8f, 0xad, 0xa6, 0x3c, 0x50, 0xc8, 0x60, 0xb8, 0xe1, 0x46, 0x70,
	0xe9, 0x58, 0x67, 0x96, 0x2c, 0xc6, 0xf9, 0x53, 0x3a, 0xff, 0x0a, 0xfd, 0x8f, 0x85, 0x53, 0x06,
	0x4f, 0xa1, 0x2f, 0xd6, 0x24, 0x5d, 0x98, 0x38, 0xce, 0x9b, 0xc4, 0xab, 0x54, 0x71, 0x51, 0xb6,
	0x6d, 0x4d, 0x82, 0x08, 0x3d, 0xc9, 0x2b, 0x62, 0xdd, 0x20, 0x86, 0x78, 0xfe, 0x0d, 0xa6, 0x9f,
	0x95, 0x74, 0x5c, 0x48, 0x32, 0xd7, 0x15, 0xbf, 0xa3, 0x1d, 0x95, 0xec, 0x29, 0xcc, 0xa0, 0xeb,
	0xf8, 0x5d, 0x3b, 0xcd, 0x87, 0x9e, 0xba, 0xe7, 0xf6, 0xfe, 0x69, 0x96, 0x8f, 0xe7, 0x7f, 0x7a,
	0xd0, 0xbf, 0xf2, 0x47, 0xc0, 0x57, 0x30, 0x29, 0x94, 0x31, 0x54, 0x72, 0x27, 0x94, 0xbc, 0xfe,
	0xd2, 0x0e, 0x3b, 0x14, 0xf1, 0x02, 0x46, 0xba, 0x5d, 0x3a, 0x8c, 0x4e, 0x2f, 0x4f, 0xff, 0x77,
	0x90, 0x7c, 0x47, 0xe1, 0x12, 0x06, 0x56, 0xd5, 0xa6, 0x68, 0x76, 0x98, 0x5e, 0x62, 0xcc, 0xdf,
	0x6c, 0xad, 0xa3, 0x2a, 0x6f, 0x09, 0x3c, 0x87, 0x51, 0x53, 0x24, 0xc3, 0x7a, 0x61, 0xfa, 0xb3,
	0x98, 0x0e, 0xe7, 0xcb, 0x77, 0x88, 0x5f, 0xc8, 0x11, 0xaf, 0x58, 0xbf, 0x59, 0xc8, 0xc7, 0xf8,
	0x01, 0x26, 0x46, 0x95, 0xa5, 0xaa, 0xdd, 0x8d, 0xe3, 0xae, 0xb6, 0x6c, 0x10, 0x7e, 0xf5, 0x45,
	0x3c, 0x27, 0x8f, 0x81, 0xfc, 0x90, 0xc7, 0x77, 0x90, 0x92, 0xdc, 0x08, 0xa3, 0xa4, 0x67, 0xd9,
	0x30, 0xb4, 0x3f, 0x8f, 0xdb, 0xaf, 0xf6, 0xe5, 0x3c, 0x66, 0x71, 0x01, 0x27, 0xf6, 0x61, 0xcb,
	0xa3, 0x3a, 0x1b, 0x85, 0x4f, 0xfb, 0x57, 0xc6, 0x97, 0x30, 0xf6, 0x8f, 0x64, 0x35, 0x2f, 0x88,
	0x8d, 0x03, 0xb3, 0x17, 0xbc, 0x87, 0x8a, 0xb2, 0xb6, 0x8e, 0x0c, 0x83, 0xc6, 0x43, 0x6d, 0x8a,
	0x33, 0x48, 0xb9, 0xd6, 0xa5, 0x28, 0xc2, 0x7b, 0xb0, 0x34, 0x54, 0x63, 0x29, 0xf8, 0x8f, 0x8c,
	0xf5, 0xd5, 0xe3, 0xd6, 0x7f, 0x4d, 0x8a, 0x17, 0xd0, 0x17, 0xde, 0x2d, 0x6c, 0x12, 0x2e, 0x7b,
	0x16, 0xaf, 0x74, 0xe8, 0xa7, 0xbc, 0x01, 0xfd, 0x57, 0x3a, 0x51, 0x91, 0x75, 0xbc, 0xd2, 0x6c,
	0x3a, 0x4b, 0x16, 0xdd, 0x7c, 0x2f, 0x2c, 0xdf, 0xc2, 0x71, 0xec, 0x7f, 0x1c, 0x43, 0xff, 0xd7,
	0xad, 0xb2, 0x36, 0x3b, 0xc2, 0x21, 0x74, 0x1f, 0xb9, 0xcd, 0x12, 0x1f, 0xdc, 0xea, 0x2a, 0xeb,
	0xe0, 0xc8, 0x3b, 0x54, 0xd8, 0xac, 0xbb, 0x3c, 0x87, 0x41, 0xf3, 0xea, 0x5e, 0xe3, 0xb5, 0xe1,
	0xd9, 0x91, 0x6f, 0xf5, 0xd5, 0x75, 0x96, 0xe0, 0x14, 0xc0, 0x87, 0x64, 0xb8, 0x53, 0x26, 0xeb,
	0x2c, 0xdf, 0xc3, 0xe4, 0xe0, 0xb9, 0x30, 0x85, 0x61, 0x2d, 0x1f, 0xa4, 0x7a, 0x94, 0xd9, 0x11,
	0x9e, 0x40, 0x2a, 0xa4, 0x70, 0x82, 0x97, 0xe2, 0x37, 0xf9, 0xf6, 0x63, 0x18, 0x15, 0xaa, 0xd2,
	0x25, 0x39, 0xca, 0x3a, 0xcb, 0x15, 0xa4, 0xf1, 0xd5, 0xa7, 0x00, 0xda, 0xa8, 0x75, 0x5d, 0xf8,
	0x4b, 0x35, 0xdd, 0x6b, 0xda, 0x50, 0xa9, 0xb4, 0x2f, 0x67, 0xc9, 0xa7, 0x33, 0x60, 0x52, 0xad,
	0x24, 0xdf, 0x34, 0x7f, 0x09, 0x36, 0x3a, 0xd1, 0xed, 0x20, 0x48, 0x6f, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xdd, 0x26, 0xde, 0xad, 0x34, 0x04, 0x00, 0x00,
}