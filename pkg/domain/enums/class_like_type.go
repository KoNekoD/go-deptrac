package enums

type ClassLikeType string

const (
	TypeClasslike ClassLikeType = "classLike"
	TypeClass     ClassLikeType = "class"
	TypeInterface ClassLikeType = "interface"
	TypeTrait     ClassLikeType = "trait"
)

func (t ClassLikeType) ToString() string {
	return string(t)
}
