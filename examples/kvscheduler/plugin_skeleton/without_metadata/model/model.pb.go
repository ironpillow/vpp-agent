// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/model.proto

package model

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

type ValueSkeleton struct {
	//
	//logical name is often defined to build unique keys for value instances
	//- alternativelly, in the value model (keys.go), you may apply the
	//WithNameTemplate() option to generate value instance name using a golang
	//template, combining multiple value attributes that collectively
	//guarantee unique value identification (i.e. primary key)
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueSkeleton) Reset()         { *m = ValueSkeleton{} }
func (m *ValueSkeleton) String() string { return proto.CompactTextString(m) }
func (*ValueSkeleton) ProtoMessage()    {}
func (*ValueSkeleton) Descriptor() ([]byte, []int) {
	return fileDescriptor_312ac5bcab6cbb43, []int{0}
}

func (m *ValueSkeleton) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueSkeleton.Unmarshal(m, b)
}
func (m *ValueSkeleton) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueSkeleton.Marshal(b, m, deterministic)
}
func (m *ValueSkeleton) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueSkeleton.Merge(m, src)
}
func (m *ValueSkeleton) XXX_Size() int {
	return xxx_messageInfo_ValueSkeleton.Size(m)
}
func (m *ValueSkeleton) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueSkeleton.DiscardUnknown(m)
}

var xxx_messageInfo_ValueSkeleton proto.InternalMessageInfo

func (m *ValueSkeleton) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ValueSkeleton)(nil), "model.ValueSkeleton")
}

func init() { proto.RegisterFile("model/model.proto", fileDescriptor_312ac5bcab6cbb43) }

var fileDescriptor_312ac5bcab6cbb43 = []byte{
	// 82 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x32,
	0x17, 0x6f, 0x58, 0x62, 0x4e, 0x69, 0x6a, 0x70, 0x76, 0x6a, 0x4e, 0x6a, 0x49, 0x7e, 0x9e, 0x90,
	0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x9d,
	0xc4, 0x06, 0xd6, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x82, 0xe4, 0x91, 0x47, 0x00,
	0x00, 0x00,
}
