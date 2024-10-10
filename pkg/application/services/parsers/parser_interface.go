package parsers

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

type ParserInterface interface {
	ParseFile(file string) (*tokens_references.FileReference, error)
}
