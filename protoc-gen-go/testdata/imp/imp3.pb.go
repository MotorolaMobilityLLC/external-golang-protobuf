// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imp/imp3.proto

package imp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ForeignImportedMessage struct {
	Tuber                *string  `protobuf:"bytes,1,opt,name=tuber" json:"tuber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForeignImportedMessage) Reset()                    { *m = ForeignImportedMessage{} }
func (m *ForeignImportedMessage) String() string            { return proto.CompactTextString(m) }
func (*ForeignImportedMessage) ProtoMessage()               {}
func (*ForeignImportedMessage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }
func (m *ForeignImportedMessage) Unmarshal(b []byte) error {
	return xxx_messageInfo_ForeignImportedMessage.Unmarshal(m, b)
}
func (m *ForeignImportedMessage) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForeignImportedMessage.Marshal(b, m, deterministic)
}
func (dst *ForeignImportedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForeignImportedMessage.Merge(dst, src)
}
func (m *ForeignImportedMessage) XXX_Size() int {
	return xxx_messageInfo_ForeignImportedMessage.Size(m)
}
func (m *ForeignImportedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ForeignImportedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ForeignImportedMessage proto.InternalMessageInfo

func (m *ForeignImportedMessage) GetTuber() string {
	if m != nil && m.Tuber != nil {
		return *m.Tuber
	}
	return ""
}

func init() {
	proto.RegisterType((*ForeignImportedMessage)(nil), "imp.ForeignImportedMessage")
}

func init() { proto.RegisterFile("imp/imp3.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2d, 0xd0,
	0xcf, 0xcc, 0x2d, 0x30, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xcc, 0x2d, 0x50,
	0xd2, 0xe3, 0x12, 0x73, 0xcb, 0x2f, 0x4a, 0xcd, 0x4c, 0xcf, 0xf3, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a,
	0x49, 0x4d, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x29, 0x4d,
	0x4a, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x9c, 0xcc, 0xa3, 0x4c, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2,
	0xf5, 0xc1, 0xe6, 0x25, 0x95, 0xa6, 0x41, 0x18, 0xc9, 0xba, 0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xf9,
	0xfa, 0x25, 0xa9, 0xc5, 0x25, 0x29, 0x89, 0x25, 0x89, 0x20, 0x4b, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0xbf, 0xbe, 0xdc, 0x7e, 0x00, 0x00, 0x00,
}
