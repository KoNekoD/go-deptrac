package nodes

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	_ "github.com/KoNekoD/go-deptrac/resources"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
	"strings"
)

type NodeNamer struct {
	projectDirectory string
}

func NewNodeNamer(projectDirectory string) *NodeNamer {
	return &NodeNamer{
		projectDirectory: projectDirectory,
	}
}

func (n *NodeNamer) GetRootPackageName(path string) (string, error) {
	fullPath := filepath.Join(n.projectDirectory, path)

	path, err := utils.GetPathWithoutFilename(fullPath)
	if err != nil {
		return "", err
	}

	// Recursively: for example path pkg/my/a/b/c we walk like this:
	// ( "pkg/my/a/b/c", "pkg/my/a/b", "pkg/my/a", "pkg/my", "pkg", "." )
	// And if it has go.mod we return the package name(module name)
	for {
		goModPath := filepath.Join(fullPath, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			fileBytes, err := os.ReadFile(goModPath)
			if err != nil {
				return "", err
			}

			parsedModfile, err := modfile.Parse(goModPath, fileBytes, nil)
			if err != nil {
				return "", err
			}
			return parsedModfile.Module.Mod.Path, nil
		}
		if filepath.Base(fullPath) == "." {
			return "", nil
		}
		fullPath = filepath.Dir(fullPath)
	}
}

func (n *NodeNamer) GetPackageName(path string) (string, error) {
	rootPackageName, err := n.GetRootPackageName(path)

	if err != nil {
		return "", err
	}

	innerPackageName, err := utils.GetPathWithoutFilename(filepath.Join(n.projectDirectory, path))
	if err != nil {
		return "", err
	}
	innerPackageName, err = filepath.Rel(n.projectDirectory, innerPackageName)

	if err != nil {
		return "", err
	}

	return filepath.Join(rootPackageName, innerPackageName), nil
}

func (n *NodeNamer) GetPackageFilename(path string) (string, error) {
	path, err := filepath.Rel(n.projectDirectory, path)
	if err != nil {
		return "", err
	}

	packageName, err := n.GetPackageName(path)

	if err != nil {
		return "", err
	}

	fileName, err := utils.GetPathWithOnlyFilename(filepath.Join(n.projectDirectory, path))
	if err != nil {
		return "", err
	}

	return filepath.Join(packageName, fileName), nil
}

func (n *NodeNamer) GetPackageStructName(path string, structName string) (string, error) {
	// Validate if structName has in path
	pathValidate := strings.Replace(path, "github.com/KoNekoD/go_deptrac/", "/home/mizuki/Documents/dev/KoNekoD/go_deptrac/", 1)
	nodes, err := parser.ParseFile(token.NewFileSet(), pathValidate, nil, 0)
	if err != nil {
		panic(err)
	}
	if nodes.Scope.Lookup(structName) == nil {
		panic("1")
	}

	packageFilename, err := n.GetPackageFilename(path)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", packageFilename, structName), err
}

func (n *NodeNamer) GetPackageFunctionName(path string, functionName string) (string, error) {
	packageFilename, err := n.GetPackageFilename(path)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", packageFilename, functionName), err
}

func (n *NodeNamer) GetPackageStructFunctionName(path string, structName string, functionName string) (string, error) {
	packageStructName, err := n.GetPackageStructName(path, structName)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s::%s", packageStructName, functionName), err
}
