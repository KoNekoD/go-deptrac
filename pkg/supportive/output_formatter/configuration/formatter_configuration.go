package configuration

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config/formatter"
)

type FormatterConfiguration struct {
	config map[formatter.FormatterType]formatter.FormatterConfigInterface
}

func NewFormatterConfiguration(config map[formatter.FormatterType]formatter.FormatterConfigInterface) *FormatterConfiguration {
	return &FormatterConfiguration{config: config}
}

func (f *FormatterConfiguration) GetConfigFor(area formatter.FormatterType) formatter.FormatterConfigInterface {
	return f.config[area]
}
