package formatter

import (
	"github.com/KoNekoD/go-deptrac/pkg/config_contract"
)

type CodeclimateConfig struct {
	Failure   config_contract.CodeclimateLevelEnum
	Skipped   config_contract.CodeclimateLevelEnum
	Uncovered config_contract.CodeclimateLevelEnum
}

func newCodeclimateConfig(failure config_contract.CodeclimateLevelEnum, skipped config_contract.CodeclimateLevelEnum, uncovered config_contract.CodeclimateLevelEnum) *CodeclimateConfig {
	return &CodeclimateConfig{
		Failure:   failure,
		Skipped:   skipped,
		Uncovered: uncovered,
	}
}

func CreateCodeclimateConfig(failure *config_contract.CodeclimateLevelEnum, skipped *config_contract.CodeclimateLevelEnum, uncovered *config_contract.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := config_contract.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := config_contract.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := config_contract.CodeclimateLevelEnumInfo
		uncovered = &uncoveredTmp
	}
	return newCodeclimateConfig(*failure, *skipped, *uncovered)
}

func (c *CodeclimateConfig) severity(failure *config_contract.CodeclimateLevelEnum, skipped *config_contract.CodeclimateLevelEnum, uncovered *config_contract.CodeclimateLevelEnum) *CodeclimateConfig {
	if failure == nil {
		failureTmp := config_contract.CodeclimateLevelEnumBlocker
		failure = &failureTmp
	}
	if skipped == nil {
		skippedTmp := config_contract.CodeclimateLevelEnumMinor
		skipped = &skippedTmp
	}
	if uncovered == nil {
		uncoveredTmp := config_contract.CodeclimateLevelEnumInfo
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
