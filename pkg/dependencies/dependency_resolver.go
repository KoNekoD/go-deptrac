package dependencies

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/configs"
	"github.com/KoNekoD/go-deptrac/pkg/domain/apperrors"
	"github.com/KoNekoD/go-deptrac/pkg/emitters"
	"github.com/KoNekoD/go-deptrac/pkg/events"
	"github.com/KoNekoD/go-deptrac/pkg/flatteners"
	"reflect"
)

type DependencyResolver struct {
	config               *configs.AnalyserConfig
	inheritanceFlattener *flatteners.InheritanceFlattener
	emitterLocator       map[emitters.EmitterType]emitters.DependencyEmitterInterface
	eventDispatcher      events.EventDispatcherInterface
}

func NewDependencyResolver(typesConfig *configs.AnalyserConfig, emitterLocator map[emitters.EmitterType]emitters.DependencyEmitterInterface, inheritanceFlattener *flatteners.InheritanceFlattener, eventDispatcher events.EventDispatcherInterface) *DependencyResolver {
	return &DependencyResolver{
		config:               typesConfig,
		emitterLocator:       emitterLocator,
		inheritanceFlattener: inheritanceFlattener,
		eventDispatcher:      eventDispatcher,
	}
}

func (r *DependencyResolver) Resolve(astMap *ast_map.AstMap) (*DependencyList, error) {
	result := NewDependencyList()

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

	errDispatchPreFlatten := r.eventDispatcher.DispatchEvent(flatteners.NewPreFlattenEvent())
	if errDispatchPreFlatten != nil {
		return nil, errDispatchPreFlatten
	}

	r.inheritanceFlattener.FlattenDependencies(*astMap, result)

	errDispatchPostFlatten := r.eventDispatcher.DispatchEvent(flatteners.NewPostFlattenEvent())
	if errDispatchPostFlatten != nil {
		return nil, errDispatchPostFlatten
	}

	return result, nil
}
