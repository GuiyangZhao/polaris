// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MatchString_MatchStringType int32

const (
	// Equivalent match
	MatchString_EXACT MatchString_MatchStringType = 0
	// Regular match
	MatchString_REGEX MatchString_MatchStringType = 1
	// Not equals match
	MatchString_NOT_EQUALS MatchString_MatchStringType = 2
	// Include match
	MatchString_IN MatchString_MatchStringType = 3
	// Not include match
	MatchString_NOT_IN MatchString_MatchStringType = 4
)

var MatchString_MatchStringType_name = map[int32]string{
	0: "EXACT",
	1: "REGEX",
	2: "NOT_EQUALS",
	3: "IN",
	4: "NOT_IN",
}
var MatchString_MatchStringType_value = map[string]int32{
	"EXACT":      0,
	"REGEX":      1,
	"NOT_EQUALS": 2,
	"IN":         3,
	"NOT_IN":     4,
}

func (x MatchString_MatchStringType) String() string {
	return proto.EnumName(MatchString_MatchStringType_name, int32(x))
}
func (MatchString_MatchStringType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_99674f3804a5acf0, []int{1, 0}
}

type MatchString_ValueType int32

const (
	MatchString_TEXT      MatchString_ValueType = 0
	MatchString_PARAMETER MatchString_ValueType = 1
	MatchString_VARIABLE  MatchString_ValueType = 2
)

var MatchString_ValueType_name = map[int32]string{
	0: "TEXT",
	1: "PARAMETER",
	2: "VARIABLE",
}
var MatchString_ValueType_value = map[string]int32{
	"TEXT":      0,
	"PARAMETER": 1,
	"VARIABLE":  2,
}

func (x MatchString_ValueType) String() string {
	return proto.EnumName(MatchString_ValueType_name, int32(x))
}
func (MatchString_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_model_99674f3804a5acf0, []int{1, 1}
}

type Location struct {
	Region               *wrappers.StringValue `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Zone                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Campus               *wrappers.StringValue `protobuf:"bytes,3,opt,name=campus,proto3" json:"campus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_99674f3804a5acf0, []int{0}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (dst *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(dst, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetRegion() *wrappers.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Location) GetZone() *wrappers.StringValue {
	if m != nil {
		return m.Zone
	}
	return nil
}

func (m *Location) GetCampus() *wrappers.StringValue {
	if m != nil {
		return m.Campus
	}
	return nil
}

type MatchString struct {
	Type                 MatchString_MatchStringType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.MatchString_MatchStringType" json:"type,omitempty"`
	Value                *wrappers.StringValue       `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ValueType            MatchString_ValueType       `protobuf:"varint,3,opt,name=value_type,json=valueType,proto3,enum=v1.MatchString_ValueType" json:"value_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MatchString) Reset()         { *m = MatchString{} }
func (m *MatchString) String() string { return proto.CompactTextString(m) }
func (*MatchString) ProtoMessage()    {}
func (*MatchString) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_99674f3804a5acf0, []int{1}
}
func (m *MatchString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchString.Unmarshal(m, b)
}
func (m *MatchString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchString.Marshal(b, m, deterministic)
}
func (dst *MatchString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchString.Merge(dst, src)
}
func (m *MatchString) XXX_Size() int {
	return xxx_messageInfo_MatchString.Size(m)
}
func (m *MatchString) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchString.DiscardUnknown(m)
}

var xxx_messageInfo_MatchString proto.InternalMessageInfo

func (m *MatchString) GetType() MatchString_MatchStringType {
	if m != nil {
		return m.Type
	}
	return MatchString_EXACT
}

func (m *MatchString) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MatchString) GetValueType() MatchString_ValueType {
	if m != nil {
		return m.ValueType
	}
	return MatchString_TEXT
}

type OptionSwitch struct {
	Options              map[string]string `protobuf:"bytes,1,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OptionSwitch) Reset()         { *m = OptionSwitch{} }
func (m *OptionSwitch) String() string { return proto.CompactTextString(m) }
func (*OptionSwitch) ProtoMessage()    {}
func (*OptionSwitch) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_99674f3804a5acf0, []int{2}
}
func (m *OptionSwitch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionSwitch.Unmarshal(m, b)
}
func (m *OptionSwitch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionSwitch.Marshal(b, m, deterministic)
}
func (dst *OptionSwitch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionSwitch.Merge(dst, src)
}
func (m *OptionSwitch) XXX_Size() int {
	return xxx_messageInfo_OptionSwitch.Size(m)
}
func (m *OptionSwitch) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionSwitch.DiscardUnknown(m)
}

var xxx_messageInfo_OptionSwitch proto.InternalMessageInfo

func (m *OptionSwitch) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "v1.Location")
	proto.RegisterType((*MatchString)(nil), "v1.MatchString")
	proto.RegisterType((*OptionSwitch)(nil), "v1.OptionSwitch")
	proto.RegisterMapType((map[string]string)(nil), "v1.OptionSwitch.OptionsEntry")
	proto.RegisterEnum("v1.MatchString_MatchStringType", MatchString_MatchStringType_name, MatchString_MatchStringType_value)
	proto.RegisterEnum("v1.MatchString_ValueType", MatchString_ValueType_name, MatchString_ValueType_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_99674f3804a5acf0) }

var fileDescriptor_model_99674f3804a5acf0 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x6f, 0x9b, 0x40,
	0x10, 0xc5, 0xbb, 0x40, 0xa8, 0x19, 0xd2, 0x74, 0xb5, 0xea, 0x81, 0x46, 0xfd, 0x13, 0x71, 0xca,
	0x89, 0x34, 0xa4, 0x52, 0xa3, 0xdc, 0x68, 0xb5, 0xaa, 0xac, 0xfa, 0x4f, 0xbb, 0xa6, 0x96, 0x6f,
	0x16, 0xa6, 0x5b, 0x8c, 0x6a, 0xb3, 0x08, 0x30, 0x16, 0xbd, 0xf6, 0xcb, 0xf4, 0xda, 0x6f, 0x58,
	0xed, 0x62, 0x22, 0xe4, 0x93, 0x6f, 0x6f, 0x67, 0x7e, 0x33, 0x6f, 0x78, 0x80, 0xbd, 0x15, 0x3f,
	0xf8, 0xc6, 0xcb, 0x0b, 0x51, 0x09, 0xa2, 0xd5, 0xb7, 0x97, 0x6f, 0x12, 0x21, 0x92, 0x0d, 0xbf,
	0x51, 0x95, 0xd5, 0xee, 0xe7, 0xcd, 0xbe, 0x88, 0xf2, 0x9c, 0x17, 0x65, 0xcb, 0xb8, 0x7f, 0x11,
	0x0c, 0x46, 0x22, 0x8e, 0xaa, 0x54, 0x64, 0xe4, 0x3d, 0x98, 0x05, 0x4f, 0x52, 0x91, 0x39, 0xe8,
	0x0a, 0x5d, 0xdb, 0xfe, 0x2b, 0xaf, 0x9d, 0xf6, 0xba, 0x69, 0x6f, 0x56, 0x15, 0x69, 0x96, 0xcc,
	0xa3, 0xcd, 0x8e, 0xb3, 0x03, 0x4b, 0xde, 0x81, 0xf1, 0x5b, 0x64, 0xdc, 0xd1, 0x4e, 0x98, 0x51,
	0xa4, 0xf4, 0x89, 0xa3, 0x6d, 0xbe, 0x2b, 0x1d, 0xfd, 0x14, 0x9f, 0x96, 0x75, 0xff, 0x69, 0x60,
	0x8f, 0xa3, 0x2a, 0x5e, 0xb7, 0x4d, 0x72, 0x07, 0x46, 0xd5, 0xe4, 0x5c, 0xdd, 0x7a, 0xe1, 0xbf,
	0xf5, 0xea, 0x5b, 0xaf, 0xd7, 0xee, 0xeb, 0xb0, 0xc9, 0x39, 0x53, 0x30, 0xf1, 0xe1, 0xac, 0x96,
	0x5b, 0x4f, 0xba, 0xb6, 0x45, 0xc9, 0x3d, 0x80, 0x12, 0x4b, 0x65, 0xa7, 0x2b, 0xbb, 0x97, 0xc7,
	0x76, 0x6a, 0x42, 0x19, 0x59, 0x75, 0x27, 0xdd, 0x2f, 0xf0, 0xfc, 0xe8, 0x0c, 0x62, 0xc1, 0x19,
	0x5d, 0x04, 0x9f, 0x42, 0xfc, 0x44, 0x4a, 0x46, 0x3f, 0xd3, 0x05, 0x46, 0xe4, 0x02, 0x60, 0x32,
	0x0d, 0x97, 0xf4, 0xdb, 0xf7, 0x60, 0x34, 0xc3, 0x1a, 0x31, 0x41, 0x1b, 0x4e, 0xb0, 0x4e, 0x00,
	0x4c, 0x59, 0x1f, 0x4e, 0xb0, 0xe1, 0xfa, 0x60, 0x3d, 0x9a, 0x90, 0x01, 0x18, 0x21, 0x5d, 0xc8,
	0x2d, 0xcf, 0xc0, 0xfa, 0x1a, 0xb0, 0x60, 0x4c, 0x43, 0xca, 0x30, 0x22, 0xe7, 0x30, 0x98, 0x07,
	0x6c, 0x18, 0x7c, 0x1c, 0x51, 0xac, 0xb9, 0x7f, 0x10, 0x9c, 0x4f, 0x73, 0xf9, 0x73, 0x67, 0xfb,
	0xb4, 0x8a, 0xd7, 0xe4, 0x03, 0x3c, 0x15, 0xea, 0x5d, 0x3a, 0xe8, 0x4a, 0xbf, 0xb6, 0xfd, 0xd7,
	0xf2, 0x43, 0xfa, 0xc8, 0xe1, 0x51, 0xd2, 0xac, 0x2a, 0x1a, 0xd6, 0xd1, 0x97, 0x0f, 0xdd, 0xa2,
	0xb6, 0x41, 0x30, 0xe8, 0xbf, 0x78, 0xa3, 0xc2, 0xb7, 0x98, 0x94, 0xe4, 0x45, 0x3f, 0x5a, 0xeb,
	0x10, 0xde, 0x83, 0x76, 0x8f, 0x56, 0xa6, 0x4a, 0xf7, 0xee, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0xee, 0xf2, 0x16, 0x9e, 0x02, 0x00, 0x00,
}
