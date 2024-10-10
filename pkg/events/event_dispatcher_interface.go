package events

type EventDispatcherInterface interface {
	DispatchEvent(event interface{}) error
}
