package di

import (
	"os"
	"reflect"
)

var currentWorkingDirectory, _ = os.Getwd()

var TestConfigFile = currentWorkingDirectory + "/deptrac.yaml"

func UseVoidConfig() {
	TestConfigFile = "resources/deptrac-empty.yaml"
}

func ProvideTestContainer() *ServiceContainerBuilder {
	cache := ""
	factory := NewServiceContainerBuilder(currentWorkingDirectory)
	factory, _ = factory.WithConfig(TestConfigFile)

	_, _ = factory.Build(&cache, false)

	return factory
}

func ProvideTestContainerService(containerProperty string) interface{} {
	container := ProvideTestContainer().GetContainer()

	refProps := reflect.ValueOf(container).Elem().FieldByName(containerProperty)

	return refProps.Interface()
}
