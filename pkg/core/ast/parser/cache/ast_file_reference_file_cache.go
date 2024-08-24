package cache

import (
	"encoding/json"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/file"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
	"path/filepath"
)

type AstFileReferenceFileCache struct {
	cache map[string]struct {
		hash      string
		reference *ast_map.FileReference
	}
	loaded       bool
	parsedFiles  map[string]bool
	cacheFile    string
	cacheVersion string
}

func NewAstFileReferenceFileCache(cacheFile string, cacheVersion string) *AstFileReferenceFileCache {
	return &AstFileReferenceFileCache{
		cache: make(map[string]struct {
			hash      string
			reference *ast_map.FileReference
		}),
		loaded:       false,
		parsedFiles:  make(map[string]bool),
		cacheFile:    cacheFile,
		cacheVersion: cacheVersion,
	}
}

func (c *AstFileReferenceFileCache) Get(filepath string) (*ast_map.FileReference, error) {
	err := c.Load()
	if err != nil {
		return nil, err
	}
	filepath = c.normalizeFilepath(filepath)

	has, err := c.has(filepath)
	if err != nil {
		return nil, err
	}

	if has {
		c.parsedFiles[filepath] = true

		return c.cache[filepath].reference, nil
	}

	return nil, nil
}

func (c *AstFileReferenceFileCache) Set(fileReference *ast_map.FileReference) error {
	err := c.Load()
	if err != nil {
		return err
	}
	normalizedFilepath := c.normalizeFilepath(*fileReference.Filepath)
	c.parsedFiles[normalizedFilepath] = true

	hash, err := util.Sha1File(normalizedFilepath)
	if err != nil {
		return err
	}

	c.cache[normalizedFilepath] = struct {
		hash      string
		reference *ast_map.FileReference
	}{hash: hash, reference: fileReference}

	return nil
}

func (c *AstFileReferenceFileCache) Load() error {
	if c.loaded {
		return nil
	}
	if !util.FileExists(c.cacheFile) || !util.IsReadable(c.cacheFile) {
		return nil
	}
	contents, err := file.FileReaderRead(c.cacheFile)
	if err != nil {
		return err
	}
	cache := struct {
		version string
		payload map[string]struct {
			hash      string
			reference *ast_map.FileReference
		}
	}{}
	err = json.Unmarshal([]byte(contents), &cache)
	c.loaded = true
	if err != nil || c.cacheVersion != cache.version {
		return nil
	}
	for filepathData, data := range cache.payload {
		c.cache[filepathData] = struct {
			hash      string
			reference *ast_map.FileReference
		}{hash: data.hash, reference: data.reference}
	}
	return nil
}

func (c *AstFileReferenceFileCache) Write() error {
	if !util.IsWriteable(filepath.Dir(c.cacheFile)) {
		return nil
	}
	cache := make(map[string]struct {
		hash      string
		reference *ast_map.FileReference
	})
	for filepathData, data := range c.cache {
		if _, ok := c.parsedFiles[filepathData]; ok {
			cache[filepathData] = data
		}
	}

	encoded, err := json.Marshal(struct {
		version string
		payload map[string]struct {
			hash      string
			reference *ast_map.FileReference
		}
	}{
		version: c.cacheVersion,
		payload: cache,
	})
	if err != nil {
		return err
	}

	err = os.WriteFile(c.cacheFile, encoded, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *AstFileReferenceFileCache) has(filepath string) (bool, error) {
	err := c.Load()
	if err != nil {
		return false, err
	}
	filepath = c.normalizeFilepath(filepath)
	if _, ok := c.cache[filepath]; !ok {
		return false, nil
	}
	hash, err := util.Sha1File(filepath)
	if err != nil {
		return false, err
	}
	if hash != c.cache[filepath].hash {
		delete(c.cache, filepath)
		return false, nil
	}

	return true, nil
}

func (c *AstFileReferenceFileCache) normalizeFilepath(path string) string {
	normalized, err := filepath.Abs(path)

	if err != nil {
		panic("File not exists: " + path)
	}

	return normalized
}
