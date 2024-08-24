package dependency_resolver

import (
	"github.com/KoNekoD/go-deptrac/pkg/contract/config"
	contractDependency "github.com/KoNekoD/go-deptrac/pkg/contract/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/core/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/core/dependency/emitter"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/dependency_injection/event_dispatcher/event_dispatcher_interface"
	"reflect"
)

type DependencyResolver struct {
	config               *config.AnalyserConfig
	inheritanceFlattener *dependency.InheritanceFlattener
	emitterLocator       map[config.EmitterType]emitter.DependencyEmitterInterface
	eventDispatcher      util.EventDispatcherInterface
}

func NewDependencyResolver(typesConfig *config.AnalyserConfig, emitterLocator map[config.EmitterType]emitter.DependencyEmitterInterface, inheritanceFlattener *dependency.InheritanceFlattener, eventDispatcher util.EventDispatcherInterface) *DependencyResolver {
	return &DependencyResolver{
		config:               typesConfig,
		emitterLocator:       emitterLocator,
		inheritanceFlattener: inheritanceFlattener,
		eventDispatcher:      eventDispatcher,
	}
}

func (r *DependencyResolver) Resolve(astMap *ast_map.AstMap) (*dependency.DependencyList, error) {
	result := dependency.NewDependencyList()

	for _, typeConfig := range r.config.Types {

		dependencyEmitterInterface, ok := r.emitterLocator[typeConfig]
		if !ok {
			return nil, dependency.NewInvalidEmitterConfigurationExceptionCouldNotLocate(typeConfig)
		}

		if reflect.TypeOf(dependencyEmitterInterface).Kind() != reflect.Ptr {
			return nil, dependency.NewInvalidEmitterConfigurationExceptionIsNotEmitter(typeConfig, dependencyEmitterInterface)
		}

		err := r.eventDispatcher.DispatchEvent(contractDependency.NewPreEmitEvent(dependencyEmitterInterface.GetName()))
		if err != nil {
			return nil, err
		}

		dependencyEmitterInterface.ApplyDependencies(*astMap, result)

		errDispatchPostEmit := r.eventDispatcher.DispatchEvent(contractDependency.NewPostEmitEvent())
		if errDispatchPostEmit != nil {
			return nil, errDispatchPostEmit
		}
	}

	errDispatchPreFlatten := r.eventDispatcher.DispatchEvent(contractDependency.NewPreFlattenEvent())
	if errDispatchPreFlatten != nil {
		return nil, errDispatchPreFlatten
	}

	r.inheritanceFlattener.FlattenDependencies(*astMap, result)

	errDispatchPostFlatten := r.eventDispatcher.DispatchEvent(contractDependency.NewPostFlattenEvent())
	if errDispatchPostFlatten != nil {
		return nil, errDispatchPostFlatten
	}

	return result, nil
}
