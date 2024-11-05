package events

// PostCreateAstMapEvent - Event triggered after the AST map of all files has been created.
type PostCreateAstMapEvent struct{}

func NewPostCreateAstMapEvent() *PostCreateAstMapEvent {
	return &PostCreateAstMapEvent{}
}
