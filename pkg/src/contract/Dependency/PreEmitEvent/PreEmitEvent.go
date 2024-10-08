package PreEmitEvent

// PreEmitEvent - Event triggered before all the dependencies have been resolved.
type PreEmitEvent struct {
	EmitterName string
}

func NewPreEmitEvent(emitterName string) *PreEmitEvent {
	return &PreEmitEvent{
		EmitterName: emitterName,
	}
}
