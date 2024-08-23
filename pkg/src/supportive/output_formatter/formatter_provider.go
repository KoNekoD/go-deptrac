package output_formatter

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/Exception/InvalidServiceInLocatorException"
	"golang.org/x/exp/maps"
	"reflect"
)

type FormatterProvider struct {
	formatterLocator map[output_formatter.OutputFormatterType]output_formatter.OutputFormatterInterface
}

func NewFormatterProvider(formatterLocator map[output_formatter.OutputFormatterType]output_formatter.OutputFormatterInterface) *FormatterProvider {
	return &FormatterProvider{
		formatterLocator: formatterLocator,
	}
}

func (f *FormatterProvider) Get(id output_formatter.OutputFormatterType) (output_formatter.OutputFormatterInterface, error) {
	service, ok := f.formatterLocator[id]

	if !ok {
		return nil, InvalidServiceInLocatorException.NewInvalidServiceInLocatorExceptionInvalidType(id, reflect.TypeOf(service).Name(), "OutputFormatterInterface.OutputFormatterInterface{}")
	}

	return service, nil
}

func (f *FormatterProvider) Has(id output_formatter.OutputFormatterType) bool {
	_, ok := f.formatterLocator[id]
	return ok
}

func (f *FormatterProvider) GetKnownFormatters() []output_formatter.OutputFormatterType {
	return maps.Keys(f.formatterLocator)
}
