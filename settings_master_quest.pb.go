// Code generated by protoc-gen-go.
// source: settings_master_quest.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DailyQuestSettings struct {
	BucketsPerDay         int32   `protobuf:"varint,1,opt,name=buckets_per_day" json:"buckets_per_day,omitempty"`
	StreakLength          int32   `protobuf:"varint,2,opt,name=streak_length" json:"streak_length,omitempty"`
	BonusMultiplier       float32 `protobuf:"fixed32,3,opt,name=bonus_multiplier" json:"bonus_multiplier,omitempty"`
	StreakBonusMultiplier float32 `protobuf:"fixed32,4,opt,name=streak_bonus_multiplier" json:"streak_bonus_multiplier,omitempty"`
}

func (m *DailyQuestSettings) Reset()                    { *m = DailyQuestSettings{} }
func (m *DailyQuestSettings) String() string            { return proto.CompactTextString(m) }
func (*DailyQuestSettings) ProtoMessage()               {}
func (*DailyQuestSettings) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func init() {
	proto.RegisterType((*DailyQuestSettings)(nil), "POGOProtos.Settings.Master.Quest.DailyQuestSettings")
}

func init() { proto.RegisterFile("settings_master_quest.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0x2d, 0x29,
	0xc9, 0xcc, 0x4b, 0x2f, 0x8e, 0xcf, 0x4d, 0x2c, 0x2e, 0x49, 0x2d, 0x8a, 0x2f, 0x2c, 0x4d, 0x2d,
	0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x08, 0xf0, 0x77, 0xf7, 0x0f, 0x00, 0x31,
	0x8b, 0xf5, 0x82, 0xa1, 0xea, 0xf4, 0x7c, 0xc1, 0xea, 0xf4, 0x02, 0x41, 0xea, 0x94, 0xea, 0xb9,
	0x84, 0x5c, 0x12, 0x33, 0x73, 0x2a, 0xc1, 0x3c, 0x98, 0x12, 0x21, 0x71, 0x2e, 0xfe, 0xa4, 0xd2,
	0xe4, 0xec, 0xd4, 0x92, 0xe2, 0xf8, 0x82, 0xd4, 0xa2, 0xf8, 0x94, 0xc4, 0x4a, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x56, 0x21, 0x51, 0x2e, 0xde, 0xe2, 0x92, 0xa2, 0xd4, 0xc4, 0xec, 0xf8, 0x9c, 0xd4,
	0xbc, 0xf4, 0x92, 0x0c, 0x09, 0x26, 0xb0, 0xb0, 0x04, 0x97, 0x40, 0x52, 0x7e, 0x5e, 0x69, 0x71,
	0x7c, 0x6e, 0x69, 0x4e, 0x49, 0x66, 0x41, 0x4e, 0x66, 0x6a, 0x91, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0x93, 0x90, 0x3c, 0x97, 0x38, 0x54, 0x03, 0x86, 0x02, 0x16, 0x90, 0x02, 0x27, 0x8e, 0x28, 0x36,
	0xb0, 0x5b, 0x8b, 0x93, 0x20, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x2a, 0x5f, 0x3a,
	0xd2, 0x00, 0x00, 0x00,
}
