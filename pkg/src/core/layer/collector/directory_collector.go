package collector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/ast"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/layer"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"regexp"
)

type DirectoryCollector struct {
	*RegexCollector
}

func NewDirectoryCollector() *DirectoryCollector {
	return &DirectoryCollector{
		RegexCollector: NewRegexCollector(),
	}
}

func (c *DirectoryCollector) Satisfy(config map[string]interface{}, reference ast.TokenReferenceInterface) (bool, error) {
	filepath := reference.GetFilepath()
	if filepath == nil {
		return false, nil
	}
	validatedPattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}
	normalizedPath := util.PathNormalize(*filepath)

	r, err := regexp.Compile(validatedPattern)
	if err != nil {
		return false, err
	}

	match := r.FindStringSubmatch(normalizedPath)

	return len(match) > 0, nil
}

func (c *DirectoryCollector) GetPattern(config map[string]interface{}) (string, error) {
	if _, ok := config["value"]; !ok {
		if _, ok2 := config["value"].(string); !ok2 {
			return "", layer.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("DirectoryCollector needs the regex configuration")
		}
	}

	// TODO: Убрать нахер везде #regex# и /regex/
	return fmt.Sprintf("%s", config["value"].(string)), nil
}
