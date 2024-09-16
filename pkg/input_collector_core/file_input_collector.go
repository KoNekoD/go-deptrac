package input_collector_core

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain_error_contract"
	"github.com/KoNekoD/go-deptrac/pkg/file_supportive/exception"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
	"path/filepath"
	"regexp"
)

type FileInputCollector struct {
	paths                []string
	excludedFilePatterns []string
	// todo: типы не берутся из типов полей структур
}

func NewFileInputCollector(originalPaths []string, excludedFilePatterns []string, basePath string) (InputCollectorInterface, error) {
	fileInfo, err := os.Stat(basePath)
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() || !util.IsReadable(basePath) {
		return nil, exception.NewInvalidPathExceptionUnreadablePath(fileInfo)
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
			return nil, exception.NewInvalidPathExceptionUnreadablePath(pathFileInfo)
		}
		paths = append(paths, util.PathCanonicalize(path))
	}

	return &FileInputCollector{paths: paths, excludedFilePatterns: excludedFilePatterns}, nil
}

func (c *FileInputCollector) Collect() ([]string, error) {
	if len(c.paths) == 0 {
		return nil, domain_error_contract.NewException("No 'paths' defined in the depfile.")
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
			return nil, NewInputExceptionCouldNotCollectFiles(err)
		}
	}
	return paths, nil
}
