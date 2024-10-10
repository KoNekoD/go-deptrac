package dispatchers

type EventDispatcherInterface interface {
	DispatchEvent(event interface{}) error
}
