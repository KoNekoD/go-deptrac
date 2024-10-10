package formatters

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
)

type FormatterConfigInterface interface {
	GetName() enums.FormatterType
	ToArray() map[string]interface{}
}
