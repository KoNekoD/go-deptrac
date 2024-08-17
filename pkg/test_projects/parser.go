package test_projects

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyContext"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/DependencyType"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/FileOccurrence"
	"github.com/KoNekoD/go-deptrac/pkg/src/contract/Ast/TokenInterface"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/AstMap"
	"github.com/KoNekoD/go-deptrac/pkg/src/core/Ast/Parser/ParserInterface"
	"github.com/KoNekoD/go-deptrac/pkg/test_projects/abs3"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

type ABS interface {
	IsMyType(typeName string) bool
}

type fint int

type fint2 ABS

var (
	a1     ABS
	a2     ABS2
	a3     abs3.ABS3
	a4, a5 ABS
	a6     fint = 1
	a7          = struct{ a int }{a: 1}
	a8     int
	a9     *int
	a10    any
)

const CHHK = false

const (
	CHK1 fint = iota
	CHK2
	CHK3
)

type FileParser struct{}

func NewFileParser() ParserInterface.ParserInterface { return &FileParser{} }

type WrappedAstFileNode struct {
	*ast.File
	Filepath            string
	HandledObjectsNames map[string]bool
	steps               int
	lvl                 int
	imports             []wrappedAstFileNodeImport
}

type wrappedAstFileNodeImport struct {
	alias *string
	value string
}

func (w *WrappedAstFileNode) IsMyType(typeName string) bool {
	return w.Scope.Lookup(typeName) != nil
}

func (w *WrappedAstFileNode) badBetween(from, to token.Pos) {
	fmt.Println("Syntax error found", "file:", w.Filepath, "from(line):", from, "to(line):", to, "from(row):", util.GetLineByPosition(w.Filepath, int(from)), "to(row):", util.GetLineByPosition(w.Filepath, int(to)))
}

func (w *WrappedAstFileNode) handleScope(scope *ast.Scope) []*AstMap.DependencyToken {
	if scope == nil {
		return nil
	}

	scopeDeps := make([]*AstMap.DependencyToken, 0)

	if scope.Outer != nil {
		scopeDeps = append(scopeDeps, w.handleScope(scope.Outer)...)
	}

	for _, object := range scope.Objects {
		scopeDeps = append(scopeDeps, w.handleObject(object)...)
	}

	return scopeDeps
}

func (w *WrappedAstFileNode) handleObject(object *ast.Object) []*AstMap.DependencyToken {
	if object == nil {
		return nil
	}

	//if object.Kind == ast.Typ {
	if !w.IsMyType(object.Name) {
		if strings.HasSuffix(w.Filepath, "analysis_result.go") {
			fmt.Println("")
		}
	}
	//}

	// Prevent recursions in type definitions
	if _, ok := w.HandledObjectsNames[object.Name]; ok {
		return nil
	}
	w.HandledObjectsNames[object.Name] = true

	return w.handleNode(object.Decl.(ast.Node))
}

func (w *WrappedAstFileNode) handleFile(n *ast.File) []*AstMap.DependencyToken {
	docDeps := w.handleNode(n.Doc)
	nameDeps := w.handleNode(n.Name)
	declsDeps := make([]*AstMap.DependencyToken, 0)
	scopeDeps := make([]*AstMap.DependencyToken, 0)
	importsDeps := make([]*AstMap.DependencyToken, 0)
	unresolvedDeps := make([]*AstMap.DependencyToken, 0)
	commentsDeps := make([]*AstMap.DependencyToken, 0)

	for _, decl := range n.Decls {
		declsDeps = append(declsDeps, w.handleNode(decl)...)
	}

	scopeDeps = append(scopeDeps, w.handleScope(n.Scope)...)

	for _, imp := range n.Imports {
		importsDeps = append(importsDeps, w.handleNode(imp)...)
	}

	for _, unresolved := range n.Unresolved {
		unresolvedDeps = append(unresolvedDeps, w.handleNode(unresolved)...)
	}

	for _, comment := range n.Comments {
		commentsDeps = append(commentsDeps, w.handleNode(comment)...)
	}

	return append(docDeps, append(nameDeps, append(declsDeps, append(scopeDeps, append(importsDeps, append(unresolvedDeps, commentsDeps...)...)...)...)...)...)
}

func (w *WrappedAstFileNode) SelfHandle() []*AstMap.DependencyToken {
	totalDeps := w.handleNode(w.File)

	theirsDeps := make([]*AstMap.DependencyToken, 0)

	for _, dep := range totalDeps {
		if !w.IsMyType(dep.Token.ToString()) {
			theirsDeps = append(theirsDeps, dep)
		}
	}

	return theirsDeps
}

func (w *WrappedAstFileNode) handleNode(node ast.Node) []*AstMap.DependencyToken {
	w.steps++
	w.lvl++
	defer func() {
		w.lvl--
	}()

	tabs := ""
	for i := 0; i < w.lvl; i++ {
		tabs += "\t"
	}

	fmt.Printf("%sNode type: %T steps: %d lvl: %d\n", tabs, node, w.steps, w.lvl)

	if node == nil {
		return nil
	}

	val := reflect.ValueOf(node)
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return nil
	}

	switch n := node.(type) {
	case *ast.ArrayType:
		lenDeps := w.handleNode(n.Len)
		eltDeps := w.handleNode(n.Elt)

		return append(lenDeps, eltDeps...)
	case *ast.AssignStmt:
		lhsDeps := make([]*AstMap.DependencyToken, 0)
		rhsDeps := make([]*AstMap.DependencyToken, 0)

		for _, lh := range n.Lhs {
			lhsDeps = append(lhsDeps, w.handleNode(lh)...)
		}
		for _, rh := range n.Rhs {
			rhsDeps = append(rhsDeps, w.handleNode(rh)...)
		}

		return append(lhsDeps, rhsDeps...)
	case *ast.BadDecl:
		w.badBetween(n.From, n.To)

		return nil
	case *ast.BadExpr:
		w.badBetween(n.From, n.To)

		return nil
	case *ast.BadStmt:
		w.badBetween(n.From, n.To)

		return nil
	case *ast.BasicLit:
		return nil
	case *ast.BinaryExpr:
		XDeps := w.handleNode(n.X)
		YDeps := w.handleNode(n.Y)

		return append(XDeps, YDeps...)
	case *ast.BlockStmt:
		listDeps := make([]*AstMap.DependencyToken, 0)

		for _, stmt := range n.List {
			listDeps = append(listDeps, w.handleNode(stmt)...)
		}

		return listDeps

	case *ast.BranchStmt:
		labelDeps := w.handleNode(n.Label)

		return labelDeps

	case *ast.CallExpr:
		return w.handleCallExpr(n)
	case *ast.CaseClause:
		listDeps := make([]*AstMap.DependencyToken, 0)
		bodyDeps := make([]*AstMap.DependencyToken, 0)

		for _, stmt := range n.List {
			listDeps = append(listDeps, w.handleNode(stmt)...)
		}

		for _, stmt := range n.Body {
			bodyDeps = append(bodyDeps, w.handleNode(stmt)...)
		}

		return append(listDeps, bodyDeps...)
	case *ast.ChanType:
		valueDeps := w.handleNode(n.Value)

		return valueDeps
	case *ast.CommClause:
		commDeps := w.handleNode(n.Comm)
		bodyDeps := make([]*AstMap.DependencyToken, 0)

		for _, stmt := range n.Body {
			bodyDeps = append(bodyDeps, w.handleNode(stmt)...)
		}

		return append(commDeps, bodyDeps...)
	case *ast.Comment:
		return nil
	case *ast.CommentGroup:
		listDeps := make([]*AstMap.DependencyToken, 0)

		if n == nil {
			panic(1)
		}

		for _, comment := range n.List {
			listDeps = append(listDeps, w.handleNode(comment)...)
		}

		return listDeps
	case *ast.CompositeLit:
		typeDeps := w.handleNode(n.Type)
		eltsDeps := make([]*AstMap.DependencyToken, 0)

		for _, elt := range n.Elts {
			eltsDeps = append(eltsDeps, w.handleNode(elt)...)
		}

		return append(typeDeps, eltsDeps...)
	case *ast.DeclStmt:
		return w.handleNode(n.Decl)
	case *ast.DeferStmt:
		return w.handleNode(n.Call)
	case *ast.Ellipsis:
		return w.handleNode(n.Elt)
	case *ast.EmptyStmt:
		return nil
	case *ast.ExprStmt:
		return w.handleNode(n.X)
	case *ast.Field:
		docDeps := w.handleNode(n.Doc)
		namesDeps := make([]*AstMap.DependencyToken, 0)
		typeDeps := w.handleNode(n.Type)
		tagDeps := w.handleNode(n.Tag)
		commentDeps := w.handleNode(n.Comment)

		for _, name := range n.Names {
			namesDeps = append(namesDeps, w.handleNode(name)...)
		}

		return append(docDeps, append(namesDeps, append(typeDeps, append(tagDeps, commentDeps...)...)...)...)
	case *ast.FieldList:
		listDeps := make([]*AstMap.DependencyToken, 0)

		for _, field := range n.List {
			listDeps = append(listDeps, w.handleNode(field)...)
		}

		return listDeps
	case *ast.File:
		return w.handleFile(n)
	case *ast.ForStmt:
		initDeps := w.handleNode(n.Init)
		condDeps := w.handleNode(n.Cond)
		postDeps := w.handleNode(n.Post)
		bodyDeps := w.handleNode(n.Body)

		return append(initDeps, append(condDeps, append(postDeps, bodyDeps...)...)...)
	case *ast.FuncDecl:
		docDeps := w.handleNode(n.Doc)
		recvDeps := w.handleNode(n.Recv)
		nameDeps := w.handleNode(n.Name)
		typeDeps := w.handleNode(n.Type)
		bodyDeps := w.handleNode(n.Body)

		return append(docDeps, append(recvDeps, append(nameDeps, append(typeDeps, bodyDeps...)...)...)...)
	case *ast.FuncLit:
		typeDeps := w.handleNode(n.Type)
		bodyDeps := w.handleNode(n.Body)

		return append(typeDeps, bodyDeps...)
	case *ast.FuncType:
		typeParamsDeps := w.handleNode(n.TypeParams)
		paramsDeps := w.handleNode(n.Params)
		resultsDeps := w.handleNode(n.Results)

		return append(typeParamsDeps, append(paramsDeps, resultsDeps...)...)
	case *ast.GenDecl:
		docDeps := w.handleNode(n.Doc)
		specsDeps := make([]*AstMap.DependencyToken, 0)

		for _, spec := range n.Specs {
			specsDeps = append(specsDeps, w.handleNode(spec)...)
		}

		return append(docDeps, specsDeps...)
	case *ast.GoStmt:
		return w.handleNode(n.Call)
	case *ast.Ident:
		return w.handleObject(n.Obj)
	case *ast.IfStmt:
		initDeps := w.handleNode(n.Init)
		condDeps := w.handleNode(n.Cond)
		bodyDeps := w.handleNode(n.Body)
		elseDeps := w.handleNode(n.Else)

		return append(initDeps, append(condDeps, append(bodyDeps, elseDeps...)...)...)
	case *ast.ImportSpec:
		return w.handleImportSpec(n)
	case *ast.IncDecStmt:
		return w.handleNode(n.X)
	case *ast.IndexExpr:
		xDeps := w.handleNode(n.X)
		indexDeps := w.handleNode(n.Index)

		return append(xDeps, indexDeps...)
	case *ast.IndexListExpr:
		xDeps := w.handleNode(n.X)
		indexDeps := make([]*AstMap.DependencyToken, 0)

		for _, index := range n.Indices {
			indexDeps = append(indexDeps, w.handleNode(index)...)
		}

		return append(xDeps, indexDeps...)
	case *ast.InterfaceType:
		return w.handleNode(n.Methods)
	case *ast.KeyValueExpr:
		keyDeps := w.handleNode(n.Key)
		valueDeps := w.handleNode(n.Value)

		return append(keyDeps, valueDeps...)
	case *ast.LabeledStmt:
		labelDeps := w.handleNode(n.Label)
		stmtDeps := w.handleNode(n.Stmt)

		return append(labelDeps, stmtDeps...)
	case *ast.MapType:
		keyDeps := w.handleNode(n.Key)
		valueDeps := w.handleNode(n.Value)

		return append(keyDeps, valueDeps...)
	case *ast.Package:
		scopeDeps := w.handleScope(n.Scope)
		importsDeps := make([]*AstMap.DependencyToken, 0)
		filesDeps := make([]*AstMap.DependencyToken, 0)

		for _, importSpec := range n.Imports {
			importsDeps = append(importsDeps, w.handleObject(importSpec)...)
		}

		for _, file := range n.Files {
			filesDeps = append(filesDeps, w.handleNode(file)...)
		}

		return append(scopeDeps, append(importsDeps, filesDeps...)...)
	case *ast.ParenExpr:
		return w.handleNode(n.X)
	case *ast.RangeStmt:
		keyDeps := w.handleNode(n.Key)
		valueDeps := w.handleNode(n.Value)
		xDeps := w.handleNode(n.X)
		bodyDeps := w.handleNode(n.Body)

		return append(keyDeps, append(valueDeps, append(xDeps, bodyDeps...)...)...)
	case *ast.ReturnStmt:
		resultsDeps := make([]*AstMap.DependencyToken, 0)

		for _, result := range n.Results {
			resultsDeps = append(resultsDeps, w.handleNode(result)...)
		}

		return resultsDeps

	case *ast.SelectStmt:
		return w.handleNode(n.Body)
	case *ast.SelectorExpr:
		return w.handleSelectorExpr(n, DependencyType.DependencyTypeAnonymousClassExtends) // TODO: DependencyTypeAnonymousClassExtends is wrong
	case *ast.SendStmt:
		chanDeps := w.handleNode(n.Chan)
		valueDeps := w.handleNode(n.Value)

		return append(chanDeps, valueDeps...)
	case *ast.SliceExpr:
		xDeps := w.handleNode(n.X)
		lowDeps := w.handleNode(n.Low)
		highDeps := w.handleNode(n.High)
		maxDeps := w.handleNode(n.Max)

		return append(xDeps, append(lowDeps, append(highDeps, maxDeps...)...)...)
	case *ast.StarExpr:
		return w.handleNode(n.X)
	case *ast.StructType:
		return w.handleNode(n.Fields)
	case *ast.SwitchStmt:
		initDeps := w.handleNode(n.Init)
		tagDeps := w.handleNode(n.Tag)
		bodyDeps := w.handleNode(n.Body)

		return append(initDeps, append(tagDeps, bodyDeps...)...)
	case *ast.TypeAssertExpr:
		xDeps := w.handleNode(n.X)
		typeDeps := w.handleNode(n.Type)

		return append(xDeps, typeDeps...)
	case *ast.TypeSpec:
		docDeps := w.handleNode(n.Doc)
		nameDeps := w.handleNode(n.Name)
		typeParamsDeps := w.handleNode(n.TypeParams)
		typeDeps := w.handleNode(n.Type)
		commentDeps := w.handleNode(n.Comment)

		return append(docDeps, append(nameDeps, append(typeParamsDeps, append(typeDeps, commentDeps...)...)...)...)
	case *ast.TypeSwitchStmt:
		initDeps := w.handleNode(n.Init)
		assignDeps := w.handleNode(n.Assign)
		bodyDeps := w.handleNode(n.Body)

		return append(initDeps, append(assignDeps, bodyDeps...)...)
	case *ast.UnaryExpr:
		return w.handleNode(n.X)
	case *ast.ValueSpec:
		docDeps := w.handleNode(n.Doc)
		namesDeps := make([]*AstMap.DependencyToken, 0)
		typeDeps := w.handleNode(n.Type)
		valuesDeps := make([]*AstMap.DependencyToken, 0)
		commentDeps := w.handleNode(n.Comment)

		for _, name := range n.Names {
			namesDeps = append(namesDeps, w.handleNode(name)...)
		}

		for _, value := range n.Values {
			valuesDeps = append(valuesDeps, w.handleNode(value)...)
		}

		return append(docDeps, append(namesDeps, append(typeDeps, append(valuesDeps, commentDeps...)...)...)...)

	default:
		panic(fmt.Sprintf("Unhandled type: %T", n))
	}
}

func (w *WrappedAstFileNode) handleImportSpec(n *ast.ImportSpec) []*AstMap.DependencyToken {
	var alias *string
	if n.Name != nil {
		alias = &n.Name.Name
	}
	w.imports = append(w.imports, wrappedAstFileNodeImport{alias: alias, value: n.Path.Value})

	docDeps := w.handleNode(n.Doc)
	nameDeps := w.handleNode(n.Name)
	pathDeps := w.handleNode(n.Path)
	commentDeps := w.handleNode(n.Comment)

	return append(docDeps, append(nameDeps, append(pathDeps, commentDeps...)...)...)
}

func (w *WrappedAstFileNode) handleSelectorExpr(n *ast.SelectorExpr, dependencyType DependencyType.DependencyType) []*AstMap.DependencyToken {

	// When this node means 'result.Type' for example
	deps := make([]*AstMap.DependencyToken, 0)
	if xIdent, ok := n.X.(*ast.Ident); ok {
		var depToken TokenInterface.TokenInterface

		if dependencyType == DependencyType.DependencyTypeAnonymousClassExtends {
			if n.Sel.Name == "PathNormalize" {
				fmt.Println("")
			}

			depToken = &AstMap.ClassLikeToken{ClassName: n.Sel.Name}
		} else if dependencyType == DependencyType.DependencyTypeUnresolvedFunctionCall {
			depToken = &AstMap.FunctionToken{FunctionName: n.Sel.Name}
		}

		deps = append(deps, AstMap.NewDependencyToken(depToken, DependencyContext.NewDependencyContext(FileOccurrence.NewFileOccurrence(w.Filepath, int(xIdent.Pos())), dependencyType)))
	}

	xDeps := w.handleNode(n.X)
	selDeps := w.handleNode(n.Sel)

	return append(deps, append(xDeps, selDeps...)...)
}

func (w *WrappedAstFileNode) handleCallExpr(n *ast.CallExpr) []*AstMap.DependencyToken {

	// When this node means 'result.Function()' for example
	deps := make([]*AstMap.DependencyToken, 0)
	if selectorExpr, ok := n.Fun.(*ast.SelectorExpr); ok {
		deps = append(deps, w.handleSelectorExpr(selectorExpr, DependencyType.DependencyTypeUnresolvedFunctionCall)...)
	} else {
		deps = w.handleNode(n.Fun)
	}

	argsDeps := make([]*AstMap.DependencyToken, 0)

	for _, arg := range n.Args {
		argsDeps = append(argsDeps, w.handleNode(arg)...)
	}

	return append(deps, argsDeps...)
}

func (p *FileParser) ParseFile(file string) (*AstMap.FileReference, error) {
	//fileReferenceBuilder := NewFileReferenceBuilder(file, f) TODO: remove function
	node := WrappedAstFileNode{File: p.loadFileNode(file), Filepath: file, HandledObjectsNames: make(map[string]bool)}

	deps := node.SelfHandle()

	typeLikeReferences := make([]*AstMap.ClassLikeReference, 0)
	functionReferences := make([]*AstMap.FunctionReference, 0)

	f := AstMap.NewFileReference(&file, typeLikeReferences, functionReferences, deps)

	for _, dep := range deps {
		if dep.Context.DependencyType == DependencyType.DependencyTypeAnonymousClassExtends {
			//t := dep.Token.(*ClassLikeToken.ClassLikeToken)

			//inherits := make([]*AstInherit.AstInherit, 0)
			//dependencies := make([]*DependencyToken.DependencyToken, 0)
			//tags := make(map[string][]string)

			//typeLikeReferences = append(typeLikeReferences, ClassLikeReference.NewClassLikeReference(t, inherits, dependencies, f, tags, f)) // TODO
		} else if dep.Context.DependencyType == DependencyType.DependencyTypeUnresolvedFunctionCall {
			//t := dep.Token.(*FunctionToken.FunctionToken)

			//dependencies := make([]*DependencyToken.DependencyToken, 0)
			//tags := make(map[string][]string)

			//functionReferences = append(functionReferences, FunctionReference.NewFunctionReference(t, dependencies, f, tags))
		}
	}

	f.ClassLikeReferences = typeLikeReferences
	f.FunctionReferences = functionReferences

	return f, nil
}

func (p *FileParser) loadFileNode(file string) *ast.File {
	nodes, err := parser.ParseFile(token.NewFileSet(), file, nil, 0)
	if err != nil {
		panic(err)
	}
	return nodes
}
