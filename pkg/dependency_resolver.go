package pkg

import (
	"github.com/KoNekoD/go-deptrac/pkg/dispatchers"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/dependencies"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	events2 "github.com/KoNekoD/go-deptrac/pkg/domain/events"
	"github.com/KoNekoD/go-deptrac/pkg/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/flatteners"
	"reflect"
)

type DependencyResolver struct {
	config               *dtos.AnalyserConfig
	inheritanceFlattener *flatteners.InheritanceFlattener
	emitterLocator       map[enums.EmitterType]emitters.DependencyEmitterInterface
	eventDispatcher      dispatchers.EventDispatcherInterface
}

func NewDependencyResolver(typesConfig *dtos.AnalyserConfig, emitterLocator map[enums.EmitterType]emitters.DependencyEmitterInterface, inheritanceFlattener *flatteners.InheritanceFlattener, eventDispatcher dispatchers.EventDispatcherInterface) *DependencyResolver {
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

		err := r.eventDispatcher.DispatchEvent(emitters.NewPreEmitEvent(dependencyEmitterInterface.GetName()))
		if err != nil {
			return nil, err
		}

		dependencyEmitterInterface.ApplyDependencies(*astMap, result)

		errDispatchPostEmit := r.eventDispatcher.DispatchEvent(emitters.NewPostEmitEvent())
		if errDispatchPostEmit != nil {
			return nil, errDispatchPostEmit
		}
	}

	errDispatchPreFlatten := r.eventDispatcher.DispatchEvent(events2.NewPreFlattenEvent())
	if errDispatchPreFlatten != nil {
		return nil, errDispatchPreFlatten
	}

	r.inheritanceFlattener.FlattenDependencies(*astMap, result)

	errDispatchPostFlatten := r.eventDispatcher.DispatchEvent(events2.NewPostFlattenEvent())
	if errDispatchPostFlatten != nil {
		return nil, errDispatchPostFlatten
	}

	return result, nil
}
