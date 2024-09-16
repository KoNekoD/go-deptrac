package parser

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type ParserInterface interface {
	ParseFile(file string) (*ast_map.FileReference, error)
}
