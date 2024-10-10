package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type FormatterConfiguration struct {
	config map[enums.FormatterType]FormatterConfigInterface
}

func NewFormatterConfiguration(config map[enums.FormatterType]FormatterConfigInterface) *FormatterConfiguration {
	return &FormatterConfiguration{config: config}
}

func (f *FormatterConfiguration) GetConfigFor(area enums.FormatterType) FormatterConfigInterface {
	return f.config[area]
}
