// Code generated by protoc-gen-go.
// source: networking_platform_requests.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BuyItemAndroidRequest struct {
	BuyItemIntent string `protobuf:"bytes,1,opt,name=buy_item_intent" json:"buy_item_intent,omitempty"`
}

func (m *BuyItemAndroidRequest) Reset()                    { *m = BuyItemAndroidRequest{} }
func (m *BuyItemAndroidRequest) String() string            { return proto.CompactTextString(m) }
func (*BuyItemAndroidRequest) ProtoMessage()               {}
func (*BuyItemAndroidRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type BuyItemPokeCoinsRequest struct {
	ItemId string `protobuf:"bytes,1,opt,name=item_id" json:"item_id,omitempty"`
}

func (m *BuyItemPokeCoinsRequest) Reset()                    { *m = BuyItemPokeCoinsRequest{} }
func (m *BuyItemPokeCoinsRequest) String() string            { return proto.CompactTextString(m) }
func (*BuyItemPokeCoinsRequest) ProtoMessage()               {}
func (*BuyItemPokeCoinsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type SendEncryptedSignatureRequest struct {
	EncryptedSignature []byte `protobuf:"bytes,1,opt,name=encrypted_signature,proto3" json:"encrypted_signature,omitempty"`
}

func (m *SendEncryptedSignatureRequest) Reset()                    { *m = SendEncryptedSignatureRequest{} }
func (m *SendEncryptedSignatureRequest) String() string            { return proto.CompactTextString(m) }
func (*SendEncryptedSignatureRequest) ProtoMessage()               {}
func (*SendEncryptedSignatureRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func init() {
	proto.RegisterType((*BuyItemAndroidRequest)(nil), "POGOProtos.Networking.Platform.Requests.BuyItemAndroidRequest")
	proto.RegisterType((*BuyItemPokeCoinsRequest)(nil), "POGOProtos.Networking.Platform.Requests.BuyItemPokeCoinsRequest")
	proto.RegisterType((*SendEncryptedSignatureRequest)(nil), "POGOProtos.Networking.Platform.Requests.SendEncryptedSignatureRequest")
}

func init() { proto.RegisterFile("networking_platform_requests.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0xca, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0xce, 0xcc, 0x4b, 0x8f, 0x2f, 0xc8, 0x49, 0x2c, 0x49, 0xcb, 0x2f, 0xca, 0x8d,
	0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52,
	0x0f, 0xf0, 0x77, 0xf7, 0x0f, 0x00, 0x31, 0x8b, 0xf5, 0xfc, 0xe0, 0xca, 0xf5, 0x02, 0xa0, 0xca,
	0xf5, 0x82, 0xa0, 0xca, 0x95, 0x0c, 0xb8, 0x44, 0x9d, 0x4a, 0x2b, 0x3d, 0x4b, 0x52, 0x73, 0x1d,
	0xf3, 0x52, 0x8a, 0xf2, 0x33, 0x53, 0xa0, 0x32, 0x42, 0xe2, 0x5c, 0xfc, 0x49, 0xa5, 0x95, 0xf1,
	0x99, 0x25, 0xa9, 0xb9, 0xf1, 0x99, 0x79, 0x25, 0xa9, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x4a, 0x5a, 0x5c, 0xe2, 0x50, 0x1d, 0x01, 0xf9, 0xd9, 0xa9, 0xce, 0xf9, 0x99, 0x79, 0xc5,
	0x30, 0x3d, 0xfc, 0x5c, 0xec, 0x10, 0xf5, 0x29, 0x50, 0xb5, 0x36, 0x5c, 0xb2, 0xc1, 0xa9, 0x79,
	0x29, 0xae, 0x79, 0xc9, 0x45, 0x95, 0x05, 0x25, 0xa9, 0x29, 0xc1, 0x99, 0xe9, 0x79, 0x89, 0x25,
	0xa5, 0x45, 0xa9, 0x30, 0x1d, 0xd2, 0x5c, 0xc2, 0xa9, 0x30, 0xc9, 0xf8, 0x62, 0x98, 0x2c, 0x58,
	0x37, 0x8f, 0x13, 0x47, 0x14, 0x1b, 0xd8, 0x37, 0xc5, 0x49, 0x10, 0xda, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x31, 0xf8, 0xea, 0x01, 0xfb, 0x00, 0x00, 0x00,
}
