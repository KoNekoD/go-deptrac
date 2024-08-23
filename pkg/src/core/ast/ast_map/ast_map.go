package ast_map

import "strings"

type AstMap struct {
	ClassReferences    map[string]*ClassLikeReference
	FileReferences     map[string]*FileReference
	FunctionReferences map[string]*FunctionReference
}

func NewAstMap(astFileReferences []*FileReference) *AstMap {
	a := &AstMap{
		ClassReferences:    make(map[string]*ClassLikeReference),
		FileReferences:     make(map[string]*FileReference),
		FunctionReferences: make(map[string]*FunctionReference),
	}
	for _, astFileReference := range astFileReferences {
		a.addAstFileReference(astFileReference)
	}
	return a
}

func (a *AstMap) GetClassLikeReferences() []*ClassLikeReference {
	values := make([]*ClassLikeReference, 0)

	for _, r := range a.ClassReferences {
		values = append(values, r)
	}

	return values
}

func (a *AstMap) GetFileReferences() []*FileReference {
	values := make([]*FileReference, 0)

	for _, fileReference := range a.FileReferences {
		values = append(values, fileReference)
	}

	return values
}

func (a *AstMap) GetFunctionReferences() []*FunctionReference {
	values := make([]*FunctionReference, 0)

	for _, functionReference := range a.FunctionReferences {
		values = append(values, functionReference)
	}

	return values
}

func (a *AstMap) GetClassReferenceForToken(structName *ClassLikeToken) *ClassLikeReference {
	// TODO: Rework to full package path
	name := structName.ToString()
	name = name[strings.LastIndex(name, "/")+1:]

	v, ok := a.ClassReferences[name]
	if !ok {
		// TODO: Possible external package
		return nil
	}

	return v
}

func (a *AstMap) GetFunctionReferenceForToken(functionName *FunctionToken) *FunctionReference {
	v, ok := a.FunctionReferences[functionName.ToString()]
	if !ok {
		return nil
	}

	return v
}

func (a *AstMap) GetFileReferenceForToken(filePath *FileToken) *FileReference {
	v, ok := a.FileReferences[filePath.ToString()]
	if !ok {
		return nil
	}

	return v
}

func (a *AstMap) GetClassInherits(structLikeName *ClassLikeToken) []*AstInherit {
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

func (a *AstMap) addClassLike(astStructReference ClassLikeReference) {
	token := astStructReference.GetToken()
	a.ClassReferences[token.ToString()] = &astStructReference
}

func (a *AstMap) addAstFileReference(astFileReference *FileReference) {
	a.FileReferences[*astFileReference.Filepath] = astFileReference
	for _, astStructReference := range astFileReference.ClassLikeReferences {
		a.addClassLike(*astStructReference)
	}

	for _, astFunctionReference := range astFileReference.FunctionReferences {
		a.addFunction(*astFunctionReference)
	}
}

func (a *AstMap) addFunction(astFunctionReference FunctionReference) {
	a.FunctionReferences[astFunctionReference.GetToken().ToString()] = &astFunctionReference
}
