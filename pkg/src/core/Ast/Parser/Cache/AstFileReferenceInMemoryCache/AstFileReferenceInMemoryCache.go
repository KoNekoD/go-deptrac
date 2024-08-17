package AstFileReferenceInMemoryCache

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"path/filepath"
)

type AstFileReferenceInMemoryCache struct {
	cache map[string]*AstMap.FileReference
}

func NewAstFileReferenceInMemoryCache() *AstFileReferenceInMemoryCache {
	return &AstFileReferenceInMemoryCache{
		cache: make(map[string]*AstMap.FileReference),
	}
}

func (c *AstFileReferenceInMemoryCache) Get(pathInput string) (*AstMap.FileReference, error) {
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

func (c *AstFileReferenceInMemoryCache) Set(fileReference *AstMap.FileReference) error {
	path, err := filepath.Abs(*fileReference.Filepath)
	if err != nil {
		return err
	}

	c.cache[path] = fileReference

	return nil
}
