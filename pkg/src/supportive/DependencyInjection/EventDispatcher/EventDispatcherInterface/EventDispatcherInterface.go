package util

type EventDispatcherInterface interface {
	DispatchEvent(event interface{}) error
}
