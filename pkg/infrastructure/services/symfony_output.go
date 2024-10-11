package services

import (
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"github.com/gookit/color"
)

type SymfonyOutput struct {
	style services2.OutputStyleInterface
}

func NewSymfonyOutput(style services2.OutputStyleInterface) *SymfonyOutput {
	return &SymfonyOutput{
		style: style,
	}
}

func (o *SymfonyOutput) WriteFormatted(message string) {
	color.Print(message)
}

func (o *SymfonyOutput) WriteLineFormatted(message services2.StringOrArrayOfStrings) {
	color.Println(message.ToString())
}

func (o *SymfonyOutput) WriteRaw(message string) {
	color.Println(message)
}

func (o *SymfonyOutput) GetStyle() services2.OutputStyleInterface {
	return o.style
}

func (o *SymfonyOutput) IsVerbose() bool {
	return o.style.IsVerbose()
}

func (o *SymfonyOutput) IsDebug() bool {
	return o.style.IsDebug()
}
