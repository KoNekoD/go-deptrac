package ast_file_reference_cache

import (
	"encoding/json"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"os"
	"path/filepath"
)

type AstFileReferenceFileCache struct {
	cache map[string]struct {
		hash      string
		reference *tokens_references.FileReference
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
			reference *tokens_references.FileReference
		}),
		loaded:       false,
		parsedFiles:  make(map[string]bool),
		cacheFile:    cacheFile,
		cacheVersion: cacheVersion,
	}
}

func (c *AstFileReferenceFileCache) Get(filepath string) (*tokens_references.FileReference, error) {
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

func (c *AstFileReferenceFileCache) Set(fileReference *tokens_references.FileReference) error {
	err := c.Load()
	if err != nil {
		return err
	}
	normalizedFilepath := c.normalizeFilepath(*fileReference.Filepath)
	c.parsedFiles[normalizedFilepath] = true

	hash, err := utils.Sha1File(normalizedFilepath)
	if err != nil {
		return err
	}

	c.cache[normalizedFilepath] = struct {
		hash      string
		reference *tokens_references.FileReference
	}{hash: hash, reference: fileReference}

	return nil
}

func (c *AstFileReferenceFileCache) Load() error {
	if c.loaded {
		return nil
	}
	if !utils.FileExists(c.cacheFile) || !utils.IsReadable(c.cacheFile) {
		return nil
	}
	contents, err := utils.FileReaderRead(c.cacheFile)
	if err != nil {
		return err
	}
	cache := struct {
		version string
		payload map[string]struct {
			hash      string
			reference *tokens_references.FileReference
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
			reference *tokens_references.FileReference
		}{hash: data.hash, reference: data.reference}
	}
	return nil
}

func (c *AstFileReferenceFileCache) Write() error {
	if !utils.IsWriteable(filepath.Dir(c.cacheFile)) {
		return nil
	}
	cache := make(map[string]struct {
		hash      string
		reference *tokens_references.FileReference
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
			reference *tokens_references.FileReference
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
	hash, err := utils.Sha1File(filepath)
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
