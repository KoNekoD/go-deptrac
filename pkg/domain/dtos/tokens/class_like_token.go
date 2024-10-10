package tokens

import (
	"regexp"
	"strings"
)

type ClassLikeToken struct {
	ClassName string
}

func NewClassLikeToken(className string) *ClassLikeToken {
	return &ClassLikeToken{ClassName: className}
}

func NewClassLikeTokenFromFQCN(className string) *ClassLikeToken {
	return &ClassLikeToken{ClassName: strings.TrimPrefix(className, "\\")}
}

func (t *ClassLikeToken) Match(pattern string) bool {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	matches := r.FindStringSubmatch(t.ClassName)
	return len(matches) > 0
}

func (t *ClassLikeToken) ToString() string {
	return t.ClassName
}

func (t *ClassLikeToken) Equals(token *ClassLikeToken) bool {
	return t.ClassName == token.ClassName
}
