package apperrors

import (
	"fmt"
	"os"
)

type CacheFileException struct {
	Message string
}

func (e *CacheFileException) Error() string {
	return e.Message
}

func newCacheFileException(message string) *CacheFileException {
	return &CacheFileException{Message: message}
}

func NewCacheFileExceptionNotWritable(cacheFile *os.File) *CacheFileException {
	return newCacheFileException(fmt.Sprintf("Cache file_supportive \"%s\" is not writable.", cacheFile.Name()))
}
