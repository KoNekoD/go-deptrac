package AstFileReferenceCacheInterface

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
)

type AstFileReferenceCacheInterface interface {
	Get(filepath string) (*AstMap.FileReference, error)
	Set(fileReference *AstMap.FileReference) error
}
