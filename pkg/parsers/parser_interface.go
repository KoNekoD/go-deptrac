package parsers

import "github.com/KoNekoD/go-deptrac/pkg/references"

type ParserInterface interface {
	ParseFile(file string) (*references.FileReference, error)
}
