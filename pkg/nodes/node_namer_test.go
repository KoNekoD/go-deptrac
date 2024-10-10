package nodes_test

import (
	"github.com/KoNekoD/go-deptrac/pkg/app"
	"github.com/KoNekoD/go-deptrac/pkg/nodes"
	"testing"
)

func TestOk(t *testing.T) {
	namer := app.ProvideTestContainerService("NodeNamer").(*nodes.NodeNamer)

	name, err := namer.GetRootPackageName("pkg/supportive/console_supportive/application")

	if err != nil {
		t.Error(err)
	}

	t.Log(name)
}
