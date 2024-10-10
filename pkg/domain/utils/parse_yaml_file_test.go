package utils

import (
	_ "github.com/KoNekoD/go-deptrac/resources"
	"testing"
)

func TestParseYamlFile(t *testing.T) {
	data, err := ParseYamlFile("pkg/util/parse_yaml_file_test_1.yaml")

	if err != nil {
		panic(err)
	}

	_, ok := data["test"]
	if !ok {
		panic("Key 'test' not found in the yaml file_supportive")
	}
}
