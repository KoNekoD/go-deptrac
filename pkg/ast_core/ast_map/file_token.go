package ast_map

import (
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
	"strings"
)

type FileToken struct {
	path string
}

func NewFileToken(path *string) *FileToken {
	return &FileToken{path: util.PathNormalize(*path)}
}

func (t *FileToken) ToString() string {
	wd, err := os.Getwd()

	if err != nil {
		return t.path
	}

	wd = util.PathNormalize(wd)

	if strings.HasPrefix(t.path, wd) {
		return strings.TrimPrefix(t.path, wd)
	}

	return t.path
}
