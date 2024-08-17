package PreFlattenEvent

// PreFlattenEvent - Event triggered before all the dependencies have been flattened. This occurs when all dependencies caused by class inheritance have been resolved.
type PreFlattenEvent struct{}

func NewPreFlattenEvent() *PreFlattenEvent {
	return &PreFlattenEvent{}
}
