package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"golang.org/x/exp/maps"
	"reflect"
)

type FormatterProvider struct {
	formatterLocator map[enums.OutputFormatterType]OutputFormatterInterface
}

func NewFormatterProvider(formatterLocator map[enums.OutputFormatterType]OutputFormatterInterface) *FormatterProvider {
	return &FormatterProvider{
		formatterLocator: formatterLocator,
	}
}

func (f *FormatterProvider) Get(id enums.OutputFormatterType) (OutputFormatterInterface, error) {
	service, ok := f.formatterLocator[id]

	if !ok {
		return nil, apperrors.NewInvalidServiceInLocatorExceptionInvalidType(string(id), reflect.TypeOf(service).Name(), "OutputFormatterInterface.OutputFormatterInterface{}")
	}

	return service, nil
}

func (f *FormatterProvider) Has(id enums.OutputFormatterType) bool {
	_, ok := f.formatterLocator[id]
	return ok
}

func (f *FormatterProvider) GetKnownFormatters() []enums.OutputFormatterType {
	return maps.Keys(f.formatterLocator)
}
