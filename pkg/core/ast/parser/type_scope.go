package parser

import (
	"go/ast"
	"strings"
)

type TypeScope struct {
	Namespace string
	Uses      map[string]string
	FileNode  *ast.File
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
		if strings.HasSuffix(key, classNameOrAlias) {
			return &value
		}
	}

	return nil
}
