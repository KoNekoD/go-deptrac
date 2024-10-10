package app

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestApplicationOk(t *testing.T) {
	os.Args = []string{
		"",
		"analyse",
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	NewApplication().Run()

	fmt.Println("", "")
}
