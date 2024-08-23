package dependency

// PostFlattenEvent - Event triggered after all the dependencies have been flattened. This occurs when all dependencies caused by class inheritance have been resolved.
type PostFlattenEvent struct{}

func NewPostFlattenEvent() *PostFlattenEvent {
	return &PostFlattenEvent{}
}
