package subscribers

type EventSubscriberInterface interface {
	InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error
}
