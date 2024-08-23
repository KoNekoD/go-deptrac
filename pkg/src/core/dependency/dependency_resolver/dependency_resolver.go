package dependency_resolver

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/AnalyserConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/EmitterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/dependency/emitter"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
	"reflect"
)

type DependencyResolver struct {
	config               *AnalyserConfig.AnalyserConfig
	inheritanceFlattener *dependency.InheritanceFlattener
	emitterLocator       map[EmitterType.EmitterType]emitter.DependencyEmitterInterface
	eventDispatcher      util.EventDispatcherInterface
}

func NewDependencyResolver(typesConfig *AnalyserConfig.AnalyserConfig, emitterLocator map[EmitterType.EmitterType]emitter.DependencyEmitterInterface, inheritanceFlattener *dependency.InheritanceFlattener, eventDispatcher util.EventDispatcherInterface) *DependencyResolver {
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

		err := r.eventDispatcher.DispatchEvent(PreEmitEvent.NewPreEmitEvent(dependencyEmitterInterface.GetName()))
		if err != nil {
			return nil, err
		}

		dependencyEmitterInterface.ApplyDependencies(*astMap, result)

		errDispatchPostEmit := r.eventDispatcher.DispatchEvent(PostEmitEvent.NewPostEmitEvent())
		if errDispatchPostEmit != nil {
			return nil, errDispatchPostEmit
		}
	}

	errDispatchPreFlatten := r.eventDispatcher.DispatchEvent(PreFlattenEvent.NewPreFlattenEvent())
	if errDispatchPreFlatten != nil {
		return nil, errDispatchPreFlatten
	}

	r.inheritanceFlattener.FlattenDependencies(*astMap, result)

	errDispatchPostFlatten := r.eventDispatcher.DispatchEvent(PostFlattenEvent.NewPostFlattenEvent())
	if errDispatchPostFlatten != nil {
		return nil, errDispatchPostFlatten
	}

	return result, nil
}
