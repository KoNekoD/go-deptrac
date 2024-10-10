package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/references"
)

type AstFileReferenceCacheInterface interface {
	Get(filepath string) (*references.FileReference, error)
	Set(fileReference *references.FileReference) error
}
