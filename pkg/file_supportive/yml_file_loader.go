package file_supportive

import (
	exception2 "github.com/KoNekoD/go-deptrac/pkg/file_supportive/exception"
	"gopkg.in/yaml.v3"
)

type YmlFileLoader struct{}

func NewYmlFileLoader() *YmlFileLoader {
	return &YmlFileLoader{}
}

type ParseFileResult struct {
	Parameters map[string]interface{} `yaml:"parameters"`
	Services   map[string]interface{} `yaml:"services"`
	Imports    []string               `yaml:"imports"`
}

func (y *YmlFileLoader) ParseFile(file string) (*ParseFileResult, error) {
	yamlMap := make(map[string]interface{})

	err := yaml.Unmarshal([]byte(file), &yamlMap)
	if err != nil {
		return nil, exception2.NewFileCannotBeParsedAsYamlExceptionFromFilenameAndException(file, err)
	}

	_, ok1 := yamlMap["parameters"]
	_, ok2 := yamlMap["services"]
	_, ok3 := yamlMap["imports"]

	if !ok1 || !ok2 || !ok3 {
		return nil, exception2.NewParsedYamlIsNotAnArrayExceptionFromFilename(file)
	}

	result := &ParseFileResult{
		Parameters: yamlMap["parameters"].(map[string]interface{}),
		Services:   yamlMap["services"].(map[string]interface{}),
		Imports:    yamlMap["imports"].([]string),
	}

	return result, nil
}
