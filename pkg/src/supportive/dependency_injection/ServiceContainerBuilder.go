package dependency_injection

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/config/deptrac_config"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/Exception/CannotLoadConfiguration"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/container_builder"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/dependency_injection/event_subscriber_interface_map"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
	"path/filepath"
	"reflect"
)

type ServiceContainerBuilder struct {
	workingDirectory string
	containerBuilder *container_builder.ContainerBuilder
	configFile       *string
	cacheFile        *string
}

func NewServiceContainerBuilder(workingDirectory string) *ServiceContainerBuilder {
	return &ServiceContainerBuilder{
		workingDirectory: workingDirectory,
	}
}

func (b *ServiceContainerBuilder) WithConfig(configFile string) (*ServiceContainerBuilder, error) {
	if configFile == "" {
		return b, nil
	}

	if !filepath.IsAbs(configFile) {
		absolutePath, err := filepath.Abs(filepath.Join(b.workingDirectory, configFile))
		if err != nil {
			return nil, err
		}
		configFile = absolutePath
	}

	b.configFile = &configFile
	return b, nil
}

func (b *ServiceContainerBuilder) withCache(cacheFile string) error {
	if !filepath.IsAbs(cacheFile) {
		absolutePath, err := filepath.Abs(filepath.Join(b.workingDirectory, cacheFile))
		if err != nil {
			return err
		}
		cacheFile = absolutePath
	}

	b.cacheFile = &cacheFile
	return nil
}

func (b *ServiceContainerBuilder) clearCache(cacheFile string) error {
	if !filepath.IsAbs(cacheFile) {
		absolutePath, err := filepath.Abs(filepath.Join(b.workingDirectory, cacheFile))
		if err != nil {
			return err
		}
		cacheFile = absolutePath
	}
	return os.Remove(cacheFile)
}

func (b *ServiceContainerBuilder) Build(cacheOverride *string, clearCache bool) (*container_builder.ContainerBuilder, error) {
	container := container_builder.NewContainerBuilder(b.workingDirectory)
	b.containerBuilder = container

	if b.configFile != nil {
		if err := b.loadConfiguration(container, *b.configFile); err != nil {
			return nil, err
		}
	}

	err := loadServices(container, b.cacheFile)
	if err != nil {
		return nil, err
	}

	// Debug event subscriber
	if b.containerBuilder.DebugBoolFlag != nil && *b.containerBuilder.DebugBoolFlag == true {
		for _, key := range event_subscriber_interface_map.Map.Keys() {
			mapByPriorities, _ := event_subscriber_interface_map.Map.Get(key)

			for _, priority := range mapByPriorities.Keys() {
				subscribers, _ := mapByPriorities.Get(priority)

				for _, subscriber := range subscribers {
					subscriberName := reflect.TypeOf(subscriber).String()

					fmt.Println(key, priority, subscriberName)
				}
			}
		}
	}

	return container, nil
}

func (b *ServiceContainerBuilder) GetContainer() *container_builder.ContainerBuilder {
	return b.containerBuilder
}

func loadServices(container *container_builder.ContainerBuilder, cacheFile *string) error {
	if cacheFile != nil {
		container.CacheFile = cacheFile
		Cache(container)
	}

	err := Services(container)
	if err != nil {
		return err
	}

	return nil
}

func (b *ServiceContainerBuilder) loadConfiguration(container *container_builder.ContainerBuilder, configFile string) *CannotLoadConfiguration.CannotLoadConfiguration {
	projectDirectory := filepath.Dir(configFile)
	if projectDirectory == "" {
		return CannotLoadConfiguration.NewCannotLoadConfigurationFromConfig(configFile, "Unable to load configuration: Invalid or missing path.")
	}

	container.ProjectDirectory = projectDirectory

	parsed, err := util.ParseYamlFile(configFile)

	if err != nil {
		return CannotLoadConfiguration.NewCannotLoadConfigurationFromConfig(configFile, err.Error())
	}

	deptracConfig, err := deptrac_config.NewDeptracConfig(parsed)

	if err != nil {
		return CannotLoadConfiguration.NewCannotLoadConfigurationFromConfig(configFile, err.Error())
	}

	b.containerBuilder.Configuration = deptracConfig

	return nil
}
