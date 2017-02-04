// Code generated by protoc-gen-go.
// source: networking_platform.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PlatformRequestType int32

const (
	PlatformRequestType_METHOD_UNSET             PlatformRequestType = 0
	PlatformRequestType_BUY_ITEM_POKECOINS       PlatformRequestType = 2
	PlatformRequestType_BUY_ITEM_ANDROID         PlatformRequestType = 3
	PlatformRequestType_BUY_ITEM_IOS             PlatformRequestType = 4
	PlatformRequestType_GET_STORE_ITEMS          PlatformRequestType = 5
	PlatformRequestType_SEND_ENCRYPTED_SIGNATURE PlatformRequestType = 6
	PlatformRequestType_UNKNOWN_PTR_8            PlatformRequestType = 8
)

var PlatformRequestType_name = map[int32]string{
	0: "METHOD_UNSET",
	2: "BUY_ITEM_POKECOINS",
	3: "BUY_ITEM_ANDROID",
	4: "BUY_ITEM_IOS",
	5: "GET_STORE_ITEMS",
	6: "SEND_ENCRYPTED_SIGNATURE",
	8: "UNKNOWN_PTR_8",
}
var PlatformRequestType_value = map[string]int32{
	"METHOD_UNSET":             0,
	"BUY_ITEM_POKECOINS":       2,
	"BUY_ITEM_ANDROID":         3,
	"BUY_ITEM_IOS":             4,
	"GET_STORE_ITEMS":          5,
	"SEND_ENCRYPTED_SIGNATURE": 6,
	"UNKNOWN_PTR_8":            8,
}

func (x PlatformRequestType) String() string {
	return proto.EnumName(PlatformRequestType_name, int32(x))
}
func (PlatformRequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

func init() {
	proto.RegisterEnum("POGOProtos.Networking.Platform.PlatformRequestType", PlatformRequestType_name, PlatformRequestType_value)
}

func init() { proto.RegisterFile("networking_platform.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xfd, 0x5b, 0xca, 0xa2, 0x38, 0x6e, 0x45, 0x14, 0xc4, 0x07, 0xf0, 0x90, 0x8b, 0x17,
	0xaf, 0x6d, 0x77, 0x88, 0x4b, 0xe9, 0xcc, 0xb2, 0x3b, 0x41, 0xea, 0x65, 0x50, 0x88, 0x22, 0x6a,
	0x13, 0x93, 0x88, 0xf8, 0x50, 0xbe, 0xa3, 0xb4, 0xda, 0x9c, 0x66, 0xbe, 0x1f, 0xbf, 0xef, 0xf0,
	0x99, 0xf3, 0x65, 0xd9, 0x7d, 0x55, 0xcd, 0xeb, 0xcb, 0xf2, 0x59, 0xeb, 0xb7, 0x87, 0xee, 0xa9,
	0x6a, 0xde, 0xb3, 0xba, 0xa9, 0xba, 0xca, 0x5e, 0x06, 0xce, 0x39, 0xac, 0xde, 0x36, 0xa3, 0xde,
	0xca, 0xc2, 0xbf, 0x75, 0xf5, 0xb3, 0x6d, 0x46, 0x9b, 0x10, 0xcb, 0x8f, 0xcf, 0xb2, 0xed, 0xe4,
	0xbb, 0x2e, 0x2d, 0x98, 0x83, 0x39, 0xca, 0x2d, 0x3b, 0x2d, 0x28, 0xa1, 0xc0, 0x96, 0x3d, 0x35,
	0x76, 0x52, 0x2c, 0xd4, 0x0b, 0xce, 0x35, 0xf0, 0x0c, 0xa7, 0xec, 0x29, 0xc1, 0x8e, 0x3d, 0x31,
	0xd0, 0xf3, 0x31, 0xb9, 0xc8, 0xde, 0xc1, 0xee, 0xaa, 0xdf, 0x53, 0xcf, 0x09, 0xf6, 0xec, 0xc8,
	0x1c, 0xe5, 0x28, 0x9a, 0x84, 0x23, 0xae, 0x79, 0x82, 0x7d, 0x7b, 0x61, 0xce, 0x12, 0x92, 0x53,
	0xa4, 0x69, 0x5c, 0x04, 0x41, 0xa7, 0xc9, 0xe7, 0x34, 0x96, 0x22, 0x22, 0x0c, 0xec, 0xb1, 0x39,
	0x2c, 0x68, 0x46, 0x7c, 0x47, 0x1a, 0x24, 0xea, 0x0d, 0x0c, 0x27, 0xc3, 0xfb, 0xc1, 0x7a, 0x58,
	0xfb, 0xf8, 0x77, 0xaf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9d, 0xe9, 0x27, 0xb5, 0xfd, 0x00,
	0x00, 0x00,
}
