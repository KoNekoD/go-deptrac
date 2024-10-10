package parsers

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/nodes"
	"github.com/KoNekoD/go-deptrac/pkg/references"
	"github.com/KoNekoD/go-deptrac/pkg/types"
	"github.com/pkg/errors"
	"go/ast"
	"go/parser"
	"go/token"
)

var classAstMap = make(map[string]*ast.Ident)

var parsedFiles = &parsedFilesBag{parsedFiles: make(map[string]*parsedFile)}

type parsedFilesBag struct {
	parsedFiles map[string]*parsedFile
}

func (p *parsedFilesBag) Add(fileReference *references.FileReference, rootNode *ast.File) {
	filepath := fileReference.GetFilepath()
	file := &parsedFile{fileReference: fileReference, rootNode: rootNode}
	p.parsedFiles[*filepath] = file
}

func (p *parsedFilesBag) Get(filepath string) *parsedFile {
	return p.parsedFiles[filepath]
}

type parsedFile struct {
	fileReference *references.FileReference
	rootNode      *ast.File
	debt          []interface{}
}

func (p *parsedFile) GetOwnType(name string) *ast.Ident {
	for _, object := range p.rootNode.Scope.Objects {
		objectType := object.Decl.(*ast.TypeSpec)

		return objectType.Name
	}

	panic("Type not found")
}

type NikicPhpParser struct {
	classAstMap  map[string]*ast.Ident
	cache        ast_map.AstFileReferenceCacheInterface
	typeResolver *types.TypeResolver
	nodeNamer    *nodes.NodeNamer
	extractors   []references.ReferenceExtractorInterface
}

func NewNikicPhpParser(cache ast_map.AstFileReferenceCacheInterface, typeResolver *types.TypeResolver, nodeNamer *nodes.NodeNamer, extractors []references.ReferenceExtractorInterface) *NikicPhpParser {
	return &NikicPhpParser{
		classAstMap:  make(map[string]*ast.Ident),
		cache:        cache,
		typeResolver: typeResolver,
		nodeNamer:    nodeNamer,
		extractors:   extractors,
	}
}

func (p *NikicPhpParser) ParseFile(file string) (*references.FileReference, error) {
	v, err := p.cache.Get(file)
	if err != nil {
		return nil, err
	}
	if v != nil {
		return v, nil
	}

	fileReferenceBuilder := references.CreateFileReferenceBuilder(file)
	visitor := references.NewFileReferenceVisitor(fileReferenceBuilder, p.typeResolver, p.nodeNamer, p.extractors...)
	rootNode := p.loadNodesFromFile(file)

	ast.Walk(visitor, rootNode)

	if err = visitor.GetError(); err != nil {
		return nil, errors.WithStack(err)
	}

	fileReference := fileReferenceBuilder.Build()

	errCacheSet := p.cache.Set(fileReference)
	if errCacheSet != nil {
		return nil, errCacheSet
	}
	parsedFiles.Add(fileReference, rootNode)

	return fileReference, nil
}

func (p *NikicPhpParser) GetNodeForClassLikeReference(classReference *references.ClassLikeReference) *ast.Ident {
	classLikeName := classReference.GetToken().ToString()
	if v, ok := classAstMap[classLikeName]; ok {
		return v
	}

	filepath := classReference.GetFilepath()
	if nil == filepath {
		return nil
	}

	parsedFileRef := parsedFiles.Get(*filepath)

	return parsedFileRef.GetOwnType(classLikeName)
}

func (p *NikicPhpParser) loadNodesFromFile(file string) *ast.File {
	nodes, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
	if err != nil {
		panic(err)
	}
	return nodes
}
