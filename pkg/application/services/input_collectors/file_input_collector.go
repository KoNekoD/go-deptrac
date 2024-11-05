package input_collectors

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"os"
	"path/filepath"
	"regexp"
)

type FileInputCollector struct {
	paths                []string
	excludedFilePatterns []string
	// todo: типы не берутся из типов полей структур
}

func NewFileInputCollector(originalPaths []string, excludedFilePatterns []string, basePath string) (InputCollector, error) {
	fileInfo, err := os.Stat(basePath)
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() || !utils.IsReadable(basePath) {
		return nil, apperrors.NewInvalidPathExceptionUnreadablePath(fileInfo)
	}
	paths := make([]string, 0)
	for _, originalPath := range originalPaths {
		var path string
		if filepath.IsAbs(path) {
			path = originalPath
		} else {
			path = filepath.Join(basePath, originalPath)
		}

		if !utils.IsReadable(path) {
			pathFileInfo, err := os.Stat(path)
			if err != nil {
				return nil, err
			}
			return nil, apperrors.NewInvalidPathExceptionUnreadablePath(pathFileInfo)
		}
		paths = append(paths, utils.PathCanonicalize(path))
	}

	return &FileInputCollector{paths: paths, excludedFilePatterns: excludedFilePatterns}, nil
}

func (c *FileInputCollector) Collect() ([]string, error) {
	if len(c.paths) == 0 {
		return nil, apperrors.NewException("No 'paths' defined in the depfile.")
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
			return nil, apperrors.NewInputExceptionCouldNotCollectFiles(err)
		}
	}
	return paths, nil
}
