// Code generated by protoc-gen-go.
// source: maps.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of FortLureInfo from map_fort.proto

// Ignoring public import of FortModifier from map_fort.proto

// Ignoring public import of FortSummary from map_fort.proto

// Ignoring public import of FortData from map_fort.proto

// Ignoring public import of FortType from map_fort.proto

// Ignoring public import of FortSponsor from map_fort.proto

// Ignoring public import of FortRenderingType from map_fort.proto

// Ignoring public import of WildPokemon from map_pokemon.proto

// Ignoring public import of NearbyPokemon from map_pokemon.proto

// Ignoring public import of MapPokemon from map_pokemon.proto

type MapObjectsStatus int32

const (
	MapObjectsStatus_UNSET_STATUS   MapObjectsStatus = 0
	MapObjectsStatus_SUCCESS        MapObjectsStatus = 1
	MapObjectsStatus_LOCATION_UNSET MapObjectsStatus = 2
)

var MapObjectsStatus_name = map[int32]string{
	0: "UNSET_STATUS",
	1: "SUCCESS",
	2: "LOCATION_UNSET",
}
var MapObjectsStatus_value = map[string]int32{
	"UNSET_STATUS":   0,
	"SUCCESS":        1,
	"LOCATION_UNSET": 2,
}

func (x MapObjectsStatus) String() string {
	return proto.EnumName(MapObjectsStatus_name, int32(x))
}
func (MapObjectsStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type SpawnPoint struct {
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,3,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *SpawnPoint) Reset()                    { *m = SpawnPoint{} }
func (m *SpawnPoint) String() string            { return proto.CompactTextString(m) }
func (*SpawnPoint) ProtoMessage()               {}
func (*SpawnPoint) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type MapCell struct {
	// S2 geographic area that the cell covers (http://s2map.com/) (https://code.google.com/archive/p/s2-geometry-library/)
	S2CellId             uint64         `protobuf:"varint,1,opt,name=s2_cell_id" json:"s2_cell_id,omitempty"`
	CurrentTimestampMs   int64          `protobuf:"varint,2,opt,name=current_timestamp_ms" json:"current_timestamp_ms,omitempty"`
	Forts                []*FortData    `protobuf:"bytes,3,rep,name=forts" json:"forts,omitempty"`
	SpawnPoints          []*SpawnPoint  `protobuf:"bytes,4,rep,name=spawn_points" json:"spawn_points,omitempty"`
	DeletedObjects       []string       `protobuf:"bytes,6,rep,name=deleted_objects" json:"deleted_objects,omitempty"`
	IsTruncatedList      bool           `protobuf:"varint,7,opt,name=is_truncated_list" json:"is_truncated_list,omitempty"`
	FortSummaries        []*FortSummary `protobuf:"bytes,8,rep,name=fort_summaries" json:"fort_summaries,omitempty"`
	DecimatedSpawnPoints []*SpawnPoint  `protobuf:"bytes,9,rep,name=decimated_spawn_points" json:"decimated_spawn_points,omitempty"`
	// Pokemon within 2 steps or less.
	WildPokemons []*WildPokemon `protobuf:"bytes,5,rep,name=wild_pokemons" json:"wild_pokemons,omitempty"`
	// Pokemon within 1 step or none.
	CatchablePokemons []*MapPokemon `protobuf:"bytes,10,rep,name=catchable_pokemons" json:"catchable_pokemons,omitempty"`
	// Pokemon farther away than 2 steps, but still in the area.
	NearbyPokemons []*NearbyPokemon `protobuf:"bytes,11,rep,name=nearby_pokemons" json:"nearby_pokemons,omitempty"`
}

func (m *MapCell) Reset()                    { *m = MapCell{} }
func (m *MapCell) String() string            { return proto.CompactTextString(m) }
func (*MapCell) ProtoMessage()               {}
func (*MapCell) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *MapCell) GetForts() []*FortData {
	if m != nil {
		return m.Forts
	}
	return nil
}

func (m *MapCell) GetSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.SpawnPoints
	}
	return nil
}

func (m *MapCell) GetFortSummaries() []*FortSummary {
	if m != nil {
		return m.FortSummaries
	}
	return nil
}

func (m *MapCell) GetDecimatedSpawnPoints() []*SpawnPoint {
	if m != nil {
		return m.DecimatedSpawnPoints
	}
	return nil
}

func (m *MapCell) GetWildPokemons() []*WildPokemon {
	if m != nil {
		return m.WildPokemons
	}
	return nil
}

func (m *MapCell) GetCatchablePokemons() []*MapPokemon {
	if m != nil {
		return m.CatchablePokemons
	}
	return nil
}

func (m *MapCell) GetNearbyPokemons() []*NearbyPokemon {
	if m != nil {
		return m.NearbyPokemons
	}
	return nil
}

func init() {
	proto.RegisterType((*SpawnPoint)(nil), "POGOProtos.Map.SpawnPoint")
	proto.RegisterType((*MapCell)(nil), "POGOProtos.Map.MapCell")
	proto.RegisterEnum("POGOProtos.Map.MapObjectsStatus", MapObjectsStatus_name, MapObjectsStatus_value)
}

func init() { proto.RegisterFile("maps.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xe1, 0x8b, 0xd3, 0x30,
	0x18, 0x87, 0xaf, 0xb7, 0xdd, 0x6d, 0x7b, 0x77, 0xee, 0xba, 0x20, 0x5a, 0x87, 0x42, 0x99, 0x08,
	0x45, 0x64, 0xe8, 0xf9, 0x45, 0xee, 0x83, 0x70, 0xce, 0x29, 0x82, 0x5b, 0x0b, 0xd9, 0x10, 0xfc,
	0x12, 0xde, 0xb5, 0x51, 0xa3, 0x49, 0x1b, 0x9a, 0x94, 0xe3, 0x3e, 0xfa, 0x9f, 0x4b, 0xd2, 0xc9,
	0xbc, 0x83, 0xe1, 0x97, 0xf2, 0xe6, 0xcd, 0xfb, 0xfc, 0xfa, 0xb4, 0x09, 0x80, 0x42, 0x6d, 0x66,
	0xba, 0xae, 0x6c, 0x45, 0x46, 0x59, 0xfa, 0x31, 0xcd, 0x5c, 0x69, 0x66, 0x4b, 0xd4, 0x93, 0x91,
	0x42, 0xcd, 0xbe, 0x55, 0xb5, 0x6d, 0xf7, 0x27, 0x63, 0xb7, 0xd6, 0xd5, 0x2f, 0xae, 0xaa, 0xb2,
	0x6d, 0x4d, 0x5f, 0x01, 0x50, 0x8d, 0xd7, 0x65, 0x56, 0x89, 0xd2, 0x92, 0x10, 0xfa, 0x12, 0xad,
	0xb0, 0x4d, 0xc1, 0xa3, 0xe3, 0x38, 0x48, 0x02, 0x32, 0x86, 0x81, 0xac, 0xca, 0xef, 0x6d, 0xab,
	0xe3, 0x5a, 0xd3, 0xdf, 0x5d, 0xe8, 0x2d, 0x51, 0xcf, 0xb9, 0x94, 0x84, 0x00, 0x98, 0x0b, 0x96,
	0x73, 0x29, 0x99, 0x28, 0xa2, 0x20, 0x0e, 0x92, 0x2e, 0x79, 0x0c, 0xf7, 0xf3, 0xa6, 0xae, 0x79,
	0x69, 0x99, 0x15, 0x8a, 0x1b, 0x8b, 0x4a, 0x33, 0x65, 0x7c, 0x60, 0x87, 0xbc, 0x80, 0x13, 0x67,
	0x64, 0xa2, 0x4e, 0xdc, 0x49, 0x86, 0x17, 0x4f, 0x66, 0xb7, 0x9d, 0x67, 0x1f, 0x9c, 0xae, 0x7b,
	0xbc, 0x47, 0x8b, 0xe4, 0x25, 0x9c, 0x19, 0xa7, 0xc7, 0xb4, 0xf3, 0x33, 0x51, 0xd7, 0x43, 0x93,
	0xbb, 0xd0, 0x3f, 0x9f, 0xf0, 0x10, 0xce, 0x0b, 0x2e, 0xb9, 0xe5, 0x05, 0xab, 0xb6, 0x3f, 0x79,
	0x6e, 0x4d, 0x74, 0x1a, 0x77, 0x92, 0x01, 0x79, 0x04, 0x63, 0x61, 0x98, 0xad, 0x9b, 0x32, 0x47,
	0xb7, 0x2b, 0x85, 0xb1, 0x51, 0x2f, 0x0e, 0x92, 0x3e, 0x79, 0x03, 0x23, 0xe7, 0xc4, 0x4c, 0xa3,
	0x14, 0xd6, 0x82, 0x9b, 0xa8, 0xef, 0xdf, 0x13, 0x1f, 0x94, 0xa3, 0x7e, 0xf2, 0x86, 0x5c, 0xc2,
	0x83, 0x82, 0xe7, 0x42, 0xf9, 0xc4, 0x5b, 0xa6, 0x83, 0xff, 0x9a, 0x5e, 0xc2, 0xbd, 0x6b, 0x21,
	0x8b, 0xbf, 0x07, 0x62, 0xa2, 0x13, 0x8f, 0x3c, 0xbd, 0x8b, 0x64, 0xbb, 0x03, 0xfb, 0x22, 0x64,
	0xb1, 0xab, 0xc9, 0x5b, 0x20, 0x39, 0xda, 0xfc, 0x07, 0x6e, 0x25, 0xdf, 0x07, 0x80, 0x0f, 0x98,
	0x1e, 0x0a, 0x58, 0xa2, 0xde, 0xf3, 0xe7, 0x25, 0xc7, 0x7a, 0x7b, 0xb3, 0x87, 0x87, 0x1e, 0x7e,
	0x76, 0x08, 0x5e, 0xf9, 0xf1, 0xdd, 0xea, 0xf9, 0x02, 0xc2, 0x25, 0xea, 0xb4, 0xfd, 0xc1, 0xd4,
	0xa2, 0x6d, 0x0c, 0x09, 0xe1, 0x6c, 0xb3, 0xa2, 0x8b, 0x35, 0xa3, 0xeb, 0xab, 0xf5, 0x86, 0x86,
	0x47, 0x64, 0x08, 0x3d, 0xba, 0x99, 0xcf, 0x17, 0x94, 0x86, 0x01, 0x21, 0x30, 0xfa, 0x9c, 0xce,
	0xaf, 0xd6, 0x9f, 0xd2, 0x15, 0xf3, 0x73, 0xe1, 0xf1, 0xbb, 0xfe, 0xd7, 0x53, 0x7f, 0x0d, 0x4d,
	0x76, 0x94, 0x05, 0xdb, 0xb6, 0x7e, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x14, 0x99, 0x4d, 0x1d,
	0xd3, 0x02, 0x00, 0x00,
}
