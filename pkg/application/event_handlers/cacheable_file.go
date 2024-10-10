package event_handlers

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/services/ast_file_reference_cache"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
)

type CacheableFile struct {
	deferredCache ast_file_reference_cache.AstFileReferenceDeferredCacheInterface
}

func NewCacheableFile(deferredCache ast_file_reference_cache.AstFileReferenceDeferredCacheInterface) *CacheableFile {
	return &CacheableFile{deferredCache: deferredCache}
}

func (s *CacheableFile) HandleEvent(rawEvent interface{}, stopPropagation func()) error {
	switch rawEvent.(type) {
	case *events.PreCreateAstMapEvent:
		err := s.deferredCache.Load()
		if err != nil {
			return err
		}
	case *events.PostCreateAstMapEvent:
		err := s.deferredCache.Write()
		if err != nil {
			return err
		}
	}

	return nil
}
