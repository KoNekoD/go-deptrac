package dependency_resolver

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	config_contract2 "github.com/KoNekoD/go-deptrac/pkg/config_contract"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_contract"
	dependency_core2 "github.com/KoNekoD/go-deptrac/pkg/dependency_core"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_core/emitter"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/event_dispatcher/event_dispatcher_interface"
	"reflect"
)

type DependencyResolver struct {
	config               *config_contract2.AnalyserConfig
	inheritanceFlattener *dependency_core2.InheritanceFlattener
	emitterLocator       map[config_contract2.EmitterType]emitter.DependencyEmitterInterface
	eventDispatcher      event_dispatcher_interface.EventDispatcherInterface
}

func NewDependencyResolver(typesConfig *config_contract2.AnalyserConfig, emitterLocator map[config_contract2.EmitterType]emitter.DependencyEmitterInterface, inheritanceFlattener *dependency_core2.InheritanceFlattener, eventDispatcher event_dispatcher_interface.EventDispatcherInterface) *DependencyResolver {
	return &DependencyResolver{
		config:               typesConfig,
		emitterLocator:       emitterLocator,
		inheritanceFlattener: inheritanceFlattener,
		eventDispatcher:      eventDispatcher,
	}
}

func (r *DependencyResolver) Resolve(astMap *ast_map.AstMap) (*dependency_core2.DependencyList, error) {
	result := dependency_core2.NewDependencyList()

	for _, typeConfig := range r.config.Types {

		dependencyEmitterInterface, ok := r.emitterLocator[typeConfig]
		if !ok {
			return nil, dependency_core2.NewInvalidEmitterConfigurationExceptionCouldNotLocate(typeConfig)
		}

		if reflect.TypeOf(dependencyEmitterInterface).Kind() != reflect.Ptr {
			return nil, dependency_core2.NewInvalidEmitterConfigurationExceptionIsNotEmitter(typeConfig, dependencyEmitterInterface)
		}

		err := r.eventDispatcher.DispatchEvent(dependency_contract.NewPreEmitEvent(dependencyEmitterInterface.GetName()))
		if err != nil {
			return nil, err
		}

		dependencyEmitterInterface.ApplyDependencies(*astMap, result)

		errDispatchPostEmit := r.eventDispatcher.DispatchEvent(dependency_contract.NewPostEmitEvent())
		if errDispatchPostEmit != nil {
			return nil, errDispatchPostEmit
		}
	}

	errDispatchPreFlatten := r.eventDispatcher.DispatchEvent(dependency_contract.NewPreFlattenEvent())
	if errDispatchPreFlatten != nil {
		return nil, errDispatchPreFlatten
	}

	r.inheritanceFlattener.FlattenDependencies(*astMap, result)

	errDispatchPostFlatten := r.eventDispatcher.DispatchEvent(dependency_contract.NewPostFlattenEvent())
	if errDispatchPostFlatten != nil {
		return nil, errDispatchPostFlatten
	}

	return result, nil
}
