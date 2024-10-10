package examples

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/app"
	_ "github.com/KoNekoD/go-deptrac/resources"
	"os"
	"testing"
)

func TestRunSimpleCleanarch(t *testing.T) {
	t.Parallel()

	t.Run("simple_clean_arch", func(t *testing.T) {
		configArg := "--config=examples/simple_clean_arch/deptrac.yaml"
		os.Args = []string{"", configArg, "analyse"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		app.NewApplication().Run()
	})

	t.Run("simple_invalid_mvc", func(t *testing.T) {
		configArg := "--config=examples/simple_invalid_mvc/deptrac.yaml"
		os.Args = []string{"", configArg, "analyse"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		app.NewApplication().Run()
	})

	t.Run("simple_mvc", func(t *testing.T) {
		configArg := "--config=examples/simple_mvc/deptrac.yaml"
		os.Args = []string{"", configArg, "analyse"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		app.NewApplication().Run()
	})
}
