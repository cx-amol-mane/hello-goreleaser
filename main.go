package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, GoReleaser! 🚀")
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s! Welcome to GoReleaser! 🚀\n", name)
}
