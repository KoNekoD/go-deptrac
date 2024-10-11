package ast_map

import (
	"encoding/json"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/di"
	"os"
	"path"
	"runtime"
	"testing"
)

func TestAstLoaderCreateAstMap(t *testing.T) {
	di.UseVoidConfig()
	loader := di.ProvideTestContainerService("AstLoader").(*AstLoader)

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../resources")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	wd, _ := os.Getwd()

	files := []string{
		wd + "/test/create_ast_map/main.go",
	}

	astMap, err := loader.CreateAstMap(files)
	if err != nil {
		t.Fatal(err)
	}

	bytes, _ := json.Marshal(astMap)
	content := string(bytes)

	snap := `{"ClassReferences":{"github.com/KoNekoD/go_deptrac/test/create_ast_map/main.go StructChild":{"Type":"class","Inherits":[],"Dependencies":[],"Tags":{}},"github.com/KoNekoD/go_deptrac/test/create_ast_map/main.go StructRoot":{"Type":"class","Inherits":[],"Dependencies":[],"Tags":{}}},"FileReferences":{"/home/mizuki/Documents/dev/KoNekoD/go_deptrac/resources/test/create_ast_map/main.go":{"Filepath":"/home/mizuki/Documents/dev/KoNekoD/go_deptrac/resources/test/create_ast_map/main.go","ClassLikeReferences":[{"Type":"class","Inherits":[],"Dependencies":[],"Tags":{}},{"Type":"class","Inherits":[],"Dependencies":[],"Tags":{}}],"FunctionReferences":[{"Tags":{},"Dependencies":[]}],"Dependencies":[]}},"FunctionReferences":{"github.com/KoNekoD/go_deptrac/test/create_ast_map/main.go StructRoot::rootMethod()":{"Tags":{},"Dependencies":[]}}}`

	if snap != content {
		t.Fatal("Snapshots mismatch")
	}
}
