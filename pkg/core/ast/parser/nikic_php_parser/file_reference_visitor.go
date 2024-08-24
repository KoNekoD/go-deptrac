package nikic_php_parser

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/ast_map"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/parser"
	"github.com/KoNekoD/go-deptrac/pkg/core/ast/parser/extractors"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"go/ast"
	"go/token"
)

type FileReferenceVisitor struct {
	dependencyResolvers  []extractors.ReferenceExtractorInterface
	currentTypeScope     *parser.TypeScope
	currentReference     ast_map.ReferenceBuilderInterface
	fileReferenceBuilder *ast_map.FileReferenceBuilder
	typeResolver         *parser.TypeResolver
	nestingStack         []ast.Node
}

func NewFileReferenceVisitor(fileReferenceBuilder *ast_map.FileReferenceBuilder, resolver *parser.TypeResolver, extractors ...extractors.ReferenceExtractorInterface) *FileReferenceVisitor {
	return &FileReferenceVisitor{
		currentReference:     fileReferenceBuilder,
		fileReferenceBuilder: fileReferenceBuilder,
		typeResolver:         resolver,
		dependencyResolvers:  extractors,
		currentTypeScope:     parser.NewTypeScope(""),
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
		f.currentTypeScope = parser.NewTypeScope(typedNode.Name.Name).SetFileNode(typedNode)
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
	var fullName *string
	if node.Recv != nil {
		if len(node.Recv.List) > 1 {
			panic("No way")
		}
		methodFile := f.fileReferenceBuilder.Filepath
		methodOwner := ""
		switch t := node.Recv.List[0].Type.(type) {
		case *ast.Ident:
			methodOwner = t.Name
		case *ast.StarExpr:
			methodOwner = t.X.(*ast.Ident).Name
		default:
			panic("No way")
		}
		name := node.Name.String()

		fn := fmt.Sprintf("%s %s::%s", methodFile, methodOwner, name)
		fullName = &fn
	} else {
		methodFile := f.fileReferenceBuilder.Filepath
		name := node.Name.String()
		fn := fmt.Sprintf("%s %s", methodFile, name)
		fullName = &fn
	}
	f.currentReference = f.fileReferenceBuilder.NewFunction(*fullName, make([]string, 0), make(map[string][]string))

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

func (f *FileReferenceVisitor) getClassReferenceName(node *ast.GenDecl) *string {
	if node.Tok == token.TYPE {
		name := node.Specs[0].(*ast.TypeSpec).Name.Name
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
			f.enterClass(typeSpec.Name.Name, make(map[string][]string)) // type T struct {}
		case *ast.Ident:
			f.enterClass(typeSpec.Name.Name, make(map[string][]string)) // type T string
		case *ast.InterfaceType:
			f.enterInterface(typeSpec.Name.Name, make(map[string][]string))
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
