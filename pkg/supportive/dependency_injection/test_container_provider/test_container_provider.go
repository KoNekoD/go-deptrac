package test_container_provider

import (
	"github.com/KoNekoD/go-deptrac/pkg/supportive/console/application"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/dependency_injection"
	"os"
	"reflect"
)

func ProvideTestContainer() *dependency_injection.ServiceContainerBuilder {
	currentWorkingDirectory, _ := os.Getwd()
	config := currentWorkingDirectory + application.DirectorySeparator + "deptrac.yaml"
	cache := ""

	factory := dependency_injection.NewServiceContainerBuilder(currentWorkingDirectory)
	factory, _ = factory.WithConfig(config)

	_, _ = factory.Build(&cache, false)

	return factory
}

func ProvideTestContainerService(containerProperty string) interface{} {
	container := ProvideTestContainer().GetContainer()

	refProps := reflect.ValueOf(container).Elem().FieldByName(containerProperty)

	return refProps.Interface()
}
