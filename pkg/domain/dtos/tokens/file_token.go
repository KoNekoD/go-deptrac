package tokens

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"os"
	"strings"
)

type FileToken struct {
	path string
}

func NewFileToken(path *string) *FileToken {
	return &FileToken{path: utils.PathNormalize(*path)}
}

func (t *FileToken) ToString() string {
	wd, err := os.Getwd()

	if err != nil {
		return t.path
	}

	wd = utils.PathNormalize(wd)

	if strings.HasPrefix(t.path, wd) {
		return strings.TrimPrefix(t.path, wd)
	}

	return t.path
}

func (t *FileToken) tokenInterface() {}
