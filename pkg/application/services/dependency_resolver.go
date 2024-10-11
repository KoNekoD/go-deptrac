package services

import (
	"github.com/KoNekoD/go-deptrac/pkg/application/event_dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/application/services/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"reflect"
)

type DependencyResolver struct {
	config               *configs.AnalyserConfig
	inheritanceFlattener *InheritanceFlattener
	emitterLocator       map[enums.EmitterType]emitters.DependencyEmitterInterface
	eventDispatcher      event_dispatchers.EventDispatcherInterface
}

func NewDependencyResolver(typesConfig *configs.AnalyserConfig, emitterLocator map[enums.EmitterType]emitters.DependencyEmitterInterface, inheritanceFlattener *InheritanceFlattener, eventDispatcher event_dispatchers.EventDispatcherInterface) *DependencyResolver {
	return &DependencyResolver{
		config:               typesConfig,
		emitterLocator:       emitterLocator,
		inheritanceFlattener: inheritanceFlattener,
		eventDispatcher:      eventDispatcher,
	}
}

func (r *DependencyResolver) Resolve(astMap *ast_map.AstMap) (*dependencies.DependencyList, error) {
	result := dependencies.NewDependencyList()

	for _, typeConfig := range r.config.Types {
		dependencyEmitterInterface, ok := r.emitterLocator[typeConfig]
		if !ok {
			return nil, apperrors.NewInvalidEmitterConfigurationExceptionCouldNotLocate(string(typeConfig))
		}

		if reflect.TypeOf(dependencyEmitterInterface).Kind() != reflect.Ptr {
			return nil, apperrors.NewInvalidEmitterConfigurationExceptionIsNotEmitter(string(typeConfig), dependencyEmitterInterface)
		}

		err := r.eventDispatcher.DispatchEvent(events.NewPreEmitEvent(dependencyEmitterInterface.GetName()))
		if err != nil {
			return nil, err
		}

		dependencyEmitterInterface.ApplyDependencies(*astMap, result)

		errDispatchPostEmit := r.eventDispatcher.DispatchEvent(events.NewPostEmitEvent())
		if errDispatchPostEmit != nil {
			return nil, errDispatchPostEmit
		}
	}

	errDispatchPreFlatten := r.eventDispatcher.DispatchEvent(events.NewPreFlattenEvent())
	if errDispatchPreFlatten != nil {
		return nil, errDispatchPreFlatten
	}

	r.inheritanceFlattener.FlattenDependencies(*astMap, result)

	errDispatchPostFlatten := r.eventDispatcher.DispatchEvent(events.NewPostFlattenEvent())
	if errDispatchPostFlatten != nil {
		return nil, errDispatchPostFlatten
	}

	return result, nil
}
