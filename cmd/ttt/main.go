package main

import "fmt"

type TTT struct {
}

func main() {
	t := &TTT{}

	fmt.Printf("%T\n", t)
}
