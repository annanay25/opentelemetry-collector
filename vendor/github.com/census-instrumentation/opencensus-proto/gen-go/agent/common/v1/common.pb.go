// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencensus/proto/agent/common/v1/common.proto

package v1 // import "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LibraryInfo_Language int32

const (
	LibraryInfo_LANGUAGE_UNSPECIFIED LibraryInfo_Language = 0
	LibraryInfo_CPP                  LibraryInfo_Language = 1
	LibraryInfo_C_SHARP              LibraryInfo_Language = 2
	LibraryInfo_ERLANG               LibraryInfo_Language = 3
	LibraryInfo_GO_LANG              LibraryInfo_Language = 4
	LibraryInfo_JAVA                 LibraryInfo_Language = 5
	LibraryInfo_NODE_JS              LibraryInfo_Language = 6
	LibraryInfo_PHP                  LibraryInfo_Language = 7
	LibraryInfo_PYTHON               LibraryInfo_Language = 8
	LibraryInfo_RUBY                 LibraryInfo_Language = 9
)

var LibraryInfo_Language_name = map[int32]string{
	0: "LANGUAGE_UNSPECIFIED",
	1: "CPP",
	2: "C_SHARP",
	3: "ERLANG",
	4: "GO_LANG",
	5: "JAVA",
	6: "NODE_JS",
	7: "PHP",
	8: "PYTHON",
	9: "RUBY",
}
var LibraryInfo_Language_value = map[string]int32{
	"LANGUAGE_UNSPECIFIED": 0,
	"CPP":                  1,
	"C_SHARP":              2,
	"ERLANG":               3,
	"GO_LANG":              4,
	"JAVA":                 5,
	"NODE_JS":              6,
	"PHP":                  7,
	"PYTHON":               8,
	"RUBY":                 9,
}

func (x LibraryInfo_Language) String() string {
	return proto.EnumName(LibraryInfo_Language_name, int32(x))
}
func (LibraryInfo_Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_16e6d73ff146d81e, []int{2, 0}
}

// Identifier metadata of the Node that connects to OpenCensus Agent.
// In the future we plan to extend the identifier proto definition to support
// additional information (e.g cloud id, monitored resource, etc.)
type Node struct {
	// Identifier that uniquely identifies a process within a VM/container.
	Identifier *ProcessIdentifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Information on the OpenCensus Library who initiates the stream.
	LibraryInfo *LibraryInfo `protobuf:"bytes,2,opt,name=library_info,json=libraryInfo,proto3" json:"library_info,omitempty"`
	// Additional informantion on service.
	ServiceInfo *ServiceInfo `protobuf:"bytes,3,opt,name=service_info,json=serviceInfo,proto3" json:"service_info,omitempty"`
	// Additional attributes.
	Attributes           map[string]string `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_16e6d73ff146d81e, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetIdentifier() *ProcessIdentifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Node) GetLibraryInfo() *LibraryInfo {
	if m != nil {
		return m.LibraryInfo
	}
	return nil
}

func (m *Node) GetServiceInfo() *ServiceInfo {
	if m != nil {
		return m.ServiceInfo
	}
	return nil
}

func (m *Node) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Identifier that uniquely identifies a process within a VM/container.
type ProcessIdentifier struct {
	// The host name. Usually refers to the machine/container name.
	// For example: os.Hostname() in Go, socket.gethostname() in Python.
	HostName string `protobuf:"bytes,1,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	// Process id.
	Pid uint32 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	// Start time of this ProcessIdentifier. Represented in epoch time.
	StartTimestamp       *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProcessIdentifier) Reset()         { *m = ProcessIdentifier{} }
func (m *ProcessIdentifier) String() string { return proto.CompactTextString(m) }
func (*ProcessIdentifier) ProtoMessage()    {}
func (*ProcessIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_16e6d73ff146d81e, []int{1}
}
func (m *ProcessIdentifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessIdentifier.Unmarshal(m, b)
}
func (m *ProcessIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessIdentifier.Marshal(b, m, deterministic)
}
func (dst *ProcessIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessIdentifier.Merge(dst, src)
}
func (m *ProcessIdentifier) XXX_Size() int {
	return xxx_messageInfo_ProcessIdentifier.Size(m)
}
func (m *ProcessIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessIdentifier proto.InternalMessageInfo

func (m *ProcessIdentifier) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *ProcessIdentifier) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ProcessIdentifier) GetStartTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

// Information on OpenCensus Library.
type LibraryInfo struct {
	// Language of OpenCensus Library.
	Language LibraryInfo_Language `protobuf:"varint,1,opt,name=language,proto3,enum=opencensus.proto.agent.common.v1.LibraryInfo_Language" json:"language,omitempty"`
	// Version of Agent exporter of Library.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LibraryInfo) Reset()         { *m = LibraryInfo{} }
func (m *LibraryInfo) String() string { return proto.CompactTextString(m) }
func (*LibraryInfo) ProtoMessage()    {}
func (*LibraryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_16e6d73ff146d81e, []int{2}
}
func (m *LibraryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LibraryInfo.Unmarshal(m, b)
}
func (m *LibraryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LibraryInfo.Marshal(b, m, deterministic)
}
func (dst *LibraryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LibraryInfo.Merge(dst, src)
}
func (m *LibraryInfo) XXX_Size() int {
	return xxx_messageInfo_LibraryInfo.Size(m)
}
func (m *LibraryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LibraryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LibraryInfo proto.InternalMessageInfo

func (m *LibraryInfo) GetLanguage() LibraryInfo_Language {
	if m != nil {
		return m.Language
	}
	return LibraryInfo_LANGUAGE_UNSPECIFIED
}

func (m *LibraryInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Additional service information.
type ServiceInfo struct {
	// Name of the service.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInfo) Reset()         { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()    {}
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_16e6d73ff146d81e, []int{3}
}
func (m *ServiceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfo.Unmarshal(m, b)
}
func (m *ServiceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfo.Marshal(b, m, deterministic)
}
func (dst *ServiceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfo.Merge(dst, src)
}
func (m *ServiceInfo) XXX_Size() int {
	return xxx_messageInfo_ServiceInfo.Size(m)
}
func (m *ServiceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfo proto.InternalMessageInfo

func (m *ServiceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Node)(nil), "opencensus.proto.agent.common.v1.Node")
	proto.RegisterMapType((map[string]string)(nil), "opencensus.proto.agent.common.v1.Node.AttributesEntry")
	proto.RegisterType((*ProcessIdentifier)(nil), "opencensus.proto.agent.common.v1.ProcessIdentifier")
	proto.RegisterType((*LibraryInfo)(nil), "opencensus.proto.agent.common.v1.LibraryInfo")
	proto.RegisterType((*ServiceInfo)(nil), "opencensus.proto.agent.common.v1.ServiceInfo")
	proto.RegisterEnum("opencensus.proto.agent.common.v1.LibraryInfo_Language", LibraryInfo_Language_name, LibraryInfo_Language_value)
}

func init() {
	proto.RegisterFile("opencensus/proto/agent/common/v1/common.proto", fileDescriptor_common_16e6d73ff146d81e)
}

var fileDescriptor_common_16e6d73ff146d81e = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6f, 0xda, 0x30,
	0x14, 0xc6, 0x17, 0x42, 0x0b, 0xbc, 0x6c, 0x6d, 0x66, 0xf5, 0x80, 0xba, 0xc3, 0x3a, 0x76, 0xe9,
	0x85, 0x44, 0x6d, 0xa5, 0x69, 0x9a, 0xb4, 0x43, 0x4a, 0xb3, 0x96, 0x0a, 0xa5, 0x91, 0x29, 0x95,
	0xba, 0x4b, 0x14, 0xc0, 0xa4, 0xd6, 0x88, 0x8d, 0x62, 0x07, 0x89, 0xd3, 0x8e, 0xd5, 0xfe, 0x81,
	0xfd, 0xbd, 0x93, 0xe3, 0x00, 0xd1, 0x7a, 0xa0, 0xb7, 0xf7, 0xfc, 0xbe, 0xef, 0x97, 0xe8, 0xf3,
	0x93, 0xa1, 0xcb, 0x17, 0x84, 0x4d, 0x08, 0x13, 0xb9, 0x70, 0x17, 0x19, 0x97, 0xdc, 0x8d, 0x13,
	0xc2, 0xa4, 0x3b, 0xe1, 0x69, 0xca, 0x99, 0xbb, 0x3c, 0x2b, 0x2b, 0xa7, 0x18, 0xa2, 0x93, 0xad,
	0x5c, 0x9f, 0x38, 0x85, 0xdc, 0x29, 0x45, 0xcb, 0xb3, 0xe3, 0x8f, 0x09, 0xe7, 0xc9, 0x9c, 0x68,
	0xd8, 0x38, 0x9f, 0xb9, 0x92, 0xa6, 0x44, 0xc8, 0x38, 0x5d, 0x68, 0x43, 0xe7, 0xaf, 0x09, 0xf5,
	0x80, 0x4f, 0x09, 0x1a, 0x02, 0xd0, 0x29, 0x61, 0x92, 0xce, 0x28, 0xc9, 0xda, 0xc6, 0x89, 0x71,
	0x6a, 0x9d, 0x5f, 0x38, 0xbb, 0x3e, 0xe0, 0x84, 0x19, 0x9f, 0x10, 0x21, 0xfa, 0x1b, 0x2b, 0xae,
	0x60, 0x50, 0x08, 0x6f, 0xe7, 0x74, 0x9c, 0xc5, 0xd9, 0x2a, 0xa2, 0x6c, 0xc6, 0xdb, 0xb5, 0x02,
	0xdb, 0xdd, 0x8d, 0x1d, 0x68, 0x57, 0x9f, 0xcd, 0x38, 0xb6, 0xe6, 0xdb, 0x46, 0x11, 0x05, 0xc9,
	0x96, 0x74, 0x42, 0x34, 0xd1, 0x7c, 0x2d, 0x71, 0xa8, 0x5d, 0x9a, 0x28, 0xb6, 0x0d, 0x7a, 0x00,
	0x88, 0xa5, 0xcc, 0xe8, 0x38, 0x97, 0x44, 0xb4, 0xeb, 0x27, 0xe6, 0xa9, 0x75, 0xfe, 0x65, 0x37,
	0x4f, 0x85, 0xe6, 0x78, 0x1b, 0xa3, 0xcf, 0x64, 0xb6, 0xc2, 0x15, 0xd2, 0xf1, 0x77, 0x38, 0xfc,
	0x6f, 0x8c, 0x6c, 0x30, 0x7f, 0x91, 0x55, 0x11, 0x6e, 0x0b, 0xab, 0x12, 0x1d, 0xc1, 0xde, 0x32,
	0x9e, 0xe7, 0xa4, 0x48, 0xa6, 0x85, 0x75, 0xf3, 0xad, 0xf6, 0xd5, 0xe8, 0x3c, 0x1b, 0xf0, 0xfe,
	0x45, 0xb8, 0xe8, 0x03, 0xb4, 0x9e, 0xb8, 0x90, 0x11, 0x8b, 0x53, 0x52, 0x72, 0x9a, 0xea, 0x20,
	0x88, 0x53, 0xa2, 0xf0, 0x0b, 0x3a, 0x2d, 0x50, 0xef, 0xb0, 0x2a, 0x51, 0x0f, 0x0e, 0x85, 0x8c,
	0x33, 0x19, 0x6d, 0xae, 0xbd, 0x0c, 0xec, 0xd8, 0xd1, 0x8b, 0xe1, 0xac, 0x17, 0xc3, 0xb9, 0x5f,
	0x2b, 0xf0, 0x41, 0x61, 0xd9, 0xf4, 0x9d, 0xe7, 0x1a, 0x58, 0x95, 0xfb, 0x40, 0x18, 0x9a, 0xf3,
	0x98, 0x25, 0x79, 0x9c, 0xe8, 0x5f, 0x38, 0x78, 0x4d, 0x5c, 0x15, 0x80, 0x33, 0x28, 0xdd, 0x78,
	0xc3, 0x41, 0x6d, 0x68, 0x2c, 0x49, 0x26, 0x28, 0x67, 0x65, 0x12, 0xeb, 0xb6, 0xf3, 0xc7, 0x80,
	0xe6, 0x60, 0x2b, 0x3b, 0x1a, 0x78, 0xc1, 0xf5, 0xc8, 0xbb, 0xf6, 0xa3, 0x51, 0x30, 0x0c, 0xfd,
	0x5e, 0xff, 0x47, 0xdf, 0xbf, 0xb2, 0xdf, 0xa0, 0x06, 0x98, 0xbd, 0x30, 0xb4, 0x0d, 0x64, 0x41,
	0xa3, 0x17, 0x0d, 0x6f, 0x3c, 0x1c, 0xda, 0x35, 0x04, 0xb0, 0xef, 0x63, 0xe5, 0xb0, 0x4d, 0x35,
	0xb8, 0xbe, 0x8b, 0x8a, 0xa6, 0x8e, 0x9a, 0x50, 0xbf, 0xf5, 0x1e, 0x3c, 0x7b, 0x4f, 0x1d, 0x07,
	0x77, 0x57, 0x7e, 0x74, 0x3b, 0xb4, 0xf7, 0x15, 0x25, 0xbc, 0x09, 0xed, 0x86, 0x32, 0x86, 0x8f,
	0xf7, 0x37, 0x77, 0x81, 0xdd, 0x54, 0x5a, 0x3c, 0xba, 0x7c, 0xb4, 0x5b, 0x9d, 0x4f, 0x60, 0x55,
	0xd6, 0x08, 0x21, 0xa8, 0x57, 0xee, 0xa1, 0xa8, 0x2f, 0x7f, 0xc3, 0x67, 0xca, 0x77, 0xc6, 0x71,
	0x69, 0xf5, 0x8a, 0x32, 0x54, 0xc3, 0xd0, 0xf8, 0xd9, 0x4f, 0xa8, 0x7c, 0xca, 0xc7, 0x4a, 0xe0,
	0x6a, 0x5f, 0x97, 0x32, 0x21, 0xb3, 0x3c, 0x25, 0x4c, 0xc6, 0x92, 0x72, 0xe6, 0x6e, 0x91, 0x5d,
	0xfd, 0x32, 0x24, 0x84, 0x75, 0x93, 0x17, 0x0f, 0xc4, 0x78, 0xbf, 0x98, 0x5e, 0xfc, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0x1d, 0xe4, 0xde, 0x4b, 0x04, 0x00, 0x00,
}