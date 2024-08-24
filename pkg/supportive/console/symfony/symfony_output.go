package symfony

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/gookit/color"
)

type SymfonyOutput struct {
	style output_formatter.OutputStyleInterface
}

func NewSymfonyOutput(style output_formatter.OutputStyleInterface) *SymfonyOutput {
	return &SymfonyOutput{
		style: style,
	}
}

func (o *SymfonyOutput) WriteFormatted(message string) {
	color.Print(message)
}

func (o *SymfonyOutput) WriteLineFormatted(message output_formatter.StringOrArrayOfStrings) {
	color.Println(message.ToString())
}

func (o *SymfonyOutput) WriteRaw(message string) {
	color.Println(message)
}

func (o *SymfonyOutput) GetStyle() output_formatter.OutputStyleInterface {
	return o.style
}

func (o *SymfonyOutput) IsVerbose() bool {
	return o.style.IsVerbose()
}

func (o *SymfonyOutput) IsDebug() bool {
	return o.style.IsDebug()
}
