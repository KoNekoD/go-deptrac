package PostEmitEvent

// PostEmitEvent - Event triggered after all the dependencies have been resolved.
type PostEmitEvent struct{}

func NewPostEmitEvent() *PostEmitEvent {
	return &PostEmitEvent{}
}
