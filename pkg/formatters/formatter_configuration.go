package formatters

type FormatterConfiguration struct {
	config map[FormatterType]FormatterConfigInterface
}

func NewFormatterConfiguration(config map[FormatterType]FormatterConfigInterface) *FormatterConfiguration {
	return &FormatterConfiguration{config: config}
}

func (f *FormatterConfiguration) GetConfigFor(area FormatterType) FormatterConfigInterface {
	return f.config[area]
}
