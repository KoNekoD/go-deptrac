package CodeclimateConfig

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/CodeclimateLevelEnum"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface/FormatterType"
)

type CodeclimateConfig struct {
	Failure   CodeclimateLevelEnum.CodeclimateLevelEnum
	Skipped   CodeclimateLevelEnum.CodeclimateLevelEnum
	Uncovered CodeclimateLevelEnum.CodeclimateLevelEnum
}

func newCodeclimateConfig(failure CodeclimateLevelEnum.CodeclimateLevelEnum, skipped CodeclimateLevelEnum.CodeclimateLevelEnum, uncovered CodeclimateLevelEnum.CodeclimateLevelEnum) *CodeclimateConfig {
	return &CodeclimateConfig{
		Failure:   failure,
		Skipped:   skipped,
		Uncovered: uncovered,
	}
}

func CreateCodeclimateConfig(failure *CodeclimateLevelEnum.CodeclimateLevelEnum, skipped *CodeclimateLevelEnum.CodeclimateLevelEnum, uncovered *CodeclimateLevelEnum.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := CodeclimateLevelEnum.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := CodeclimateLevelEnum.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := CodeclimateLevelEnum.CodeclimateLevelEnumInfo
		uncovered = &uncoveredTmp
	}
	return newCodeclimateConfig(*failure, *skipped, *uncovered)
}

func (c *CodeclimateConfig) severity(failure *CodeclimateLevelEnum.CodeclimateLevelEnum, skipped *CodeclimateLevelEnum.CodeclimateLevelEnum, uncovered *CodeclimateLevelEnum.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := CodeclimateLevelEnum.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := CodeclimateLevelEnum.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := CodeclimateLevelEnum.CodeclimateLevelEnumInfo
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

func (c *CodeclimateConfig) GetName() FormatterType.FormatterType {
	return FormatterType.FormatterTypeCodeclimateConfig
}
