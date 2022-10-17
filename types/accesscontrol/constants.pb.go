// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/constants.proto

package accesscontrol

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AccessType int32

const (
	AccessType_UNKNOWN AccessType = 0
	AccessType_READ    AccessType = 1
	AccessType_WRITE   AccessType = 2
	AccessType_COMMIT  AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "UNKNOWN",
	1: "READ",
	2: "WRITE",
	3: "COMMIT",
}

var AccessType_value = map[string]int32{
	"UNKNOWN": 0,
	"READ":    1,
	"WRITE":   2,
	"COMMIT":  3,
}

func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}

func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{0}
}

type ResourceType int32

const (
	ResourceType_ANY        ResourceType = 0
	ResourceType_KV         ResourceType = 1
	ResourceType_Mem        ResourceType = 2
	ResourceType_DexMem     ResourceType = 3
	ResourceType_KV_BANK    ResourceType = 4
	ResourceType_KV_STAKING ResourceType = 5
	ResourceType_KV_WASM    ResourceType = 6
	ResourceType_KV_ORACLE  ResourceType = 7
	ResourceType_KV_DEX     ResourceType = 8
	ResourceType_KV_EPOCH   ResourceType = 9
)

var ResourceType_name = map[int32]string{
	0: "ANY",
	1: "KV",
	2: "Mem",
	3: "DexMem",
	4: "KV_BANK",
	5: "KV_STAKING",
	6: "KV_WASM",
	7: "KV_ORACLE",
	8: "KV_DEX",
	9: "KV_EPOCH",
}

var ResourceType_value = map[string]int32{
	"ANY":        0,
	"KV":         1,
	"Mem":        2,
	"DexMem":     3,
	"KV_BANK":    4,
	"KV_STAKING": 5,
	"KV_WASM":    6,
	"KV_ORACLE":  7,
	"KV_DEX":     8,
	"KV_EPOCH":   9,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x5b, 0xfe, 0x14, 0x18, 0xd1, 0x4c, 0xf6, 0x6c, 0x7a, 0xf2, 0x44, 0x62, 0x2b, 0xf1,
	0xe6, 0x6d, 0x81, 0x46, 0x71, 0x6d, 0x6b, 0x4a, 0x5d, 0xd4, 0x4b, 0x03, 0xeb, 0x46, 0x8d, 0xc2,
	0x12, 0x76, 0x31, 0xf2, 0x09, 0xbc, 0xfa, 0xb1, 0x3c, 0x72, 0xf4, 0x68, 0xe0, 0x8b, 0x98, 0x52,
	0x2e, 0x7a, 0x9a, 0xc9, 0xbc, 0xf7, 0xe6, 0x25, 0x3f, 0x38, 0x12, 0x4a, 0x4f, 0x94, 0xf6, 0x47,
	0x42, 0x48, 0xad, 0x85, 0x9a, 0x9a, 0xb9, 0x7a, 0xf5, 0x85, 0x9a, 0x6a, 0x33, 0x9a, 0x1a, 0xed,
	0xcd, 0xe6, 0xca, 0x28, 0x72, 0x58, 0xb8, 0xbc, 0x3f, 0x2e, 0xef, 0xad, 0x3d, 0x96, 0x66, 0xd4,
	0x6e, 0x9d, 0x01, 0xd0, 0xad, 0x90, 0x2e, 0x67, 0x92, 0xec, 0x41, 0xed, 0x26, 0x62, 0x51, 0x3c,
	0x8c, 0xd0, 0x22, 0x75, 0xa8, 0x24, 0x01, 0xed, 0xa1, 0x4d, 0x1a, 0x50, 0x1d, 0x26, 0xfd, 0x34,
	0xc0, 0x12, 0x01, 0x70, 0xba, 0x71, 0x18, 0xf6, 0x53, 0x2c, 0xb7, 0x3e, 0x6c, 0x68, 0x26, 0x52,
	0xab, 0xc5, 0x5c, 0xc8, 0x6d, 0xbc, 0x06, 0x65, 0x1a, 0xdd, 0xa1, 0x45, 0x1c, 0x28, 0x31, 0x8e,
	0x76, 0x7e, 0x08, 0xe5, 0xa4, 0x88, 0xf5, 0xe4, 0x7b, 0xbe, 0x97, 0xf3, 0x12, 0xc6, 0xb3, 0x0e,
	0x8d, 0x18, 0x56, 0xc8, 0x01, 0x00, 0xe3, 0xd9, 0x20, 0xa5, 0xac, 0x1f, 0x9d, 0x63, 0x75, 0x27,
	0x0e, 0xe9, 0x20, 0x44, 0x87, 0xec, 0x43, 0x83, 0xf1, 0x2c, 0x4e, 0x68, 0xf7, 0x2a, 0xc0, 0x5a,
	0xfe, 0x84, 0xf1, 0xac, 0x17, 0xdc, 0x62, 0x9d, 0x34, 0xa1, 0xce, 0x78, 0x16, 0x5c, 0xc7, 0xdd,
	0x0b, 0x6c, 0x74, 0x2e, 0xbf, 0xd6, 0xae, 0xbd, 0x5a, 0xbb, 0xf6, 0xcf, 0xda, 0xb5, 0x3f, 0x37,
	0xae, 0xb5, 0xda, 0xb8, 0xd6, 0xf7, 0xc6, 0xb5, 0xee, 0x4f, 0x1e, 0x9f, 0xcd, 0xd3, 0x62, 0xec,
	0x09, 0x35, 0xf1, 0x77, 0xb8, 0x8a, 0x71, 0xac, 0x1f, 0x5e, 0x7c, 0xb3, 0x9c, 0xc9, 0x7f, 0xfc,
	0xc6, 0xce, 0x16, 0xdb, 0xe9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0x3c, 0xaa, 0x30, 0x5e,
	0x01, 0x00, 0x00,
}
