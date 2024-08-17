package FormatterConfigInterface

import "github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/Formatter/FormatterConfigInterface/FormatterType"

type FormatterConfigInterface interface {
	GetName() FormatterType.FormatterType
	ToArray() map[string]interface{}
}
