package collectors_shared

import (
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	filepath2 "path/filepath"
	"regexp"
)

type GlobCollector struct {
	*RegexCollector
	basePath string
}

func NewGlobCollector(basePath string) *GlobCollector {
	return &GlobCollector{
		RegexCollector: NewRegexCollector(),
		basePath:       utils.PathNormalize(basePath),
	}
}

func (c *GlobCollector) Satisfy(config map[string]interface{}, reference tokens_references.TokenReferenceInterface) (bool, error) {
	filepath := reference.GetFilepath()

	if filepath == nil {
		return false, nil
	}

	validatedPattern, err := c.GetValidatedPattern(config, c.GetPattern)
	if err != nil {
		return false, err
	}

	normalizedPath := utils.PathNormalize(*filepath)

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
	if !utils.MapKeyExists(config, "value") || !utils.MapKeyIsString(config, "value") {
		return "", apperrors.NewInvalidCollectorDefinitionInvalidCollectorConfiguration("GlobCollector needs the glob pattern configuration.")
	}

	return utils.GlogToRegex(config["value"].(string)), nil
}
