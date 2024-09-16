package simple_cleanarch

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/console_supportive/application"
	_ "github.com/KoNekoD/go-deptrac/resources"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	os.Args = []string{
		"",
		"--config_contract-file_supportive=pkg/test_projects/examples/simple-invalid-mvc/depfile.yaml",
		"analyse",
	}

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//err := flag.CommandLine.Parse(os.Args)
	//if err != nil {
	//	t.Fatal(err)
	//}

	application.NewApplication().
		Run()
}
