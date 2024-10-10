package simple_cleanarch

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/app"
	_ "github.com/KoNekoD/go-deptrac/resources"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	os.Args = []string{
		"",
		"--config=pkg/test_projects/examples/simple_mvc/deptrac.yaml",
		"analyse",
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//err := flag.CommandLine.Parse(os.Args)
	//if err != nil {
	//	t.Fatal(err)
	//}

	app.NewApplication().
		Run()
}
