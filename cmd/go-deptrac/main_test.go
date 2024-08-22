package main

import (
	"flag"
	"github.com/KoNekoD/go-deptrac/pkg/src/supportive/Console/Application"
	"os"
	"testing"
)

func TestMainOk(t *testing.T) {

	os.Args = []string{
		"",
		"analyse",
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	//flag.CommandLine.Parse(os.Args)

	Application.NewApplication().Run()
}
