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

type AccessOperationSelectorType int32

const (
	AccessOperationSelectorType_NONE AccessOperationSelectorType = 0
	AccessOperationSelectorType_JQ   AccessOperationSelectorType = 1
)

var AccessOperationSelectorType_name = map[int32]string{
	0: "NONE",
	1: "JQ",
}

var AccessOperationSelectorType_value = map[string]int32{
	"NONE": 0,
	"JQ":   1,
}

func (x AccessOperationSelectorType) String() string {
	return proto.EnumName(AccessOperationSelectorType_name, int32(x))
}

func (AccessOperationSelectorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

type ResourceType int32

const (
	ResourceType_ANY                                     ResourceType = 0
	ResourceType_KV                                      ResourceType = 1
	ResourceType_Mem                                     ResourceType = 2
	ResourceType_DexMem                                  ResourceType = 3
	ResourceType_KV_BANK                                 ResourceType = 4
	ResourceType_KV_STAKING                              ResourceType = 5
	ResourceType_KV_WASM                                 ResourceType = 6
	ResourceType_KV_ORACLE                               ResourceType = 7
	ResourceType_KV_DEX                                  ResourceType = 8
	ResourceType_KV_EPOCH                                ResourceType = 9
	ResourceType_KV_TOKENFACTORY                         ResourceType = 10
	ResourceType_KV_ORACLE_VOTE_TARGETS                  ResourceType = 11
	ResourceType_KV_ORACLE_AGGREGATE_VOTES               ResourceType = 12
	ResourceType_KV_ORACLE_FEEDERS                       ResourceType = 13
	ResourceType_KV_STAKING_DELEGATION                   ResourceType = 14
	ResourceType_KV_STAKING_VALIDATOR                    ResourceType = 15
	ResourceType_KV_AUTH                                 ResourceType = 16
	ResourceType_KV_AUTH_ADDRESS_STORE                   ResourceType = 17
	ResourceType_KV_BANK_SUPPLY                          ResourceType = 18
	ResourceType_KV_BANK_DENOM                           ResourceType = 19
	ResourceType_KV_BANK_BALANCES                        ResourceType = 20
	ResourceType_KV_TOKENFACTORY_DENOM                   ResourceType = 21
	ResourceType_KV_TOKENFACTORY_METADATA                ResourceType = 22
	ResourceType_KV_TOKENFACTORY_ADMIN                   ResourceType = 23
	ResourceType_KV_TOKENFACTORY_CREATOR                 ResourceType = 24
	ResourceType_KV_ORACLE_EXCHANGE_RATE                 ResourceType = 25
	ResourceType_KV_ORACLE_VOTE_PENALTY_COUNTER          ResourceType = 26
	ResourceType_KV_ORACLE_PRICE_SNAPSHOT                ResourceType = 27
	ResourceType_KV_STAKING_VALIDATION_POWER             ResourceType = 28
	ResourceType_KV_STAKING_TOTAL_POWER                  ResourceType = 29
	ResourceType_KV_STAKING_VALIDATORS_CON_ADDR          ResourceType = 30
	ResourceType_KV_STAKING_UNBONDING_DELEGATION         ResourceType = 31
	ResourceType_KV_STAKING_UNBONDING_DELEGATION_VAL     ResourceType = 32
	ResourceType_KV_STAKING_REDELEGATION                 ResourceType = 33
	ResourceType_KV_STAKING_REDELEGATION_VAL_SRC         ResourceType = 34
	ResourceType_KV_STAKING_REDELEGATION_VAL_DST         ResourceType = 35
	ResourceType_KV_STAKING_REDELEGATION_QUEUE           ResourceType = 36
	ResourceType_KV_STAKING_VALIDATOR_QUEUE              ResourceType = 37
	ResourceType_KV_STAKING_HISTORICAL_INFO              ResourceType = 38
	ResourceType_KV_STAKING_UNBONDING                    ResourceType = 39
	ResourceType_KV_STAKING_VALIDATORS_BY_POWER          ResourceType = 41
	ResourceType_KV_DISTRIBUTION                         ResourceType = 40
	ResourceType_KV_DISTRIBUTION_FEE_POOL                ResourceType = 42
	ResourceType_KV_DISTRIBUTION_PROPOSER_KEY            ResourceType = 43
	ResourceType_KV_DISTRIBUTION_OUTSTANDING_REWARDS     ResourceType = 44
	ResourceType_KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR ResourceType = 45
	ResourceType_KV_DISTRIBUTION_DELEGATOR_STARTING_INFO ResourceType = 46
	ResourceType_KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS  ResourceType = 47
	ResourceType_KV_DISTRIBUTION_VAL_CURRENT_REWARDS     ResourceType = 48
	ResourceType_KV_DISTRIBUTION_VAL_ACCUM_COMMISSION    ResourceType = 49
	ResourceType_KV_DISTRIBUTION_SLASH_EVENT             ResourceType = 50
)

var ResourceType_name = map[int32]string{
	0:  "ANY",
	1:  "KV",
	2:  "Mem",
	3:  "DexMem",
	4:  "KV_BANK",
	5:  "KV_STAKING",
	6:  "KV_WASM",
	7:  "KV_ORACLE",
	8:  "KV_DEX",
	9:  "KV_EPOCH",
	10: "KV_TOKENFACTORY",
	11: "KV_ORACLE_VOTE_TARGETS",
	12: "KV_ORACLE_AGGREGATE_VOTES",
	13: "KV_ORACLE_FEEDERS",
	14: "KV_STAKING_DELEGATION",
	15: "KV_STAKING_VALIDATOR",
	16: "KV_AUTH",
	17: "KV_AUTH_ADDRESS_STORE",
	18: "KV_BANK_SUPPLY",
	19: "KV_BANK_DENOM",
	20: "KV_BANK_BALANCES",
	21: "KV_TOKENFACTORY_DENOM",
	22: "KV_TOKENFACTORY_METADATA",
	23: "KV_TOKENFACTORY_ADMIN",
	24: "KV_TOKENFACTORY_CREATOR",
	25: "KV_ORACLE_EXCHANGE_RATE",
	26: "KV_ORACLE_VOTE_PENALTY_COUNTER",
	27: "KV_ORACLE_PRICE_SNAPSHOT",
	28: "KV_STAKING_VALIDATION_POWER",
	29: "KV_STAKING_TOTAL_POWER",
	30: "KV_STAKING_VALIDATORS_CON_ADDR",
	31: "KV_STAKING_UNBONDING_DELEGATION",
	32: "KV_STAKING_UNBONDING_DELEGATION_VAL",
	33: "KV_STAKING_REDELEGATION",
	34: "KV_STAKING_REDELEGATION_VAL_SRC",
	35: "KV_STAKING_REDELEGATION_VAL_DST",
	36: "KV_STAKING_REDELEGATION_QUEUE",
	37: "KV_STAKING_VALIDATOR_QUEUE",
	38: "KV_STAKING_HISTORICAL_INFO",
	39: "KV_STAKING_UNBONDING",
	41: "KV_STAKING_VALIDATORS_BY_POWER",
	40: "KV_DISTRIBUTION",
	42: "KV_DISTRIBUTION_FEE_POOL",
	43: "KV_DISTRIBUTION_PROPOSER_KEY",
	44: "KV_DISTRIBUTION_OUTSTANDING_REWARDS",
	45: "KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR",
	46: "KV_DISTRIBUTION_DELEGATOR_STARTING_INFO",
	47: "KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS",
	48: "KV_DISTRIBUTION_VAL_CURRENT_REWARDS",
	49: "KV_DISTRIBUTION_VAL_ACCUM_COMMISSION",
	50: "KV_DISTRIBUTION_SLASH_EVENT",
}

var ResourceType_value = map[string]int32{
	"ANY":                                     0,
	"KV":                                      1,
	"Mem":                                     2,
	"DexMem":                                  3,
	"KV_BANK":                                 4,
	"KV_STAKING":                              5,
	"KV_WASM":                                 6,
	"KV_ORACLE":                               7,
	"KV_DEX":                                  8,
	"KV_EPOCH":                                9,
	"KV_TOKENFACTORY":                         10,
	"KV_ORACLE_VOTE_TARGETS":                  11,
	"KV_ORACLE_AGGREGATE_VOTES":               12,
	"KV_ORACLE_FEEDERS":                       13,
	"KV_STAKING_DELEGATION":                   14,
	"KV_STAKING_VALIDATOR":                    15,
	"KV_AUTH":                                 16,
	"KV_AUTH_ADDRESS_STORE":                   17,
	"KV_BANK_SUPPLY":                          18,
	"KV_BANK_DENOM":                           19,
	"KV_BANK_BALANCES":                        20,
	"KV_TOKENFACTORY_DENOM":                   21,
	"KV_TOKENFACTORY_METADATA":                22,
	"KV_TOKENFACTORY_ADMIN":                   23,
	"KV_TOKENFACTORY_CREATOR":                 24,
	"KV_ORACLE_EXCHANGE_RATE":                 25,
	"KV_ORACLE_VOTE_PENALTY_COUNTER":          26,
	"KV_ORACLE_PRICE_SNAPSHOT":                27,
	"KV_STAKING_VALIDATION_POWER":             28,
	"KV_STAKING_TOTAL_POWER":                  29,
	"KV_STAKING_VALIDATORS_CON_ADDR":          30,
	"KV_STAKING_UNBONDING_DELEGATION":         31,
	"KV_STAKING_UNBONDING_DELEGATION_VAL":     32,
	"KV_STAKING_REDELEGATION":                 33,
	"KV_STAKING_REDELEGATION_VAL_SRC":         34,
	"KV_STAKING_REDELEGATION_VAL_DST":         35,
	"KV_STAKING_REDELEGATION_QUEUE":           36,
	"KV_STAKING_VALIDATOR_QUEUE":              37,
	"KV_STAKING_HISTORICAL_INFO":              38,
	"KV_STAKING_UNBONDING":                    39,
	"KV_STAKING_VALIDATORS_BY_POWER":          41,
	"KV_DISTRIBUTION":                         40,
	"KV_DISTRIBUTION_FEE_POOL":                42,
	"KV_DISTRIBUTION_PROPOSER_KEY":            43,
	"KV_DISTRIBUTION_OUTSTANDING_REWARDS":     44,
	"KV_DISTRIBUTION_DELEGATOR_WITHDRAW_ADDR": 45,
	"KV_DISTRIBUTION_DELEGATOR_STARTING_INFO": 46,
	"KV_DISTRIBUTION_VAL_HISTORICAL_REWARDS":  47,
	"KV_DISTRIBUTION_VAL_CURRENT_REWARDS":     48,
	"KV_DISTRIBUTION_VAL_ACCUM_COMMISSION":    49,
	"KV_DISTRIBUTION_SLASH_EVENT":             50,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{2}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessOperationSelectorType", AccessOperationSelectorType_name, AccessOperationSelectorType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0xcd, 0x72, 0xdb, 0x36,
	0x10, 0xc7, 0xa5, 0xd8, 0xf1, 0xc7, 0xfa, 0x23, 0x6b, 0xc4, 0x4e, 0xfc, 0x29, 0x27, 0xb1, 0x1b,
	0xa7, 0x4a, 0x23, 0xc5, 0xed, 0xad, 0x37, 0x88, 0x84, 0x25, 0x86, 0x14, 0xc0, 0x00, 0x20, 0x15,
	0xf5, 0x82, 0xb1, 0x55, 0x4e, 0x9b, 0x69, 0x6c, 0x7a, 0x24, 0xa5, 0xd3, 0x3c, 0x41, 0xaf, 0x7d,
	0xac, 0x1e, 0x73, 0xec, 0xb1, 0x63, 0xbf, 0x48, 0x87, 0x14, 0x64, 0xcb, 0x8c, 0x53, 0x9f, 0x24,
	0xe1, 0xff, 0xdb, 0xc5, 0xee, 0x7f, 0x01, 0x01, 0xf6, 0x7b, 0xe9, 0xe0, 0x34, 0x1d, 0xd4, 0x8f,
	0x7b, 0xbd, 0x64, 0x30, 0xe8, 0xa5, 0x67, 0xc3, 0x7e, 0xfa, 0xa1, 0xde, 0x4b, 0xcf, 0x06, 0xc3,
	0xe3, 0xb3, 0xe1, 0xa0, 0x76, 0xde, 0x4f, 0x87, 0x29, 0xd9, 0x1e, 0x51, 0xb5, 0x1b, 0x54, 0xed,
	0xf7, 0xc3, 0x93, 0x64, 0x78, 0x7c, 0x58, 0xfd, 0x11, 0x80, 0xe6, 0x82, 0xfe, 0x74, 0x9e, 0x90,
	0x05, 0x98, 0x8d, 0xb8, 0xcf, 0x45, 0x87, 0x63, 0x89, 0xcc, 0xc1, 0xb4, 0x64, 0xd4, 0xc5, 0x32,
	0x99, 0x87, 0xfb, 0x1d, 0xe9, 0x69, 0x86, 0xf7, 0x08, 0xc0, 0x8c, 0x23, 0xda, 0x6d, 0x4f, 0xe3,
	0x54, 0xb5, 0x0e, 0x5b, 0xa3, 0x58, 0x71, 0x9e, 0xf4, 0x8f, 0x87, 0xef, 0xd3, 0x33, 0x95, 0x7c,
	0x48, 0x7a, 0xc3, 0xb4, 0x9f, 0x27, 0x9b, 0x83, 0x69, 0x2e, 0x38, 0xc3, 0x12, 0x99, 0x81, 0x7b,
	0x6f, 0xde, 0x62, 0xb9, 0xfa, 0xe7, 0x02, 0x2c, 0xca, 0x64, 0x90, 0x7e, 0xec, 0xf7, 0x92, 0x1c,
	0x99, 0x85, 0x29, 0xca, 0xbb, 0x23, 0xc2, 0x8f, 0xb1, 0x9c, 0x2d, 0xb4, 0x93, 0xd3, 0xd1, 0x3e,
	0x6e, 0xf2, 0x47, 0xf6, 0x7d, 0x2a, 0xab, 0xca, 0x8f, 0x4d, 0x83, 0x72, 0x1f, 0xa7, 0xc9, 0x32,
	0x80, 0x1f, 0x1b, 0xa5, 0xa9, 0xef, 0xf1, 0x26, 0xde, 0xb7, 0x62, 0x87, 0xaa, 0x36, 0xce, 0x90,
	0x25, 0x98, 0xf7, 0x63, 0x23, 0x24, 0x75, 0x02, 0x86, 0xb3, 0x59, 0x12, 0x3f, 0x36, 0x2e, 0x7b,
	0x87, 0x73, 0x64, 0x11, 0xe6, 0xfc, 0xd8, 0xb0, 0x50, 0x38, 0x2d, 0x9c, 0x27, 0x0f, 0xe1, 0x81,
	0x1f, 0x1b, 0x2d, 0x7c, 0xc6, 0x8f, 0xa8, 0xa3, 0x85, 0xec, 0x22, 0x90, 0x4d, 0x78, 0x74, 0x15,
	0x6d, 0x62, 0xa1, 0x99, 0xd1, 0x54, 0x36, 0x99, 0x56, 0xb8, 0x40, 0x76, 0x60, 0xe3, 0x5a, 0xa3,
	0xcd, 0xa6, 0x64, 0x4d, 0xaa, 0x47, 0x94, 0xc2, 0x45, 0xb2, 0x06, 0x2b, 0xd7, 0xf2, 0x11, 0x63,
	0x2e, 0x93, 0x0a, 0x97, 0xc8, 0x06, 0xac, 0x5d, 0x17, 0x6b, 0x5c, 0x16, 0x64, 0x51, 0x9e, 0xe0,
	0xb8, 0x4c, 0xd6, 0x61, 0x75, 0x42, 0x8a, 0x69, 0xe0, 0xb9, 0x54, 0x0b, 0x89, 0x0f, 0x6c, 0x47,
	0x34, 0xd2, 0x2d, 0x44, 0x9b, 0x21, 0xfb, 0x61, 0xa8, 0xeb, 0x4a, 0xa6, 0x94, 0x51, 0x5a, 0x48,
	0x86, 0x2b, 0x84, 0xc0, 0xb2, 0xb5, 0xc5, 0xa8, 0x28, 0x0c, 0x83, 0x2e, 0x12, 0xb2, 0x02, 0x4b,
	0xe3, 0x35, 0x97, 0x71, 0xd1, 0xc6, 0x87, 0x64, 0x15, 0x70, 0xbc, 0xd4, 0xa0, 0x01, 0xe5, 0x0e,
	0x53, 0xb8, 0x6a, 0xf3, 0x4e, 0x1a, 0x60, 0x03, 0xd6, 0xc8, 0x36, 0xac, 0x17, 0xa5, 0x36, 0xd3,
	0xd4, 0xa5, 0x9a, 0xe2, 0xa3, 0xdb, 0x02, 0xa9, 0xdb, 0xf6, 0x38, 0x3e, 0x26, 0x5b, 0xf0, 0xb8,
	0x28, 0x39, 0x92, 0xe5, 0x5d, 0xad, 0x5b, 0xd1, 0x3a, 0xc4, 0xde, 0x39, 0x2d, 0xca, 0x9b, 0xcc,
	0x48, 0xaa, 0x19, 0x6e, 0x90, 0x67, 0x50, 0x29, 0x38, 0x1f, 0x32, 0x4e, 0x03, 0xdd, 0x35, 0x8e,
	0x88, 0xb8, 0x66, 0x12, 0x37, 0x6d, 0x59, 0x96, 0x09, 0xa5, 0xe7, 0x30, 0xa3, 0x38, 0x0d, 0x55,
	0x4b, 0x68, 0xdc, 0x22, 0xbb, 0xb0, 0xf5, 0xa5, 0x9d, 0x9e, 0xe0, 0x26, 0x14, 0x1d, 0x26, 0x71,
	0xdb, 0x0e, 0x77, 0x0c, 0x68, 0xa1, 0x69, 0x60, 0xb5, 0x1d, 0xbb, 0xfd, 0x17, 0xb3, 0x50, 0xc6,
	0x11, 0x3c, 0xb7, 0x1d, 0x2b, 0x64, 0x0f, 0x76, 0x27, 0x98, 0x88, 0x37, 0x04, 0x77, 0x0b, 0x43,
	0xdd, 0x25, 0x07, 0xb0, 0x77, 0x07, 0x94, 0x65, 0xc7, 0x27, 0xd6, 0x8d, 0x31, 0x28, 0xd9, 0x44,
	0x96, 0xa7, 0x85, 0xad, 0x26, 0xc5, 0x2c, 0xda, 0x28, 0xe9, 0xe0, 0xb3, 0xbb, 0x20, 0x57, 0x69,
	0xdc, 0x23, 0x4f, 0x61, 0xe7, 0x6b, 0xd0, 0xdb, 0x88, 0x45, 0x0c, 0xf7, 0x49, 0x05, 0x36, 0x6f,
	0xeb, 0xdd, 0xea, 0xdf, 0x14, 0xf4, 0x96, 0x97, 0x9d, 0x3e, 0xcf, 0xa1, 0x81, 0xf1, 0xf8, 0x91,
	0xc0, 0xe7, 0x85, 0x73, 0x7c, 0xd5, 0x32, 0x1e, 0x7c, 0xdd, 0xd5, 0x46, 0xd7, 0x3a, 0xff, 0xad,
	0xbd, 0x87, 0xae, 0xa7, 0xb4, 0xf4, 0x1a, 0x51, 0xde, 0xff, 0x0b, 0x3b, 0xe9, 0xc9, 0xc5, 0xec,
	0x4a, 0x99, 0x50, 0x88, 0x00, 0xab, 0xe4, 0x09, 0x6c, 0x17, 0xd5, 0x50, 0x8a, 0x50, 0x28, 0x26,
	0x8d, 0xcf, 0xba, 0xf8, 0xd2, 0x4e, 0xe1, 0x06, 0x21, 0x22, 0xad, 0x34, 0x1d, 0x0d, 0x43, 0xb2,
	0x0e, 0x95, 0xae, 0xc2, 0xef, 0xc8, 0x4b, 0x38, 0x28, 0x82, 0xd6, 0x21, 0x21, 0x4d, 0xc7, 0xd3,
	0x2d, 0x57, 0xd2, 0xce, 0xe8, 0x00, 0xbc, 0xfa, 0x7f, 0x58, 0x69, 0x2a, 0x75, 0x96, 0x3c, 0x77,
	0xa5, 0x46, 0xaa, 0xf0, 0xbc, 0x08, 0x67, 0x53, 0x99, 0xb0, 0x6f, 0x5c, 0x45, 0xfd, 0xb6, 0x72,
	0x33, 0xd6, 0x89, 0xa4, 0x64, 0x5c, 0x5f, 0x81, 0xaf, 0xc9, 0x0b, 0xd8, 0xbf, 0x0d, 0xa4, 0x8e,
	0x13, 0xb5, 0x4d, 0xfe, 0xaf, 0xac, 0x54, 0xe6, 0xe0, 0xa1, 0xbd, 0x0d, 0x37, 0x48, 0x15, 0x50,
	0xd5, 0x32, 0x2c, 0x66, 0x5c, 0xe3, 0xf7, 0x8d, 0x37, 0x7f, 0x5f, 0x54, 0xca, 0x9f, 0x2f, 0x2a,
	0xe5, 0x7f, 0x2f, 0x2a, 0xe5, 0xbf, 0x2e, 0x2b, 0xa5, 0xcf, 0x97, 0x95, 0xd2, 0x3f, 0x97, 0x95,
	0xd2, 0x4f, 0xaf, 0x7f, 0x79, 0x3f, 0xfc, 0xf5, 0xe3, 0x49, 0xad, 0x97, 0x9e, 0xd6, 0xed, 0xfb,
	0x32, 0xfa, 0x78, 0x35, 0xf8, 0xf9, 0xb7, 0xfa, 0xf0, 0xd3, 0x79, 0x52, 0x78, 0x70, 0x4e, 0x66,
	0xf2, 0x77, 0xe6, 0x87, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xc1, 0x8f, 0x41, 0x8f, 0x06,
	0x00, 0x00,
}
