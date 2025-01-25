package main

const (
	input = iota
	output
)

type Variable struct {
	Name  string
	Value interface{}
}

type Pin struct {
	Name  string
	Value interface{}
	Mode  int
}
