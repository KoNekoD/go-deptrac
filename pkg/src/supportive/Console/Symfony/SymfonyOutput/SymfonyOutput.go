package SymfonyOutput

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/OutputFormatter/OutputStyleInterface"
)

type SymfonyOutput struct {
	style OutputStyleInterface.OutputStyleInterface
}

func NewSymfonyOutput(style OutputStyleInterface.OutputStyleInterface) *SymfonyOutput {
	return &SymfonyOutput{
		style: style,
	}
}

func (o *SymfonyOutput) WriteFormatted(message string) {
	fmt.Print(message)
}

func (o *SymfonyOutput) WriteLineFormatted(message OutputStyleInterface.StringOrArrayOfStrings) {
	fmt.Println(message.ToString())
}

func (o *SymfonyOutput) WriteRaw(message string) {
	fmt.Println(message)
}

func (o *SymfonyOutput) GetStyle() OutputStyleInterface.OutputStyleInterface {
	return o.style
}

func (o *SymfonyOutput) IsVerbose() bool {
	return o.style.IsVerbose()
}

func (o *SymfonyOutput) IsDebug() bool {
	return o.style.IsDebug()
}
