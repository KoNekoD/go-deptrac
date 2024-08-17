package DirectoryCollector

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/RegexCollector"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"regexp"
)

type DirectoryCollector struct {
	*RegexCollector.RegexCollector
}

func NewDirectoryCollector() *DirectoryCollector {
	return &DirectoryCollector{
		RegexCollector: RegexCollector.NewRegexCollector(),
	}
}

func (c *DirectoryCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
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
			return "", InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("DirectoryCollector needs the regex configuration")
		}
	}

	// TODO: Убрать нахер везде #regex# и /regex/
	return fmt.Sprintf("%s", config["value"].(string)), nil
}
