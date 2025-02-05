package ast_maps

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/ast_inherits"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/tokens_references"
	"strings"
)

type AstMap struct {
	ClassReferences    map[string]*tokens_references.ClassLikeReference
	FileReferences     map[string]*tokens_references.FileReference
	FunctionReferences map[string]*tokens_references.FunctionReference
}

func NewAstMap(astFileReferences []*tokens_references.FileReference) *AstMap {
	a := &AstMap{
		ClassReferences:    make(map[string]*tokens_references.ClassLikeReference),
		FileReferences:     make(map[string]*tokens_references.FileReference),
		FunctionReferences: make(map[string]*tokens_references.FunctionReference),
	}
	for _, astFileReference := range astFileReferences {
		a.addAstFileReference(astFileReference)
	}
	return a
}

func (a *AstMap) GetClassLikeReferences() []*tokens_references.ClassLikeReference {
	values := make([]*tokens_references.ClassLikeReference, 0)

	for _, r := range a.ClassReferences {
		values = append(values, r)
	}

	return values
}

func (a *AstMap) GetFileReferences() []*tokens_references.FileReference {
	values := make([]*tokens_references.FileReference, 0)

	for _, fileReference := range a.FileReferences {
		values = append(values, fileReference)
	}

	return values
}

func (a *AstMap) GetFunctionReferences() []*tokens_references.FunctionReference {
	values := make([]*tokens_references.FunctionReference, 0)

	for _, functionReference := range a.FunctionReferences {
		values = append(values, functionReference)
	}

	return values
}

func (a *AstMap) GetClassReferenceForToken(structName *tokens.ClassLikeToken) *tokens_references.ClassLikeReference {
	// TODO: Rework to full package path
	name := structName.ToString()

	v, ok := a.ClassReferences[name]
	if !ok {

		// TODO: debug
		for refName, reference := range a.ClassReferences {
			refNameFile := strings.Split(refName, " ")[0]
			nameFile := strings.Split(name, " ")[0]

			if refNameFile == nameFile {
				fmt.Println(reference)
				//panic("1")
				//fmt.Println(reference)
				// todo почему то есть файлы которых он не находит... и они пападают в этот кейс
			}
		}

		// TODO: Possible external package
		return nil
	}

	return v
}

func (a *AstMap) GetFunctionReferenceForToken(functionName *tokens.FunctionToken) *tokens_references.FunctionReference {
	v, ok := a.FunctionReferences[functionName.ToString()]
	if !ok {
		return nil
	}

	return v
}

func (a *AstMap) GetFileReferenceForToken(filePath *tokens.FileToken) *tokens_references.FileReference {
	v, ok := a.FileReferences[filePath.ToString()]
	if !ok {
		return nil
	}

	return v
}

func (a *AstMap) GetClassInherits(structLikeName *tokens.ClassLikeToken) []*ast_inherits.AstInherit {
	structReference := a.GetClassReferenceForToken(structLikeName)
	if structReference == nil {
		return nil
	}
	inherits := make([]*ast_inherits.AstInherit, 0)
	for _, dep := range structReference.Inherits {
		inherits = append(inherits, dep)
		outArr := a.recursivelyResolveDependencies(dep, nil, nil)
		for _, inherit := range outArr {
			inherits = append(inherits, inherit)
		}
	}
	return inherits
}

type stack struct {
	s []*ast_inherits.AstInherit
}

func (s *stack) Push(v *ast_inherits.AstInherit) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() *ast_inherits.AstInherit {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}

func (a *AstMap) recursivelyResolveDependencies(inheritDependency *ast_inherits.AstInherit, alreadyResolved map[string]bool, pathStack *stack) []*ast_inherits.AstInherit {
	if alreadyResolved == nil {
		alreadyResolved = make(map[string]bool)
	}
	if pathStack == nil {
		pathStack = &stack{s: make([]*ast_inherits.AstInherit, 0)}
		pathStack.Push(inheritDependency)
	}
	structName := inheritDependency.ClassLikeName.ToString()
	if _, ok := alreadyResolved[structName]; ok {
		pathStack.Pop()
		return nil
	}
	structReference := a.GetClassReferenceForToken(inheritDependency.ClassLikeName)
	if structReference == nil {
		return nil
	}
	out := make([]*ast_inherits.AstInherit, 0)
	for _, inherit := range structReference.Inherits {
		alreadyResolved[structName] = true
		path := pathStack.s
		out = append(out, inherit.ReplacePath(path))
		pathStack.Push(inherit)
		outArr := a.recursivelyResolveDependencies(inherit, alreadyResolved, pathStack)
		for _, astInherit := range outArr {
			out = append(out, astInherit)
		}
		delete(alreadyResolved, structName)
		pathStack.Pop()
	}
	return out
}

func (a *AstMap) addClassLike(astStructReference tokens_references.ClassLikeReference) {
	token := astStructReference.GetToken()

	// If token.ToString() contains :: then panic
	if strings.Contains(token.ToString(), "::") {
		panic(token.ToString())
	}

	a.ClassReferences[token.ToString()] = &astStructReference
}

func (a *AstMap) addAstFileReference(astFileReference *tokens_references.FileReference) {
	a.FileReferences[*astFileReference.Filepath] = astFileReference
	for _, astStructReference := range astFileReference.ClassLikeReferences {
		a.addClassLike(*astStructReference)
	}

	for _, astFunctionReference := range astFileReference.FunctionReferences {
		a.addFunction(*astFunctionReference)
	}
}

func (a *AstMap) addFunction(astFunctionReference tokens_references.FunctionReference) {
	a.FunctionReferences[astFunctionReference.GetToken().ToString()] = &astFunctionReference
}
