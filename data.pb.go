// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of QuestType from enums.proto

// Ignoring public import of EncounterType from enums.proto

// Ignoring public import of Filter from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of Costume from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of Form from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of Slot from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of ItemAward from inventory_item.proto

// Ignoring public import of ItemData from inventory_item.proto

// Ignoring public import of ItemType from inventory_item.proto

// Ignoring public import of ItemId from inventory_item.proto

// Ignoring public import of PlayerPublicProfile from data_player.proto

// Ignoring public import of PlayerCurrency from data_player.proto

// Ignoring public import of Currency from data_player.proto

// Ignoring public import of PlayerAvatar from data_player.proto

// Ignoring public import of ContactSettings from data_player.proto

// Ignoring public import of PlayerCamera from data_player.proto

// Ignoring public import of EquippedBadge from data_player.proto

// Ignoring public import of PlayerStats from data_player.proto

// Ignoring public import of DailyBonus from data_player.proto

// Ignoring public import of PlayerAvatarType from data_player.proto

type PokedexEntry struct {
	PokemonId            PokemonId `protobuf:"varint,1,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	TimesEncountered     int32     `protobuf:"varint,2,opt,name=times_encountered" json:"times_encountered,omitempty"`
	TimesCaptured        int32     `protobuf:"varint,3,opt,name=times_captured" json:"times_captured,omitempty"`
	EvolutionStonePieces int32     `protobuf:"varint,4,opt,name=evolution_stone_pieces" json:"evolution_stone_pieces,omitempty"`
	EvolutionStones      int32     `protobuf:"varint,5,opt,name=evolution_stones" json:"evolution_stones,omitempty"`
	CapturedCostumes     []Costume `protobuf:"varint,6,rep,packed,name=captured_costumes,enum=POGOProtos.Enums.Costume" json:"captured_costumes,omitempty"`
	CapturedForms        []Form    `protobuf:"varint,7,rep,packed,name=captured_forms,enum=POGOProtos.Enums.Form" json:"captured_forms,omitempty"`
	CapturedGenders      []Gender  `protobuf:"varint,8,rep,packed,name=captured_genders,enum=POGOProtos.Enums.Gender" json:"captured_genders,omitempty"`
	CapturedShiny        bool      `protobuf:"varint,9,opt,name=captured_shiny" json:"captured_shiny,omitempty"`
}

func (m *PokedexEntry) Reset()                    { *m = PokedexEntry{} }
func (m *PokedexEntry) String() string            { return proto.CompactTextString(m) }
func (*PokedexEntry) ProtoMessage()               {}
func (*PokedexEntry) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

type DownloadUrlEntry struct {
	AssetId  string `protobuf:"bytes,1,opt,name=asset_id" json:"asset_id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Checksum uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *DownloadUrlEntry) Reset()                    { *m = DownloadUrlEntry{} }
func (m *DownloadUrlEntry) String() string            { return proto.CompactTextString(m) }
func (*DownloadUrlEntry) ProtoMessage()               {}
func (*DownloadUrlEntry) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{1} }

type PokemonDisplay struct {
	Costume Costume `protobuf:"varint,1,opt,name=costume,enum=POGOProtos.Enums.Costume" json:"costume,omitempty"`
	Gender  Gender  `protobuf:"varint,2,opt,name=gender,enum=POGOProtos.Enums.Gender" json:"gender,omitempty"`
	Shiny   bool    `protobuf:"varint,3,opt,name=shiny" json:"shiny,omitempty"`
	Form    Form    `protobuf:"varint,4,opt,name=form,enum=POGOProtos.Enums.Form" json:"form,omitempty"`
}

func (m *PokemonDisplay) Reset()                    { *m = PokemonDisplay{} }
func (m *PokemonDisplay) String() string            { return proto.CompactTextString(m) }
func (*PokemonDisplay) ProtoMessage()               {}
func (*PokemonDisplay) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{2} }

type PlayerBadge struct {
	BadgeType    BadgeType `protobuf:"varint,1,opt,name=badge_type,enum=POGOProtos.Enums.BadgeType" json:"badge_type,omitempty"`
	Rank         int32     `protobuf:"varint,2,opt,name=rank" json:"rank,omitempty"`
	StartValue   int32     `protobuf:"varint,3,opt,name=start_value" json:"start_value,omitempty"`
	EndValue     int32     `protobuf:"varint,4,opt,name=end_value" json:"end_value,omitempty"`
	CurrentValue float64   `protobuf:"fixed64,5,opt,name=current_value" json:"current_value,omitempty"`
}

func (m *PlayerBadge) Reset()                    { *m = PlayerBadge{} }
func (m *PlayerBadge) String() string            { return proto.CompactTextString(m) }
func (*PlayerBadge) ProtoMessage()               {}
func (*PlayerBadge) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{3} }

type ClientVersion struct {
	MinVersion string `protobuf:"bytes,1,opt,name=min_version" json:"min_version,omitempty"`
}

func (m *ClientVersion) Reset()                    { *m = ClientVersion{} }
func (m *ClientVersion) String() string            { return proto.CompactTextString(m) }
func (*ClientVersion) ProtoMessage()               {}
func (*ClientVersion) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{4} }

type PokemonData struct {
	Id                     uint64          `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	PokemonId              PokemonId       `protobuf:"varint,2,opt,name=pokemon_id,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	Cp                     int32           `protobuf:"varint,3,opt,name=cp" json:"cp,omitempty"`
	Stamina                int32           `protobuf:"varint,4,opt,name=stamina" json:"stamina,omitempty"`
	StaminaMax             int32           `protobuf:"varint,5,opt,name=stamina_max" json:"stamina_max,omitempty"`
	Move_1                 PokemonMove     `protobuf:"varint,6,opt,name=move_1,enum=POGOProtos.Enums.PokemonMove" json:"move_1,omitempty"`
	Move_2                 PokemonMove     `protobuf:"varint,7,opt,name=move_2,enum=POGOProtos.Enums.PokemonMove" json:"move_2,omitempty"`
	DeployedFortId         string          `protobuf:"bytes,8,opt,name=deployed_fort_id" json:"deployed_fort_id,omitempty"`
	OwnerName              string          `protobuf:"bytes,9,opt,name=owner_name" json:"owner_name,omitempty"`
	IsEgg                  bool            `protobuf:"varint,10,opt,name=is_egg" json:"is_egg,omitempty"`
	EggKmWalkedTarget      float64         `protobuf:"fixed64,11,opt,name=egg_km_walked_target" json:"egg_km_walked_target,omitempty"`
	EggKmWalkedStart       float64         `protobuf:"fixed64,12,opt,name=egg_km_walked_start" json:"egg_km_walked_start,omitempty"`
	Origin                 int32           `protobuf:"varint,14,opt,name=origin" json:"origin,omitempty"`
	HeightM                float32         `protobuf:"fixed32,15,opt,name=height_m" json:"height_m,omitempty"`
	WeightKg               float32         `protobuf:"fixed32,16,opt,name=weight_kg" json:"weight_kg,omitempty"`
	IndividualAttack       int32           `protobuf:"varint,17,opt,name=individual_attack" json:"individual_attack,omitempty"`
	IndividualDefense      int32           `protobuf:"varint,18,opt,name=individual_defense" json:"individual_defense,omitempty"`
	IndividualStamina      int32           `protobuf:"varint,19,opt,name=individual_stamina" json:"individual_stamina,omitempty"`
	CpMultiplier           float32         `protobuf:"fixed32,20,opt,name=cp_multiplier" json:"cp_multiplier,omitempty"`
	Pokeball               ItemId          `protobuf:"varint,21,opt,name=pokeball,enum=POGOProtos.Inventory.Item.ItemId" json:"pokeball,omitempty"`
	CapturedCellId         uint64          `protobuf:"varint,22,opt,name=captured_cell_id" json:"captured_cell_id,omitempty"`
	BattlesAttacked        int32           `protobuf:"varint,23,opt,name=battles_attacked" json:"battles_attacked,omitempty"`
	BattlesDefended        int32           `protobuf:"varint,24,opt,name=battles_defended" json:"battles_defended,omitempty"`
	EggIncubatorId         string          `protobuf:"bytes,25,opt,name=egg_incubator_id" json:"egg_incubator_id,omitempty"`
	CreationTimeMs         uint64          `protobuf:"varint,26,opt,name=creation_time_ms" json:"creation_time_ms,omitempty"`
	NumUpgrades            int32           `protobuf:"varint,27,opt,name=num_upgrades" json:"num_upgrades,omitempty"`
	AdditionalCpMultiplier float32         `protobuf:"fixed32,28,opt,name=additional_cp_multiplier" json:"additional_cp_multiplier,omitempty"`
	Favorite               int32           `protobuf:"varint,29,opt,name=favorite" json:"favorite,omitempty"`
	Nickname               string          `protobuf:"bytes,30,opt,name=nickname" json:"nickname,omitempty"`
	FromFort               int32           `protobuf:"varint,31,opt,name=from_fort" json:"from_fort,omitempty"`
	BuddyCandyAwarded      int32           `protobuf:"varint,32,opt,name=buddy_candy_awarded" json:"buddy_candy_awarded,omitempty"`
	BuddyTotalKmWalked     float32         `protobuf:"fixed32,33,opt,name=buddy_total_km_walked" json:"buddy_total_km_walked,omitempty"`
	DisplayPokemonId       int32           `protobuf:"varint,34,opt,name=display_pokemon_id" json:"display_pokemon_id,omitempty"`
	DisplayCp              int32           `protobuf:"varint,35,opt,name=display_cp" json:"display_cp,omitempty"`
	PokemonDisplay         *PokemonDisplay `protobuf:"bytes,36,opt,name=pokemon_display" json:"pokemon_display,omitempty"`
}

func (m *PokemonData) Reset()                    { *m = PokemonData{} }
func (m *PokemonData) String() string            { return proto.CompactTextString(m) }
func (*PokemonData) ProtoMessage()               {}
func (*PokemonData) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{5} }

func (m *PokemonData) GetPokemonDisplay() *PokemonDisplay {
	if m != nil {
		return m.PokemonDisplay
	}
	return nil
}

type PlayerData struct {
	CreationTimestampMs     int64            `protobuf:"varint,1,opt,name=creation_timestamp_ms" json:"creation_timestamp_ms,omitempty"`
	Username                string           `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Team                    TeamColor        `protobuf:"varint,5,opt,name=team,enum=POGOProtos.Enums.TeamColor" json:"team,omitempty"`
	TutorialState           []TutorialState  `protobuf:"varint,7,rep,packed,name=tutorial_state,enum=POGOProtos.Enums.TutorialState" json:"tutorial_state,omitempty"`
	Avatar                  *PlayerAvatar    `protobuf:"bytes,8,opt,name=avatar" json:"avatar,omitempty"`
	MaxPokemonStorage       int32            `protobuf:"varint,9,opt,name=max_pokemon_storage" json:"max_pokemon_storage,omitempty"`
	MaxItemStorage          int32            `protobuf:"varint,10,opt,name=max_item_storage" json:"max_item_storage,omitempty"`
	DailyBonus              *DailyBonus      `protobuf:"bytes,11,opt,name=daily_bonus" json:"daily_bonus,omitempty"`
	EquippedBadge           *EquippedBadge   `protobuf:"bytes,12,opt,name=equipped_badge" json:"equipped_badge,omitempty"`
	ContactSettings         *ContactSettings `protobuf:"bytes,13,opt,name=contact_settings" json:"contact_settings,omitempty"`
	Currencies              []*Currency      `protobuf:"bytes,14,rep,name=currencies" json:"currencies,omitempty"`
	RemainingCodenameClaims int32            `protobuf:"varint,15,opt,name=remaining_codename_claims" json:"remaining_codename_claims,omitempty"`
	BuddyPokemon            *BuddyPokemon    `protobuf:"bytes,16,opt,name=buddy_pokemon" json:"buddy_pokemon,omitempty"`
	BattleLockoutEndMs      int64            `protobuf:"varint,17,opt,name=battle_lockout_end_ms" json:"battle_lockout_end_ms,omitempty"`
	SecondaryPlayerAvatar   *PlayerAvatar    `protobuf:"bytes,18,opt,name=secondary_player_avatar" json:"secondary_player_avatar,omitempty"`
}

func (m *PlayerData) Reset()                    { *m = PlayerData{} }
func (m *PlayerData) String() string            { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()               {}
func (*PlayerData) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{6} }

func (m *PlayerData) GetAvatar() *PlayerAvatar {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *PlayerData) GetDailyBonus() *DailyBonus {
	if m != nil {
		return m.DailyBonus
	}
	return nil
}

func (m *PlayerData) GetEquippedBadge() *EquippedBadge {
	if m != nil {
		return m.EquippedBadge
	}
	return nil
}

func (m *PlayerData) GetContactSettings() *ContactSettings {
	if m != nil {
		return m.ContactSettings
	}
	return nil
}

func (m *PlayerData) GetCurrencies() []*Currency {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func (m *PlayerData) GetBuddyPokemon() *BuddyPokemon {
	if m != nil {
		return m.BuddyPokemon
	}
	return nil
}

func (m *PlayerData) GetSecondaryPlayerAvatar() *PlayerAvatar {
	if m != nil {
		return m.SecondaryPlayerAvatar
	}
	return nil
}

type BackgroundToken struct {
	Token          []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpirationTime int64  `protobuf:"varint,2,opt,name=expiration_time" json:"expiration_time,omitempty"`
	Iv             []byte `protobuf:"bytes,3,opt,name=iv,proto3" json:"iv,omitempty"`
}

func (m *BackgroundToken) Reset()                    { *m = BackgroundToken{} }
func (m *BackgroundToken) String() string            { return proto.CompactTextString(m) }
func (*BackgroundToken) ProtoMessage()               {}
func (*BackgroundToken) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{7} }

type AssetDigestEntry struct {
	AssetId    string `protobuf:"bytes,1,opt,name=asset_id" json:"asset_id,omitempty"`
	BundleName string `protobuf:"bytes,2,opt,name=bundle_name" json:"bundle_name,omitempty"`
	Version    int64  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Checksum   uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
	Size       int32  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Key        []byte `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *AssetDigestEntry) Reset()                    { *m = AssetDigestEntry{} }
func (m *AssetDigestEntry) String() string            { return proto.CompactTextString(m) }
func (*AssetDigestEntry) ProtoMessage()               {}
func (*AssetDigestEntry) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{8} }

type BuddyPokemon struct {
	Id            uint64  `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	StartKmWalked float64 `protobuf:"fixed64,2,opt,name=start_km_walked" json:"start_km_walked,omitempty"`
	LastKmAwarded float64 `protobuf:"fixed64,3,opt,name=last_km_awarded" json:"last_km_awarded,omitempty"`
}

func (m *BuddyPokemon) Reset()                    { *m = BuddyPokemon{} }
func (m *BuddyPokemon) String() string            { return proto.CompactTextString(m) }
func (*BuddyPokemon) ProtoMessage()               {}
func (*BuddyPokemon) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{9} }

func init() {
	proto.RegisterType((*PokedexEntry)(nil), "POGOProtos.Data.PokedexEntry")
	proto.RegisterType((*DownloadUrlEntry)(nil), "POGOProtos.Data.DownloadUrlEntry")
	proto.RegisterType((*PokemonDisplay)(nil), "POGOProtos.Data.PokemonDisplay")
	proto.RegisterType((*PlayerBadge)(nil), "POGOProtos.Data.PlayerBadge")
	proto.RegisterType((*ClientVersion)(nil), "POGOProtos.Data.ClientVersion")
	proto.RegisterType((*PokemonData)(nil), "POGOProtos.Data.PokemonData")
	proto.RegisterType((*PlayerData)(nil), "POGOProtos.Data.PlayerData")
	proto.RegisterType((*BackgroundToken)(nil), "POGOProtos.Data.BackgroundToken")
	proto.RegisterType((*AssetDigestEntry)(nil), "POGOProtos.Data.AssetDigestEntry")
	proto.RegisterType((*BuddyPokemon)(nil), "POGOProtos.Data.BuddyPokemon")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 1346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x6e, 0x1b, 0xbb,
	0x11, 0xae, 0x24, 0x5b, 0x96, 0x47, 0xb2, 0x24, 0xaf, 0xff, 0x68, 0x3b, 0x4e, 0x14, 0xd5, 0x45,
	0xd5, 0x02, 0x75, 0x51, 0x25, 0x40, 0x0b, 0x14, 0xbd, 0x88, 0x7f, 0x1a, 0x18, 0x68, 0x11, 0x15,
	0x49, 0x7b, 0xd1, 0x1b, 0x82, 0xda, 0x1d, 0xaf, 0x09, 0xed, 0x92, 0x5b, 0x92, 0x2b, 0x5b, 0x7d,
	0x85, 0xf3, 0x08, 0xe7, 0x5d, 0xce, 0xc5, 0x79, 0x90, 0xf3, 0x2c, 0x07, 0x9c, 0xdd, 0x95, 0xed,
	0x38, 0xca, 0xc9, 0x8d, 0xbd, 0xe2, 0x7c, 0x43, 0x72, 0xbe, 0xf9, 0x66, 0x86, 0x00, 0x91, 0x70,
	0xe2, 0x2c, 0x33, 0xda, 0xe9, 0xa0, 0x37, 0xf9, 0xf0, 0xfe, 0xc3, 0xc4, 0x7f, 0xda, 0xb3, 0x4b,
	0xe1, 0xc4, 0x51, 0x1b, 0x55, 0x9e, 0xda, 0xc2, 0x7a, 0xb4, 0x2b, 0xd5, 0x1c, 0x95, 0xd3, 0x66,
	0xc1, 0xa5, 0xc3, 0xb4, 0x5c, 0xdd, 0xf6, 0xfe, 0x3c, 0x4b, 0xc4, 0x02, 0x4d, 0xb1, 0x34, 0xfc,
	0xa9, 0x0e, 0x9d, 0x89, 0x9e, 0x61, 0x84, 0xf7, 0x57, 0xca, 0x99, 0x45, 0xf0, 0x47, 0x80, 0x4c,
	0xcf, 0x30, 0xd5, 0x8a, 0xcb, 0x88, 0xd5, 0x06, 0xb5, 0x51, 0x77, 0x7c, 0x7c, 0xf6, 0xe8, 0xb0,
	0x2b, 0x3a, 0x66, 0x52, 0x60, 0xae, 0xa3, 0xe0, 0x10, 0xb6, 0x9d, 0x4c, 0xd1, 0x72, 0x54, 0xa1,
	0xce, 0x95, 0x43, 0x83, 0x11, 0xab, 0x0f, 0x6a, 0xa3, 0xf5, 0x60, 0x1f, 0xba, 0x85, 0x29, 0x14,
	0x99, 0xcb, 0xfd, 0x7a, 0x83, 0xd6, 0x5f, 0xc2, 0x3e, 0xce, 0x75, 0x92, 0x3b, 0xa9, 0x15, 0xb7,
	0x4e, 0x2b, 0xe4, 0x99, 0xc4, 0x10, 0x2d, 0x5b, 0x23, 0x3b, 0x83, 0xfe, 0x67, 0x76, 0xcb, 0xd6,
	0xc9, 0xf2, 0x16, 0xb6, 0xab, 0xbd, 0x78, 0xa8, 0xad, 0xcb, 0x53, 0xb4, 0xac, 0x39, 0x68, 0x8c,
	0xba, 0xe3, 0xc3, 0xe7, 0x97, 0xbc, 0x28, 0x10, 0xc1, 0x19, 0x74, 0x97, 0x5e, 0x37, 0xda, 0xa4,
	0x96, 0x6d, 0x90, 0xcb, 0xfe, 0x73, 0x97, 0xbf, 0x6b, 0x93, 0x06, 0x63, 0xe8, 0x2f, 0xf1, 0x31,
	0xaa, 0x08, 0x8d, 0x65, 0x2d, 0xf2, 0x60, 0xcf, 0x3d, 0xde, 0x13, 0xc0, 0xc7, 0xba, 0xf4, 0xb1,
	0xb7, 0x52, 0x2d, 0xd8, 0xe6, 0xa0, 0x36, 0x6a, 0x0d, 0xff, 0x05, 0xfd, 0x4b, 0x7d, 0xa7, 0x12,
	0x2d, 0xa2, 0x7f, 0x9b, 0xa4, 0xe0, 0xb8, 0x0f, 0x2d, 0x61, 0x2d, 0xba, 0x8a, 0xe1, 0xcd, 0xa0,
	0x0d, 0x8d, 0xdc, 0x24, 0x44, 0xdb, 0x66, 0xd0, 0x81, 0x35, 0x2b, 0xff, 0x8f, 0x25, 0x59, 0x7d,
	0x68, 0x85, 0xb7, 0x18, 0xce, 0x6c, 0x9e, 0x12, 0x3d, 0x1b, 0xc3, 0xef, 0x6b, 0xd0, 0x2d, 0xf9,
	0xbf, 0x94, 0xd6, 0xa7, 0x33, 0xf8, 0x3d, 0x6c, 0x94, 0x74, 0x94, 0x29, 0xfb, 0x0a, 0x1b, 0x23,
	0x68, 0x16, 0x41, 0xd1, 0x71, 0x5f, 0x8b, 0x69, 0x0b, 0xd6, 0x8b, 0x50, 0xfc, 0x4d, 0x5a, 0xc1,
	0x29, 0xac, 0x79, 0xf6, 0xe8, 0x16, 0x2b, 0xc9, 0x1b, 0x7e, 0x57, 0x83, 0xf6, 0x84, 0x24, 0x76,
	0x2e, 0xa2, 0x18, 0xbd, 0xa0, 0xa6, 0xfe, 0x83, 0xbb, 0x45, 0x86, 0xab, 0x05, 0x45, 0xe0, 0x4f,
	0x8b, 0x0c, 0x7d, 0xf8, 0x46, 0xa8, 0x59, 0xa9, 0xa1, 0x1d, 0x68, 0x5b, 0x27, 0x8c, 0xe3, 0x73,
	0x91, 0xe4, 0x15, 0x27, 0xdb, 0xb0, 0x89, 0x2a, 0x2a, 0x97, 0x0a, 0xcd, 0xec, 0xc1, 0x56, 0x98,
	0x1b, 0x83, 0xaa, 0x42, 0x7a, 0xc1, 0xd4, 0x86, 0xa7, 0xb0, 0x75, 0x91, 0x48, 0x54, 0xee, 0x3f,
	0x68, 0xac, 0xd4, 0xca, 0xef, 0x97, 0x4a, 0xc5, 0xe7, 0xc5, 0xcf, 0x82, 0xfe, 0xe1, 0x0f, 0x1b,
	0xd0, 0xae, 0x18, 0x15, 0x4e, 0x04, 0x00, 0xf5, 0x32, 0x35, 0xcd, 0xcf, 0x0a, 0xa2, 0xfe, 0xcb,
	0x05, 0x01, 0x50, 0x0f, 0xb3, 0xf2, 0xa2, 0x3d, 0xd8, 0xb0, 0x4e, 0xa4, 0x52, 0x89, 0xf2, 0x9a,
	0x45, 0x38, 0x7e, 0x81, 0xa7, 0xe2, 0xbe, 0x54, 0xf5, 0x1f, 0xa0, 0x99, 0xea, 0x39, 0xf2, 0x3f,
	0xb1, 0x26, 0x6d, 0x7f, 0xb2, 0x72, 0xfb, 0x7f, 0xea, 0x39, 0x2e, 0xe1, 0x63, 0xb6, 0xf1, 0x2d,
	0x70, 0x06, 0xfd, 0x08, 0xb3, 0x44, 0x2f, 0x0a, 0xf5, 0x93, 0xea, 0x5a, 0x24, 0xb4, 0x00, 0x40,
	0xdf, 0x29, 0x34, 0x5c, 0x89, 0x14, 0x49, 0xaf, 0x9b, 0x41, 0x17, 0x9a, 0xd2, 0x72, 0x8c, 0x63,
	0x06, 0x94, 0xf4, 0x17, 0xb0, 0x8b, 0x71, 0xcc, 0x67, 0x29, 0xbf, 0x13, 0xc9, 0x0c, 0x23, 0xee,
	0x84, 0x89, 0xd1, 0xb1, 0xb6, 0xa7, 0x37, 0x38, 0x86, 0x9d, 0xa7, 0x56, 0xca, 0x15, 0xeb, 0x90,
	0xb1, 0x0b, 0x4d, 0x6d, 0x64, 0x2c, 0x15, 0xeb, 0x56, 0x4a, 0xbe, 0x45, 0x19, 0xdf, 0x3a, 0x9e,
	0xb2, 0xde, 0xa0, 0x36, 0xaa, 0xfb, 0x3c, 0xde, 0x15, 0x2b, 0xb3, 0x98, 0xf5, 0x69, 0xe9, 0x10,
	0xb6, 0xa5, 0x8a, 0xe4, 0x5c, 0x46, 0xb9, 0x48, 0xb8, 0x70, 0x4e, 0x84, 0x33, 0xb6, 0x4d, 0xfe,
	0x47, 0x10, 0x3c, 0x32, 0x45, 0x78, 0x83, 0xca, 0x22, 0x0b, 0xbe, 0x60, 0xab, 0x38, 0xdf, 0x59,
	0x4a, 0x23, 0xe3, 0x69, 0x9e, 0x38, 0x99, 0x25, 0x12, 0x0d, 0xdb, 0xa5, 0x93, 0xde, 0x40, 0xcb,
	0x27, 0x76, 0x2a, 0x92, 0x84, 0xed, 0x11, 0x91, 0xaf, 0x1f, 0x13, 0x79, 0x5d, 0x75, 0xd0, 0xb3,
	0x6b, 0xdf, 0x41, 0xfd, 0x9f, 0xeb, 0xc8, 0x93, 0xf9, 0xd0, 0x80, 0x30, 0x49, 0x3c, 0x99, 0xfb,
	0x83, 0xda, 0x68, 0xcd, 0x5b, 0xa6, 0xc2, 0xb9, 0x04, 0x6d, 0x79, 0x6b, 0x8c, 0xd8, 0x41, 0xd5,
	0xce, 0x2a, 0x0b, 0x5d, 0x3a, 0xc2, 0x88, 0xb1, 0x65, 0xa3, 0x8b, 0x63, 0x2e, 0x55, 0x98, 0x4f,
	0x85, 0xd3, 0xc6, 0xef, 0x76, 0x48, 0x69, 0xf0, 0xe7, 0x18, 0x14, 0xd4, 0x01, 0x7d, 0x0f, 0xe5,
	0xa9, 0x65, 0x47, 0x74, 0xce, 0x2e, 0x74, 0x54, 0x9e, 0xf2, 0x3c, 0x8b, 0x8d, 0x88, 0xd0, 0xb2,
	0x63, 0xda, 0x69, 0x00, 0x4c, 0x44, 0x91, 0xf4, 0x78, 0x91, 0xf0, 0xa7, 0xe1, 0xbe, 0xa0, 0x70,
	0xfb, 0xd0, 0xba, 0x11, 0x73, 0x6d, 0xa4, 0x43, 0x76, 0x52, 0xe5, 0x43, 0xc9, 0x70, 0x46, 0xc9,
	0x7f, 0x49, 0xa7, 0x6e, 0xc3, 0xe6, 0x8d, 0xd1, 0x29, 0xc9, 0x84, 0xbd, 0x22, 0xd0, 0x31, 0xec,
	0x4c, 0xf3, 0x28, 0x5a, 0xf0, 0x50, 0xa8, 0x68, 0xc1, 0xc5, 0x9d, 0x30, 0xfe, 0xfe, 0x03, 0x32,
	0x9e, 0xc0, 0x5e, 0x61, 0x74, 0xda, 0x89, 0xe4, 0x41, 0x06, 0xec, 0x35, 0x1d, 0x79, 0x04, 0x41,
	0x54, 0x34, 0x28, 0xfe, 0xa8, 0x84, 0x86, 0xe4, 0x1a, 0x00, 0x54, 0xb6, 0x30, 0x63, 0xbf, 0xa6,
	0xb5, 0xbf, 0x40, 0xaf, 0xc2, 0x95, 0x36, 0x76, 0x3a, 0xa8, 0x8d, 0xda, 0xe3, 0x57, 0x67, 0x9f,
	0x4d, 0xbb, 0xb3, 0xa7, 0xfd, 0x6f, 0xf8, 0xe3, 0x3a, 0x40, 0xd1, 0x74, 0xa8, 0x7e, 0x4f, 0x60,
	0xef, 0x09, 0x7b, 0x5e, 0x0f, 0x99, 0xa7, 0xd0, 0x97, 0x74, 0xc3, 0x07, 0x9e, 0x5b, 0x34, 0x14,
	0x78, 0xd1, 0x72, 0x7f, 0x07, 0x6b, 0x0e, 0x45, 0x4a, 0xf5, 0xf8, 0xc5, 0xf2, 0xfe, 0x84, 0x22,
	0xbd, 0xd0, 0x89, 0x36, 0xc1, 0x5f, 0xa1, 0xeb, 0x72, 0xa7, 0x8d, 0x2c, 0x74, 0xe6, 0xb0, 0x1c,
	0x26, 0xaf, 0xbe, 0xe0, 0x54, 0xe2, 0x3e, 0x7a, 0xd8, 0x79, 0xbd, 0x5f, 0x0b, 0xde, 0x42, 0x53,
	0xcc, 0x85, 0x13, 0x86, 0x2a, 0xb0, 0x3d, 0x3e, 0x7d, 0x1e, 0x58, 0x31, 0x9d, 0x8b, 0x7f, 0xef,
	0x08, 0xeb, 0x73, 0x90, 0x8a, 0xfb, 0x25, 0x87, 0xd6, 0x69, 0x23, 0xe2, 0xa2, 0x60, 0x49, 0x43,
	0xde, 0xe8, 0xc7, 0xfc, 0xd2, 0x02, 0x64, 0xf9, 0x33, 0xb4, 0x23, 0x21, 0x93, 0x05, 0x9f, 0x6a,
	0x95, 0x5b, 0xaa, 0xd8, 0xf6, 0x78, 0xb8, 0xea, 0xc4, 0x4b, 0x0f, 0x3d, 0xf7, 0xc8, 0xe0, 0x6f,
	0xd0, 0xc5, 0xff, 0xe5, 0x32, 0xcb, 0x30, 0xe2, 0xd4, 0xbb, 0xa9, 0xa0, 0xdb, 0xe3, 0xdf, 0xac,
	0xf2, 0xbd, 0x2a, 0xd1, 0x45, 0xc7, 0x7f, 0x07, 0xfd, 0x50, 0x2b, 0x27, 0x42, 0xc7, 0x2d, 0x3a,
	0x27, 0x55, 0x6c, 0xd9, 0x16, 0x6d, 0xf0, 0xdb, 0x55, 0x1b, 0x5c, 0x14, 0xf8, 0x8f, 0x25, 0x3c,
	0x78, 0x0b, 0x50, 0x74, 0xf3, 0x50, 0xa2, 0x65, 0xdd, 0x41, 0x63, 0xd4, 0x1e, 0x0f, 0x56, 0x3a,
	0x17, 0xc8, 0x45, 0xf0, 0x1a, 0x0e, 0x0d, 0xa6, 0x42, 0x2a, 0xa9, 0x62, 0x1e, 0xea, 0x08, 0x7d,
	0x86, 0x79, 0x98, 0x08, 0x99, 0x5a, 0xea, 0x38, 0xfe, 0x01, 0xb1, 0x55, 0x28, 0xb6, 0x24, 0x93,
	0xba, 0x4e, 0xfb, 0x69, 0x0b, 0xa5, 0xbd, 0xcf, 0x3d, 0xaa, 0x54, 0x19, 0xe9, 0x9c, 0x2a, 0x98,
	0x27, 0x3a, 0x9c, 0xe9, 0xdc, 0x71, 0x3f, 0x7e, 0x52, 0x4b, 0x8d, 0xa9, 0x11, 0x5c, 0xc1, 0x81,
	0xc5, 0x50, 0xab, 0x48, 0x98, 0x45, 0xf9, 0xbc, 0xe2, 0x65, 0x9a, 0x83, 0x6f, 0x4f, 0xf3, 0xf0,
	0x0a, 0x7a, 0xe7, 0x22, 0x9c, 0xc5, 0x46, 0xe7, 0x2a, 0xfa, 0xa4, 0x67, 0xa8, 0xfc, 0x04, 0x76,
	0xfe, 0x83, 0x84, 0xdb, 0x09, 0x0e, 0xa0, 0x87, 0xf7, 0x99, 0x34, 0x0f, 0xca, 0x26, 0xfd, 0x36,
	0x68, 0x60, 0xcd, 0x69, 0xe6, 0x74, 0x86, 0x77, 0xd0, 0x7f, 0xe7, 0x5f, 0x17, 0x97, 0x32, 0x46,
	0xeb, 0x56, 0xbd, 0x38, 0x76, 0xa0, 0x3d, 0xcd, 0x55, 0x94, 0x20, 0x7f, 0x54, 0x06, 0x3d, 0xd8,
	0xa8, 0x06, 0x63, 0xa3, 0xaa, 0x94, 0xa7, 0x8f, 0x8f, 0xe5, 0xe3, 0xa4, 0x98, 0x5c, 0x6d, 0x68,
	0xcc, 0x70, 0x41, 0x63, 0xab, 0x33, 0xfc, 0x07, 0x74, 0x9e, 0xb0, 0xf6, 0x78, 0x8a, 0x1e, 0x40,
	0xaf, 0x18, 0xe3, 0x0f, 0x3d, 0xa2, 0x4e, 0x43, 0xe2, 0x00, 0x7a, 0x89, 0xb0, 0xb4, 0x5e, 0xf5,
	0x16, 0x7f, 0x74, 0xed, 0xbc, 0xf5, 0xdf, 0x26, 0x3d, 0x51, 0xed, 0xe4, 0x57, 0x93, 0xda, 0xa4,
	0x3e, 0x2d, 0x7e, 0xbd, 0xf9, 0x39, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xe5, 0x14, 0xdf, 0x05, 0x0b,
	0x00, 0x00,
}