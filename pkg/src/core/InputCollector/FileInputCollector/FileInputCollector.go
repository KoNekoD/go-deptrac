package FileInputCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ExceptionInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/InputCollector/InputCollectorInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/InputCollector/InputException"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/File/Exception/InvalidPathException"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
	"path/filepath"
	"regexp"
)

type FileInputCollector struct {
	paths                []string
	excludedFilePatterns []string
}

func NewFileInputCollector(originalPaths []string, excludedFilePatterns []string, basePath string) (InputCollectorInterface.InputCollectorInterface, error) {
	fileInfo, err := os.Stat(basePath)
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() || !util.IsReadable(basePath) {
		return nil, InvalidPathException.NewInvalidPathExceptionUnreadablePath(fileInfo)
	}
	paths := make([]string, 0)
	for _, originalPath := range originalPaths {
		var path string
		if filepath.IsAbs(path) {
			path = originalPath
		} else {
			path = filepath.Join(basePath, originalPath)
		}

		if !util.IsReadable(path) {
			pathFileInfo, err := os.Stat(path)
			if err != nil {
				return nil, err
			}
			return nil, InvalidPathException.NewInvalidPathExceptionUnreadablePath(pathFileInfo)
		}
		paths = append(paths, util.PathCanonicalize(path))
	}

	return &FileInputCollector{paths: paths, excludedFilePatterns: excludedFilePatterns}, nil
}

func (c *FileInputCollector) Collect() ([]string, error) {
	if len(c.paths) == 0 {
		return nil, ExceptionInterface.NewException("No 'paths' defined in the depfile.")
	}

	regex, err := regexp.Compile(".*\\.go")
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0)
	for _, path := range c.paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err == nil && regex.MatchString(info.Name()) {
				paths = append(paths, path)
			}

			return err
		})
		if err != nil {
			return nil, InputException.NewInputExceptionCouldNotCollectFiles(err)
		}
	}
	return paths, nil
}
