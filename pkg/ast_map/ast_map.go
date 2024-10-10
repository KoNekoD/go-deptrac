package ast_map

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/tokens"
	"strings"
)

type AstMap struct {
	ClassReferences    map[string]*references.ClassLikeReference
	FileReferences     map[string]*references.FileReference
	FunctionReferences map[string]*references.FunctionReference
}

func NewAstMap(astFileReferences []*references.FileReference) *AstMap {
	a := &AstMap{
		ClassReferences:    make(map[string]*references.ClassLikeReference),
		FileReferences:     make(map[string]*references.FileReference),
		FunctionReferences: make(map[string]*references.FunctionReference),
	}
	for _, astFileReference := range astFileReferences {
		a.addAstFileReference(astFileReference)
	}
	return a
}

func (a *AstMap) GetClassLikeReferences() []*references.ClassLikeReference {
	values := make([]*references.ClassLikeReference, 0)

	for _, r := range a.ClassReferences {
		values = append(values, r)
	}

	return values
}

func (a *AstMap) GetFileReferences() []*references.FileReference {
	values := make([]*references.FileReference, 0)

	for _, fileReference := range a.FileReferences {
		values = append(values, fileReference)
	}

	return values
}

func (a *AstMap) GetFunctionReferences() []*references.FunctionReference {
	values := make([]*references.FunctionReference, 0)

	for _, functionReference := range a.FunctionReferences {
		values = append(values, functionReference)
	}

	return values
}

func (a *AstMap) GetClassReferenceForToken(structName *tokens.ClassLikeToken) *references.ClassLikeReference {
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

func (a *AstMap) GetFunctionReferenceForToken(functionName *tokens.FunctionToken) *references.FunctionReference {
	v, ok := a.FunctionReferences[functionName.ToString()]
	if !ok {
		return nil
	}

	return v
}

func (a *AstMap) GetFileReferenceForToken(filePath *tokens.FileToken) *references.FileReference {
	v, ok := a.FileReferences[filePath.ToString()]
	if !ok {
		return nil
	}

	return v
}

func (a *AstMap) GetClassInherits(structLikeName *tokens.ClassLikeToken) []*AstInherit {
	structReference := a.GetClassReferenceForToken(structLikeName)
	if structReference == nil {
		return nil
	}
	inherits := make([]*AstInherit, 0)
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
	s []*AstInherit
}

func (s *stack) Push(v *AstInherit) {
	s.s = append(s.s, v)
}

func (s *stack) Pop() *AstInherit {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}

func (a *AstMap) recursivelyResolveDependencies(inheritDependency *AstInherit, alreadyResolved map[string]bool, pathStack *stack) []*AstInherit {
	if alreadyResolved == nil {
		alreadyResolved = make(map[string]bool)
	}
	if pathStack == nil {
		pathStack = &stack{s: make([]*AstInherit, 0)}
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
	out := make([]*AstInherit, 0)
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

func (a *AstMap) addClassLike(astStructReference references.ClassLikeReference) {
	token := astStructReference.GetToken()

	// If token.ToString() contains :: then panic
	if strings.Contains(token.ToString(), "::") {
		panic(token.ToString())
	}

	// github.com/KoNekoD/go_deptrac/pkg/core/ast_contract/ast_map/emitter.go AstMap
	if strings.Contains(token.ToString(), "github.com/KoNekoD/go_deptrac/pkg/core/ast_contract/ast_map/emitter.go AstMap") {
		panic(token.ToString())
	}

	a.ClassReferences[token.ToString()] = &astStructReference
}

func (a *AstMap) addAstFileReference(astFileReference *references.FileReference) {
	a.FileReferences[*astFileReference.Filepath] = astFileReference
	for _, astStructReference := range astFileReference.ClassLikeReferences {
		a.addClassLike(*astStructReference)
	}

	for _, astFunctionReference := range astFileReference.FunctionReferences {
		a.addFunction(*astFunctionReference)
	}
}

func (a *AstMap) addFunction(astFunctionReference references.FunctionReference) {
	a.FunctionReferences[astFunctionReference.GetToken().ToString()] = &astFunctionReference
}
