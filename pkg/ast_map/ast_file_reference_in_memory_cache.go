package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"path/filepath"
)

type AstFileReferenceInMemoryCache struct {
	cache map[string]*references.FileReference
}

func NewAstFileReferenceInMemoryCache() *AstFileReferenceInMemoryCache {
	return &AstFileReferenceInMemoryCache{
		cache: make(map[string]*references.FileReference),
	}
}

func (c *AstFileReferenceInMemoryCache) Get(pathInput string) (*references.FileReference, error) {
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

func (c *AstFileReferenceInMemoryCache) Set(fileReference *references.FileReference) error {
	path, err := filepath.Abs(*fileReference.Filepath)
	if err != nil {
		return err
	}

	c.cache[path] = fileReference

	return nil
}
