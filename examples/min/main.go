package main

//go:generate ../../dist/goval -t Request

type Request struct {
	Name string `goval:"min=3;max=10"`
	Age  int    `goval:"min=18"`
}
