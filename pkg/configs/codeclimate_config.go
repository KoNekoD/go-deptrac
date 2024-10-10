package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/enums"
	"github.com/KoNekoD/go-deptrac/pkg/formatters"
)

type CodeclimateConfig struct {
	Failure   enums.CodeclimateLevelEnum
	Skipped   enums.CodeclimateLevelEnum
	Uncovered enums.CodeclimateLevelEnum
}

func newCodeclimateConfig(failure enums.CodeclimateLevelEnum, skipped enums.CodeclimateLevelEnum, uncovered enums.CodeclimateLevelEnum) *CodeclimateConfig {
	return &CodeclimateConfig{
		Failure:   failure,
		Skipped:   skipped,
		Uncovered: uncovered,
	}
}

func CreateCodeclimateConfig(failure *enums.CodeclimateLevelEnum, skipped *enums.CodeclimateLevelEnum, uncovered *enums.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := enums.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := enums.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := enums.CodeclimateLevelEnumInfo
		uncovered = &uncoveredTmp
	}
	return newCodeclimateConfig(*failure, *skipped, *uncovered)
}

func (c *CodeclimateConfig) severity(failure *enums.CodeclimateLevelEnum, skipped *enums.CodeclimateLevelEnum, uncovered *enums.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := enums.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := enums.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := enums.CodeclimateLevelEnumInfo
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

func (c *CodeclimateConfig) GetName() formatters.FormatterType {
	return formatters.FormatterTypeCodeclimateConfig
}
