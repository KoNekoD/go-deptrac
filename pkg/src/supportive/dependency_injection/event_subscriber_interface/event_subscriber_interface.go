package event_subscriber_interface

type EventSubscriberInterface interface {
	InvokeEventSubscriber(rawEvent interface{}, stopPropagation func()) error
}
