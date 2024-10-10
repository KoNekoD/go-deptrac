package services_test

import (
	"github.com/KoNekoD/go-deptrac/pkg/app"
	"github.com/KoNekoD/go-deptrac/pkg/domain/services"
	"testing"
)

func TestOk(t *testing.T) {
	namer := app.ProvideTestContainerService("NodeNamer").(*services.NodeNamer)

	name, err := namer.GetRootPackageName("pkg/supportive/console_supportive/application")

	if err != nil {
		t.Error(err)
	}

	t.Log(name)
}
