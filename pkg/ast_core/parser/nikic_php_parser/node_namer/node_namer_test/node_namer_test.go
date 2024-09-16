package node_namer_test

import (
	"github.com/KoNekoD/go-deptrac/pkg/ast_core/parser/nikic_php_parser/node_namer"
	"github.com/KoNekoD/go-deptrac/pkg/dependency_injection_supportive/test_container_provider"
	"testing"
)

func TestOk(t *testing.T) {
	namer := test_container_provider.ProvideTestContainerService("NodeNamer").(*node_namer.NodeNamer)

	name, err := namer.GetRootPackageName("pkg/supportive/console_supportive/application")

	if err != nil {
		t.Error(err)
	}

	t.Log(name)
}
