package types

import (
	"go/ast"
	"strings"
)

type TypeScope struct {
	Namespace string
	Uses      map[string]string
	FileNode  *ast.File
	FilePath  string
}

func NewTypeScope(namespace string) *TypeScope {
	return &TypeScope{
		Namespace: namespace,
		Uses:      make(map[string]string),
		FileNode:  nil,
	}
}

func (s *TypeScope) SetFileNode(fileNode *ast.File) *TypeScope {
	s.FileNode = fileNode

	return s
}

func (s *TypeScope) SetFilePath(filepath string) *TypeScope {
	s.FilePath = filepath

	return s
}

func (s *TypeScope) AddUse(className string, alias *string) {
	key := className
	if alias != nil {
		key = *alias
	}

	// Trim " from key
	key = strings.Trim(key, "\"")
	className = strings.Trim(className, "\"")

	s.Uses[key] = className
}

func (s *TypeScope) GetUses() map[string]string {
	return s.Uses
}

func (s *TypeScope) GetUse(classNameOrAlias string) *string {
	if resolvedByAlias, ok := s.Uses[classNameOrAlias]; ok {
		return &resolvedByAlias
	}

	for key, value := range s.Uses {
		// split by / and check last element
		splittedKeys := strings.Split(key, "/")
		lastSplittedKey := splittedKeys[len(splittedKeys)-1]
		if lastSplittedKey == classNameOrAlias {
			return &value
		}
	}

	return nil
}
