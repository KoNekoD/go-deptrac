package cache

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	"path/filepath"
)

type AstFileReferenceInMemoryCache struct {
	cache map[string]*ast_map.FileReference
}

func NewAstFileReferenceInMemoryCache() *AstFileReferenceInMemoryCache {
	return &AstFileReferenceInMemoryCache{
		cache: make(map[string]*ast_map.FileReference),
	}
}

func (c *AstFileReferenceInMemoryCache) Get(pathInput string) (*ast_map.FileReference, error) {
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

func (c *AstFileReferenceInMemoryCache) Set(fileReference *ast_map.FileReference) error {
	path, err := filepath.Abs(*fileReference.Filepath)
	if err != nil {
		return err
	}

	c.cache[path] = fileReference

	return nil
}
