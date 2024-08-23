package formatter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config"
)

type CodeclimateConfig struct {
	Failure   config.CodeclimateLevelEnum
	Skipped   config.CodeclimateLevelEnum
	Uncovered config.CodeclimateLevelEnum
}

func newCodeclimateConfig(failure config.CodeclimateLevelEnum, skipped config.CodeclimateLevelEnum, uncovered config.CodeclimateLevelEnum) *CodeclimateConfig {
	return &CodeclimateConfig{
		Failure:   failure,
		Skipped:   skipped,
		Uncovered: uncovered,
	}
}

func CreateCodeclimateConfig(failure *config.CodeclimateLevelEnum, skipped *config.CodeclimateLevelEnum, uncovered *config.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := config.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := config.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := config.CodeclimateLevelEnumInfo
		uncovered = &uncoveredTmp
	}
	return newCodeclimateConfig(*failure, *skipped, *uncovered)
}

func (c *CodeclimateConfig) severity(failure *config.CodeclimateLevelEnum, skipped *config.CodeclimateLevelEnum, uncovered *config.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := config.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := config.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := config.CodeclimateLevelEnumInfo
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

func (c *CodeclimateConfig) GetName() FormatterType {
	return FormatterTypeCodeclimateConfig
}
