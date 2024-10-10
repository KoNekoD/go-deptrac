package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/formatters_configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type FormatterConfiguration struct {
	config map[enums.FormatterType]formatters_configs.FormatterConfigInterface
}

func NewFormatterConfiguration(config map[enums.FormatterType]formatters_configs.FormatterConfigInterface) *FormatterConfiguration {
	return &FormatterConfiguration{config: config}
}

func (f *FormatterConfiguration) GetConfigFor(area enums.FormatterType) formatters_configs.FormatterConfigInterface {
	return f.config[area]
}
