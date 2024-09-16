package symfony

import (
	"github.com/KoNekoD/go-deptrac/pkg/output_formatter_contract"
	"github.com/gookit/color"
)

type SymfonyOutput struct {
	style output_formatter_contract.OutputStyleInterface
}

func NewSymfonyOutput(style output_formatter_contract.OutputStyleInterface) *SymfonyOutput {
	return &SymfonyOutput{
		style: style,
	}
}

func (o *SymfonyOutput) WriteFormatted(message string) {
	color.Print(message)
}

func (o *SymfonyOutput) WriteLineFormatted(message output_formatter_contract.StringOrArrayOfStrings) {
	color.Println(message.ToString())
}

func (o *SymfonyOutput) WriteRaw(message string) {
	color.Println(message)
}

func (o *SymfonyOutput) GetStyle() output_formatter_contract.OutputStyleInterface {
	return o.style
}

func (o *SymfonyOutput) IsVerbose() bool {
	return o.style.IsVerbose()
}

func (o *SymfonyOutput) IsDebug() bool {
	return o.style.IsDebug()
}
