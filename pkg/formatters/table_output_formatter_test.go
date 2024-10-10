package formatters

import (
	"fmt"
	"github.com/gookit/color"
	"testing"
)

func TestColorOk(t *testing.T) {
	fmt.Println(color.Sprint(color.Render("<info>123</>")))
}
