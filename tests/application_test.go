package tests

import (
	"flag"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/app"
	"os"
	"testing"
)

func TestApplicationOk(t *testing.T) {
	os.Args = []string{
		"",
		"analyse",
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	app.NewApp().Run()

	fmt.Println("", "")
}
