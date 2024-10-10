package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"golang.org/x/exp/maps"
	"reflect"
)

type FormatterProvider struct {
	formatterLocator map[OutputFormatterType]OutputFormatterInterface
}

func NewFormatterProvider(formatterLocator map[OutputFormatterType]OutputFormatterInterface) *FormatterProvider {
	return &FormatterProvider{
		formatterLocator: formatterLocator,
	}
}

func (f *FormatterProvider) Get(id OutputFormatterType) (OutputFormatterInterface, error) {
	service, ok := f.formatterLocator[id]

	if !ok {
		return nil, apperrors.NewInvalidServiceInLocatorExceptionInvalidType(string(id), reflect.TypeOf(service).Name(), "OutputFormatterInterface.OutputFormatterInterface{}")
	}

	return service, nil
}

func (f *FormatterProvider) Has(id OutputFormatterType) bool {
	_, ok := f.formatterLocator[id]
	return ok
}

func (f *FormatterProvider) GetKnownFormatters() []OutputFormatterType {
	return maps.Keys(f.formatterLocator)
}
