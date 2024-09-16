package output_formatter_supportive

import (
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/Exception/InvalidServiceInLocatorException"
	output_formatter_contract2 "github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	"golang.org/x/exp/maps"
	"reflect"
)

type FormatterProvider struct {
	formatterLocator map[output_formatter_contract2.OutputFormatterType]output_formatter_contract2.OutputFormatterInterface
}

func NewFormatterProvider(formatterLocator map[output_formatter_contract2.OutputFormatterType]output_formatter_contract2.OutputFormatterInterface) *FormatterProvider {
	return &FormatterProvider{
		formatterLocator: formatterLocator,
	}
}

func (f *FormatterProvider) Get(id output_formatter_contract2.OutputFormatterType) (output_formatter_contract2.OutputFormatterInterface, error) {
	service, ok := f.formatterLocator[id]

	if !ok {
		return nil, InvalidServiceInLocatorException.NewInvalidServiceInLocatorExceptionInvalidType(id, reflect.TypeOf(service).Name(), "OutputFormatterInterface.OutputFormatterInterface{}")
	}

	return service, nil
}

func (f *FormatterProvider) Has(id output_formatter_contract2.OutputFormatterType) bool {
	_, ok := f.formatterLocator[id]
	return ok
}

func (f *FormatterProvider) GetKnownFormatters() []output_formatter_contract2.OutputFormatterType {
	return maps.Keys(f.formatterLocator)
}
