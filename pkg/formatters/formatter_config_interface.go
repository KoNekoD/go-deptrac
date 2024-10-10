package formatters

type FormatterConfigInterface interface {
	GetName() FormatterType
	ToArray() map[string]interface{}
}
