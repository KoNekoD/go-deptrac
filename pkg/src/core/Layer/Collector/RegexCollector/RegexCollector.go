package RegexCollector

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Layer/InvalidCollectorDefinitionException"
	"regexp"
)

type RegexCollector struct{}

func NewRegexCollector() *RegexCollector {
	return &RegexCollector{}
}

func (c *RegexCollector) GetValidatedPattern(config map[string]interface{}, getPattern func(config map[string]interface{}) (string, error)) (string, error) {
	pattern, err := getPattern(config)
	if err != nil {
		return "", err
	}

	if _, err = regexp.Compile(pattern); err != nil {
		return "", InvalidCollectorDefinitionException.NewInvalidCollectorDefinitionExceptionInvalidCollectorConfiguration("Invalid regex pattern " + pattern)
	}

	return pattern, nil
}
