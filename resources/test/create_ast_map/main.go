package main

type StructChild struct {
}

type StructRoot struct {
	child StructChild
}

func (r StructRoot) rootMethod() {}
