package AstMap

// TaggedTokenReference - Helper trait for implementing TaggedTokenReferenceInterface.
type TaggedTokenReference struct {
	Tags map[string][]string
}

func NewTaggedTokenReference(tags map[string][]string) *TaggedTokenReference {
	return &TaggedTokenReference{
		Tags: tags,
	}
}

func (a *TaggedTokenReference) HasTag(name string) bool {
	if a.Tags == nil {
		return false
	}

	_, ok := a.Tags[name]
	return ok
}

func (a *TaggedTokenReference) GetTagLines(name string) []string {
	if a.Tags == nil {
		return nil
	}

	tags, ok := a.Tags[name]

	if !ok {
		return nil
	}

	return tags
}
