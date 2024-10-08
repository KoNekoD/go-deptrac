package GlobCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenReferenceInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Layer/Collector/RegexCollector"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	filepath2 "path/filepath"
	"regexp"
)

type GlobCollector struct {
	*RegexCollector.RegexCollector
	basePath string
}

func NewGlobCollector(basePath string) *GlobCollector {
	return &GlobCollector{
		RegexCollector: RegexCollector.NewRegexCollector(),
		basePath:       util.PathNormalize(basePath),
	}
}

func (c *GlobCollector) Satisfy(config map[string]interface{}, reference TokenReferenceInterface.TokenReferenceInterface) (bool, error) {
	filepath := reference.GetFilepath()

	if filepath == nil {
		return false, nil
	}

	validatedPattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	normalizedPath := util.PathNormalize(*filepath)

	relativeFilePath, err := filepath2.Rel(c.basePath, normalizedPath)
	if err != nil {
		return false, err
	}

	r, err := regexp.Compile(validatedPattern)
	if err != nil {
		return false, err
	}

	return r.MatchString(relativeFilePath), nil
}

func (c *GlobCollector) GetPattern(config map[string]interface{}) (string, error) {
	if !util.MapKeyExists(config, "value") || !util.MapKeyIsString(config, "value") {
		return "", InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("GlobCollector needs the glob pattern configuration.")
	}

	return util.GlogToRegex(config["value"].(string)), nil
}
