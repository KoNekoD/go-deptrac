package ast_file_reference_cache

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"path/filepath"
)

type AstFileReferenceInMemoryCache struct {
	cache map[string]*tokens_references.FileReference
}

func NewAstFileReferenceInMemoryCache() *AstFileReferenceInMemoryCache {
	return &AstFileReferenceInMemoryCache{
		cache: make(map[string]*tokens_references.FileReference),
	}
}

func (c *AstFileReferenceInMemoryCache) Get(pathInput string) (*tokens_references.FileReference, error) {
	path, err := filepath.Abs(pathInput)
	if err != nil {
		return nil, err
	}

	v, ok := c.cache[path]
	if !ok {
		return nil, nil
	}

	return v, nil
}

func (c *AstFileReferenceInMemoryCache) Set(fileReference *tokens_references.FileReference) error {
	path, err := filepath.Abs(*fileReference.Filepath)
	if err != nil {
		return err
	}

	c.cache[path] = fileReference

	return nil
}
