package main

import (
	"fmt"

	"github.com/negasus/goval"
)

func main() {
	fmt.Printf("validator hello from cmd\n")

	v := goval.New()
	v.Generate()
}
