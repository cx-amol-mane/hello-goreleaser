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
	fmt.Println("This is a simple Go program that prints a greeting message.")
}

// This is a simple Go program that prints a greeting message.
// You can run it with an optional name argument, e.g., `go run main.go
// Alice` to greet Alice instead of the default "World".
// --- IGNORE ---
// --- IGNORE ---
// --- IGNORE ---
// --- IGNORE ---
