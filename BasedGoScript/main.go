package main

import (
	"fmt"
	"os"

	"project/BasedGoScript/repl"
)

func main() {

	fmt.Printf("Hello! This is a language based on Golang\n")
	fmt.Printf("Now use Start Game to begin your journey\n")
	fmt.Printf("Have a good time!\n")
	repl.Start(os.Stdin, os.Stdout)
}
