package FormatterConfiguration

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface/FormatterType"
)

type FormatterConfiguration struct {
	config map[FormatterType.FormatterType]FormatterConfigInterface.FormatterConfigInterface
}

func NewFormatterConfiguration(config map[FormatterType.FormatterType]FormatterConfigInterface.FormatterConfigInterface) *FormatterConfiguration {
	return &FormatterConfiguration{config: config}
}

func (f *FormatterConfiguration) GetConfigFor(area FormatterType.FormatterType) FormatterConfigInterface.FormatterConfigInterface {
	return f.config[area]
}
