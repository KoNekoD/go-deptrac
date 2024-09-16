package nikic_php_parser

import (
	ast_map2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/ast_map"
	parser2 "github.com/KoNekoD/go-deptrac/pkg/ast_core/parser"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/nikic_php_parser/node_namer"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"go/ast"
	"go/token"
)

type FileReferenceVisitor struct {
	dependencyResolvers  []extractors.ReferenceExtractorInterface
	currentTypeScope     *parser2.TypeScope
	currentReference     ast_map2.ReferenceBuilderInterface
	fileReferenceBuilder *ast_map2.FileReferenceBuilder
	typeResolver         *parser2.TypeResolver
	nodeNamer            *node_namer.NodeNamer
	errors               []error
	nestingStack         []ast.Node
}

func NewFileReferenceVisitor(fileReferenceBuilder *ast_map2.FileReferenceBuilder, resolver *parser2.TypeResolver, nodeNamer *node_namer.NodeNamer, extractors ...extractors.ReferenceExtractorInterface) *FileReferenceVisitor {
	return &FileReferenceVisitor{
		currentReference:     fileReferenceBuilder,
		fileReferenceBuilder: fileReferenceBuilder,
		typeResolver:         resolver,
		nodeNamer:            nodeNamer,
		errors:               make([]error, 0),
		dependencyResolvers:  extractors,
		currentTypeScope:     parser2.NewTypeScope(""),
	}
}

func (f *FileReferenceVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if len(f.nestingStack) > 0 {
		lastNesting := f.nestingStack[len(f.nestingStack)-1]
		if lastNesting == node {
			return f
		}
	}

	if node != nil {
		f.nestingStack = append(f.nestingStack, node)
		f.enterNode(node)

		defer func() {
			nodeForLeave := f.nestingStack[:len(f.nestingStack)-1]

			f.leaveNode(node)

			f.nestingStack = nodeForLeave
		}()
	}

	if node == nil {
		return f
	} else {
		ast.Walk(f, node)
	}

	return nil
}

func (f *FileReferenceVisitor) enterNode(node ast.Node) {
	switch typedNode := node.(type) {
	case *ast.File:
		packageFileName, err := f.nodeNamer.GetPackageFilename(f.fileReferenceBuilder.Filepath) // TODO: Possible bug when file_supportive dir  != package declared in file_supportive
		f.addErrIfNeeded(err)

		f.currentTypeScope = parser2.NewTypeScope(packageFileName).SetFileNode(typedNode)
	case *ast.FuncDecl:
		f.enterFunction(typedNode)
	case *ast.GenDecl:
		f.enterGenDecl(typedNode)
	}
}
func (f *FileReferenceVisitor) leaveNode(node ast.Node) {
	switch v := node.(type) {
	case *ast.FuncDecl:
		f.currentReference = f.fileReferenceBuilder
	case *ast.GenDecl:
		if v.Tok == token.TYPE {
			if f.getClassReferenceName(v) != nil {
				f.currentReference = f.fileReferenceBuilder
			}
		}
	case *ast.ImportSpec:
		f.leaveUse(v)
	}

	for _, resolver := range f.dependencyResolvers {
		resolver.ProcessNode(node, f.currentReference, f.currentTypeScope)
	}
}

func (f *FileReferenceVisitor) enterFunction(node *ast.FuncDecl) {
	var (
		fullName string
		err      error
	)

	if node.Recv != nil { // Function is a method
		if len(node.Recv.List) > 1 {
			panic("No way")
		}
		methodOwner := ""
		switch t := node.Recv.List[0].Type.(type) {
		case *ast.Ident:
			methodOwner = t.Name
		case *ast.StarExpr:
			methodOwner = t.X.(*ast.Ident).Name
		default:
			panic("No way")
		}
		methodName := node.Name.String()

		fullName, err = f.nodeNamer.GetPackageStructFunctionName(f.fileReferenceBuilder.Filepath, methodOwner, methodName)
		f.addErrIfNeeded(err)
	} else { // Function is a function
		methodFile := f.fileReferenceBuilder.Filepath
		name := node.Name.String()
		fullName, err = f.nodeNamer.GetPackageFunctionName(methodFile, name)
		f.addErrIfNeeded(err)
	}

	f.currentReference = f.fileReferenceBuilder.NewFunction(fullName, make([]string, 0), make(map[string][]string))

	for _, param := range node.Type.Params.List {
		if param.Type != nil {
			for _, classLikeName := range f.typeResolver.ResolvePHPParserTypes(f.currentTypeScope, param.Type) {
				f.currentReference.Parameter(classLikeName, util.GetLineByPosition(f.fileReferenceBuilder.Filepath, int(param.Type.Pos())))
			}
		}
	}

	if node.Type.Results != nil {
		for _, returnType := range node.Type.Results.List {
			if returnType.Type != nil {
				for _, classLikeName := range f.typeResolver.ResolvePHPParserTypes(f.currentTypeScope, returnType.Type) {
					f.currentReference.ReturnType(classLikeName, util.GetLineByPosition(f.fileReferenceBuilder.Filepath, int(returnType.Type.Pos())))
				}
			}
		}
	}
}

func (f *FileReferenceVisitor) addErrIfNeeded(errToAdd error) {
	if errToAdd != nil {
		f.errors = append(f.errors, errToAdd)
	}
}

func (f *FileReferenceVisitor) getClassReferenceName(node *ast.GenDecl) *string {
	if node.Tok == token.TYPE {
		structName := node.Specs[0].(*ast.TypeSpec).Name.Name
		name, err := f.nodeNamer.GetPackageStructName(f.fileReferenceBuilder.Filepath, structName)
		f.addErrIfNeeded(err)
		return &name
	}
	panic("1")
}

func (f *FileReferenceVisitor) enterGenDecl(node *ast.GenDecl) {
	if node.Tok != token.TYPE {
		return
	}

	for _, spec := range node.Specs {
		typeSpec := spec.(*ast.TypeSpec)

		switch typeSpec.Type.(type) {
		case *ast.StructType:
			structName := typeSpec.Name.Name
			packaeStructName, err := f.nodeNamer.GetPackageStructName(f.fileReferenceBuilder.Filepath, structName)
			f.addErrIfNeeded(err)
			f.enterClass(packaeStructName, make(map[string][]string)) // type T struct {}
		case *ast.Ident:
			structName := typeSpec.Name.Name
			packaeStructName, err := f.nodeNamer.GetPackageStructName(f.fileReferenceBuilder.Filepath, structName)
			f.addErrIfNeeded(err)
			f.enterClass(packaeStructName, make(map[string][]string)) // type T string
		case *ast.InterfaceType:
			structName := typeSpec.Name.Name
			packaeStructName, err := f.nodeNamer.GetPackageStructName(f.fileReferenceBuilder.Filepath, structName)
			f.addErrIfNeeded(err)
			f.enterInterface(packaeStructName, make(map[string][]string))
		default:
			panic("2")
		}
	}
}

func (f *FileReferenceVisitor) enterInterface(name string, tags map[string][]string) {
	f.currentReference = f.fileReferenceBuilder.NewInterface(name, make([]string, 0), tags)
}

func (f *FileReferenceVisitor) enterClass(name string, tags map[string][]string) {
	f.currentReference = f.fileReferenceBuilder.NewClass(name, make([]string, 0), tags)
}

func (f *FileReferenceVisitor) leaveUse(node *ast.ImportSpec) {
	classLikeName := node.Path.Value

	var alias *string
	if node.Name != nil {
		alias = &node.Name.Name
	}

	f.currentTypeScope.AddUse(classLikeName, alias)
	f.fileReferenceBuilder.UseStatement(classLikeName, util.GetLineByPosition(f.fileReferenceBuilder.Filepath, int(node.Path.Pos())))
}
