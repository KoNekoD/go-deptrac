package services

import (
	"github.com/gookit/color"
)

type SymfonyOutput struct {
	style OutputStyleInterface
}

func NewSymfonyOutput(style OutputStyleInterface) *SymfonyOutput {
	return &SymfonyOutput{
		style: style,
	}
}

func (o *SymfonyOutput) WriteFormatted(message string) {
	color.Print(message)
}

func (o *SymfonyOutput) WriteLineFormatted(message StringOrArrayOfStrings) {
	color.Println(message.ToString())
}

func (o *SymfonyOutput) WriteRaw(message string) {
	color.Println(message)
}

func (o *SymfonyOutput) GetStyle() OutputStyleInterface {
	return o.style
}

func (o *SymfonyOutput) IsVerbose() bool {
	return o.style.IsVerbose()
}

func (o *SymfonyOutput) IsDebug() bool {
	return o.style.IsDebug()
}
