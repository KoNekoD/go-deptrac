package configs

import (
	enums2 "github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type CodeclimateConfig struct {
	Failure   enums2.CodeclimateLevelEnum
	Skipped   enums2.CodeclimateLevelEnum
	Uncovered enums2.CodeclimateLevelEnum
}

func newCodeclimateConfig(failure enums2.CodeclimateLevelEnum, skipped enums2.CodeclimateLevelEnum, uncovered enums2.CodeclimateLevelEnum) *CodeclimateConfig {
	return &CodeclimateConfig{
		Failure:   failure,
		Skipped:   skipped,
		Uncovered: uncovered,
	}
}

func CreateCodeclimateConfig(failure *enums2.CodeclimateLevelEnum, skipped *enums2.CodeclimateLevelEnum, uncovered *enums2.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := enums2.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := enums2.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := enums2.CodeclimateLevelEnumInfo
		uncovered = &uncoveredTmp
	}
	return newCodeclimateConfig(*failure, *skipped, *uncovered)
}

func (c *CodeclimateConfig) severity(failure *enums2.CodeclimateLevelEnum, skipped *enums2.CodeclimateLevelEnum, uncovered *enums2.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := enums2.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := enums2.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := enums2.CodeclimateLevelEnumInfo
		uncovered = &uncoveredTmp
	}
	c.Failure = *failure
	c.Skipped = *skipped
	c.Uncovered = *uncovered
	return c
}

func (c *CodeclimateConfig) ToArray() map[string]interface{} {
	return map[string]interface{}{
		"severity": map[string]interface{}{
			"failure":   string(c.Failure),
			"skipped":   string(c.Skipped),
			"uncovered": string(c.Uncovered),
		},
	}
}

func (c *CodeclimateConfig) GetName() enums2.FormatterType {
	return enums2.FormatterTypeCodeclimateConfig
}
