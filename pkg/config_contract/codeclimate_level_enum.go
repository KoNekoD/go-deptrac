package config_contract

type CodeclimateLevelEnum string

const (
	CodeclimateLevelEnumInfo     CodeclimateLevelEnum = "info"
	CodeclimateLevelEnumMinor    CodeclimateLevelEnum = "minor"
	CodeclimateLevelEnumMajor    CodeclimateLevelEnum = "major"
	CodeclimateLevelEnumCritical CodeclimateLevelEnum = "critical"
	CodeclimateLevelEnumBlocker  CodeclimateLevelEnum = "blocker"
)
