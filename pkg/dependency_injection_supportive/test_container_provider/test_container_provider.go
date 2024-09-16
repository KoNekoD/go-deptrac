package test_container_provider

import (
	"github.com/KoNekoD/go-deptrac/pkg/console_supportive/application"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive"
	"os"
	"reflect"
)

func ProvideTestContainer() *dependency_injection_supportive.ServiceContainerBuilder {
	currentWorkingDirectory, _ := os.Getwd()
	config := currentWorkingDirectory + application.DirectorySeparator + "deptrac.yaml"
	cache := ""

	factory := dependency_injection_supportive.NewServiceContainerBuilder(currentWorkingDirectory)
	factory, _ = factory.WithConfig(config)

	_, _ = factory.Build(&cache, false)

	return factory
}

func ProvideTestContainerService(containerProperty string) interface{} {
	container := ProvideTestContainer().GetContainer()

	refProps := reflect.ValueOf(container).Elem().FieldByName(containerProperty)

	return refProps.Interface()
}
