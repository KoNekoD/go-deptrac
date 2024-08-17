package ParserInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
)

type ParserInterface interface {
	ParseFile(file string) (*AstMap.FileReference, error)
}
