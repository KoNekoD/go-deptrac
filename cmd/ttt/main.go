package main

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/commands_options"
	"github.com/alexflint/go-arg"
)

func main() {
	options := commands_options.InitOptions{}
	arg.MustParse(&options)

	fmt.Println(options)
}
