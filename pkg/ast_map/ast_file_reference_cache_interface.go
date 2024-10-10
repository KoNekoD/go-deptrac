package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
)

type AstFileReferenceCacheInterface interface {
	Get(filepath string) (*tokens_references.FileReference, error)
	Set(fileReference *tokens_references.FileReference) error
}
