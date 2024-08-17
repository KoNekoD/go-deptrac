package FormatterProvider

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputFormatterInterface/OutputFormatterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/Exception/InvalidServiceInLocatorException"
	"golang.org/x/exp/maps"
	"reflect"
)

type FormatterProvider struct {
	formatterLocator map[OutputFormatterType.OutputFormatterType]OutputFormatterInterface.OutputFormatterInterface
}

func NewFormatterProvider(formatterLocator map[OutputFormatterType.OutputFormatterType]OutputFormatterInterface.OutputFormatterInterface) *FormatterProvider {
	return &FormatterProvider{
		formatterLocator: formatterLocator,
	}
}

func (f *FormatterProvider) Get(id OutputFormatterType.OutputFormatterType) (OutputFormatterInterface.OutputFormatterInterface, error) {
	service, ok := f.formatterLocator[id]

	if !ok {
		return nil, InvalidServiceInLocatorException.NewInvalidServiceInLocatorExceptionInvalidType(id, reflect.TypeOf(service).Name(), "OutputFormatterInterface.OutputFormatterInterface{}")
	}

	return service, nil
}

func (f *FormatterProvider) Has(id OutputFormatterType.OutputFormatterType) bool {
	_, ok := f.formatterLocator[id]
	return ok
}

func (f *FormatterProvider) GetKnownFormatters() []OutputFormatterType.OutputFormatterType {
	return maps.Keys(f.formatterLocator)
}
