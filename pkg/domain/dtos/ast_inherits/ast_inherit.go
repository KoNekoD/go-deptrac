package ast_inherits

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/enums"
	"strings"
)

type AstInherit struct {
	ClassLikeName  *tokens.ClassLikeToken
	FileOccurrence *dtos.FileOccurrence
	Type           enums.AstInheritType
	path           []*AstInherit
}

func NewAstInherit(classLikeName *tokens.ClassLikeToken, fileOccurrence *dtos.FileOccurrence, astInheritType enums.AstInheritType, path []*AstInherit) *AstInherit {
	return &AstInherit{
		ClassLikeName:  classLikeName,
		FileOccurrence: fileOccurrence,
		Type:           astInheritType,
		path:           path,
	}
}

func (i *AstInherit) GetPath() []*AstInherit {
	return i.path
}

func ArrayReverse(s []*AstInherit) []*AstInherit {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (i *AstInherit) ToString() string {
	description := fmt.Sprintf("%s::%d (%s)", i.ClassLikeName.ToString(), i.FileOccurrence.Line, i.Type)
	if len(i.path) == 0 {
		return description
	}

	path := make([]string, 0)
	reverse := ArrayReverse(i.path)
	for _, p := range reverse {
		path = append(path, p.ToString())
	}
	return fmt.Sprintf("%s (path: %s)", description, strings.Join(path, " -> "))
}

func (i *AstInherit) ReplacePath(path []*AstInherit) *AstInherit {
	return &AstInherit{
		ClassLikeName:  i.ClassLikeName,
		FileOccurrence: i.FileOccurrence,
		Type:           i.Type,
		path:           path,
	}
}
