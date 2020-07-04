// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login_10.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	login_10.proto

It has these top-level messages:
	C2S_Login_10001
	S2C_Login_10001
	C2S_Create_10002
	S2C_Create_10002
	PlayerInfo
*/
package protos

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

type C2S_Login_10001 struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId" json:"AccountId,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"`
}

func (m *C2S_Login_10001) Reset()                    { *m = C2S_Login_10001{} }
func (m *C2S_Login_10001) String() string            { return proto.CompactTextString(m) }
func (*C2S_Login_10001) ProtoMessage()               {}
func (*C2S_Login_10001) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S_Login_10001) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *C2S_Login_10001) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type S2C_Login_10001 struct {
	Code int32       `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
	Info *PlayerInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *S2C_Login_10001) Reset()                    { *m = S2C_Login_10001{} }
func (m *S2C_Login_10001) String() string            { return proto.CompactTextString(m) }
func (*S2C_Login_10001) ProtoMessage()               {}
func (*S2C_Login_10001) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C_Login_10001) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2C_Login_10001) GetInfo() *PlayerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type C2S_Create_10002 struct {
	Nick string `protobuf:"bytes,1,opt,name=Nick" json:"Nick,omitempty"`
}

func (m *C2S_Create_10002) Reset()                    { *m = C2S_Create_10002{} }
func (m *C2S_Create_10002) String() string            { return proto.CompactTextString(m) }
func (*C2S_Create_10002) ProtoMessage()               {}
func (*C2S_Create_10002) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *C2S_Create_10002) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

type S2C_Create_10002 struct {
	Code int32       `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
	Info *PlayerInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *S2C_Create_10002) Reset()                    { *m = S2C_Create_10002{} }
func (m *S2C_Create_10002) String() string            { return proto.CompactTextString(m) }
func (*S2C_Create_10002) ProtoMessage()               {}
func (*S2C_Create_10002) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *S2C_Create_10002) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *S2C_Create_10002) GetInfo() *PlayerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type PlayerInfo struct {
	Id      int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Nick    string `protobuf:"bytes,2,opt,name=nick" json:"nick,omitempty"`
	HeadImg string `protobuf:"bytes,3,opt,name=headImg" json:"headImg,omitempty"`
	Vip     int32  `protobuf:"varint,4,opt,name=vip" json:"vip,omitempty"`
	Level   int32  `protobuf:"varint,5,opt,name=level" json:"level,omitempty"`
	Coin    int64  `protobuf:"varint,6,opt,name=coin" json:"coin,omitempty"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PlayerInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlayerInfo) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *PlayerInfo) GetHeadImg() string {
	if m != nil {
		return m.HeadImg
	}
	return ""
}

func (m *PlayerInfo) GetVip() int32 {
	if m != nil {
		return m.Vip
	}
	return 0
}

func (m *PlayerInfo) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *PlayerInfo) GetCoin() int64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

func init() {
	proto.RegisterType((*C2S_Login_10001)(nil), "protos.C2S_Login_10001")
	proto.RegisterType((*S2C_Login_10001)(nil), "protos.S2C_Login_10001")
	proto.RegisterType((*C2S_Create_10002)(nil), "protos.C2S_Create_10002")
	proto.RegisterType((*S2C_Create_10002)(nil), "protos.S2C_Create_10002")
	proto.RegisterType((*PlayerInfo)(nil), "protos.PlayerInfo")
}

func init() { proto.RegisterFile("login_10.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xc9, 0xbf, 0x4a, 0x47, 0x68, 0xc3, 0xe0, 0x61, 0x0f, 0x1e, 0x24, 0x87, 0xe2, 0xa9,
	0xa4, 0xf1, 0x09, 0x24, 0x78, 0x08, 0x68, 0x91, 0xd4, 0x7b, 0x89, 0xc9, 0xb6, 0x2e, 0x5d, 0x77,
	0x4a, 0x1a, 0x0b, 0xbe, 0x80, 0xcf, 0x2d, 0x33, 0x5b, 0x29, 0x7a, 0xf4, 0x94, 0x6f, 0xbe, 0xc9,
	0xfe, 0xe6, 0xdb, 0x59, 0x98, 0x58, 0xda, 0x1a, 0xb7, 0x5e, 0xe4, 0xf3, 0x7d, 0x4f, 0x03, 0xe1,
	0x48, 0x3e, 0x87, 0xec, 0x01, 0xa6, 0x65, 0xb1, 0x5a, 0x3f, 0x9e, 0xba, 0x79, 0xbe, 0xc0, 0x6b,
	0x18, 0xdf, 0xb7, 0x2d, 0x7d, 0xb8, 0xa1, 0xea, 0x54, 0x70, 0x13, 0xdc, 0x8e, 0xeb, 0xb3, 0x81,
	0x57, 0x90, 0xbc, 0xd0, 0x4e, 0x3b, 0x15, 0x4a, 0xc7, 0x17, 0xd9, 0x13, 0x4c, 0x57, 0x45, 0xf9,
	0x0b, 0x83, 0x10, 0x97, 0xd4, 0x69, 0x21, 0x24, 0xb5, 0x68, 0x9c, 0x41, 0x6c, 0xdc, 0x86, 0xe4,
	0xec, 0x65, 0x81, 0x3e, 0xcb, 0x61, 0xfe, 0x6c, 0x9b, 0x4f, 0xdd, 0x57, 0x6e, 0x43, 0xb5, 0xf4,
	0xb3, 0x19, 0xa4, 0x9c, 0xaa, 0xec, 0x75, 0x33, 0x68, 0xe1, 0x15, 0xcc, 0x5b, 0x9a, 0x76, 0x77,
	0x4a, 0x24, 0x3a, 0x5b, 0x42, 0xca, 0x63, 0xff, 0xfe, 0xf7, 0xef, 0xb9, 0x5f, 0x01, 0xc0, 0xd9,
	0xc4, 0x09, 0x84, 0xc6, 0xaf, 0x20, 0xaa, 0x43, 0xd3, 0x31, 0xda, 0x71, 0x04, 0x7f, 0x75, 0xd1,
	0xa8, 0xe0, 0xe2, 0x4d, 0x37, 0x5d, 0xf5, 0xbe, 0x55, 0x91, 0xd8, 0x3f, 0x25, 0xa6, 0x10, 0x1d,
	0xcd, 0x5e, 0xc5, 0x92, 0x83, 0x25, 0xef, 0xce, 0xea, 0xa3, 0xb6, 0x2a, 0x11, 0xcf, 0x17, 0x4c,
	0x6d, 0xc9, 0x38, 0x35, 0x92, 0x39, 0xa2, 0x5f, 0xfd, 0xf3, 0xdc, 0x7d, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xaf, 0x1d, 0xf6, 0x1c, 0xb7, 0x01, 0x00, 0x00,
}