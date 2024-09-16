package configuration

import (
	formatter2 "github.com/KoNekoD/go-deptrac/pkg/config_contract/formatter"
)

type FormatterConfiguration struct {
	config map[formatter2.FormatterType]formatter2.FormatterConfigInterface
}

func NewFormatterConfiguration(config map[formatter2.FormatterType]formatter2.FormatterConfigInterface) *FormatterConfiguration {
	return &FormatterConfiguration{config: config}
}

func (f *FormatterConfiguration) GetConfigFor(area formatter2.FormatterType) formatter2.FormatterConfigInterface {
	return f.config[area]
}
