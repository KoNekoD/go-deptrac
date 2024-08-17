package DependencyResolver

import (
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/AnalyserConfig"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Config/EmitterType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PostFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreEmitEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Dependency/PreFlattenEvent"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/DependencyList"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/Emitter/DependencyEmitterInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/InheritanceFlattener"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Dependency/InvalidEmitterConfigurationException"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/DependencyInjection/EventDispatcher/EventDispatcherInterface"
	"reflect"
)

type DependencyResolver struct {
	config               *AnalyserConfig.AnalyserConfig
	inheritanceFlattener *InheritanceFlattener.InheritanceFlattener
	emitterLocator       map[EmitterType.EmitterType]DependencyEmitterInterface.DependencyEmitterInterface
	eventDispatcher      util.EventDispatcherInterface
}

func NewDependencyResolver(typesConfig *AnalyserConfig.AnalyserConfig, emitterLocator map[EmitterType.EmitterType]DependencyEmitterInterface.DependencyEmitterInterface, inheritanceFlattener *InheritanceFlattener.InheritanceFlattener, eventDispatcher util.EventDispatcherInterface) *DependencyResolver {
	return &DependencyResolver{
		config:               typesConfig,
		emitterLocator:       emitterLocator,
		inheritanceFlattener: inheritanceFlattener,
		eventDispatcher:      eventDispatcher,
	}
}

func (r *DependencyResolver) Resolve(astMap *AstMap.AstMap) (*DependencyList.DependencyList, error) {
	result := DependencyList.NewDependencyList()

	for _, typeConfig := range r.config.Types {

		emitter, ok := r.emitterLocator[typeConfig]
		if !ok {
			return nil, InvalidEmitterConfigurationException.NewInvalidEmitterConfigurationExceptionCouldNotLocate(typeConfig)
		}

		if reflect.TypeOf(emitter).Kind() != reflect.Ptr {
			return nil, InvalidEmitterConfigurationException.NewInvalidEmitterConfigurationExceptionIsNotEmitter(typeConfig, emitter)
		}

		err := r.eventDispatcher.DispatchEvent(PreEmitEvent.NewPreEmitEvent(emitter.GetName()))
		if err != nil {
			return nil, err
		}

		emitter.ApplyDependencies(*astMap, result)

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
