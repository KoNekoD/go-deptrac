package parser

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
)

type ParserInterface interface {
	ParseFile(file string) (*ast_map.FileReference, error)
}
