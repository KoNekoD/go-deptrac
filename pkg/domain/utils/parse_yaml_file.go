package utils

import (
	_ "github.com/KoNekoD/go-deptrac/resources"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
)

func ParseYamlFile(file string) (map[string]interface{}, error) {
	output := make(map[string]interface{})

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "could not read yaml file_supportive")
	}
	err = yaml.Unmarshal(yamlFile, &output)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal yaml file_supportive")
	}

	return output, nil
}
