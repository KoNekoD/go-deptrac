package cache

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
)

type AstFileReferenceCacheInterface interface {
	Get(filepath string) (*ast_map.FileReference, error)
	Set(fileReference *ast_map.FileReference) error
}
