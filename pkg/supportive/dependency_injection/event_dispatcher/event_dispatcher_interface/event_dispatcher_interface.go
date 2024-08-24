package event_dispatcher_interface

type EventDispatcherInterface interface {
	DispatchEvent(event interface{}) error
}
