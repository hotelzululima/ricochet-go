// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package ricochet is a generated protocol buffer package.

It is generated from these files:
	config.proto
	contact.proto
	conversation.proto
	core.proto
	identity.proto
	network.proto

It has these top-level messages:
	Config
	ConfigRequest
	Contact
	ContactRequest
	MonitorContactsRequest
	ContactEvent
	AddContactReply
	DeleteContactRequest
	DeleteContactReply
	RejectInboundRequestReply
	ConversationEvent
	Entity
	Message
	ServerStatusRequest
	ServerStatusReply
	Identity
	IdentityRequest
	MonitorNetworkRequest
	TorProcessStatus
	TorControlStatus
	TorConnectionStatus
	NetworkStatus
	StartNetworkRequest
	StopNetworkRequest
*/
package ricochet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ConfigRequest struct {
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Config)(nil), "ricochet.Config")
	proto.RegisterType((*ConfigRequest)(nil), "ricochet.ConfigRequest")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0xca, 0x4c, 0xce, 0x4f, 0xce,
	0x48, 0x2d, 0x51, 0xe2, 0xe0, 0x62, 0x73, 0x06, 0xcb, 0x28, 0xf1, 0x73, 0xf1, 0x42, 0x58, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x49, 0x6c, 0x60, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xad, 0xf5, 0x95, 0x4e, 0x3b, 0x00, 0x00, 0x00,
}