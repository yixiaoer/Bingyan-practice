package main

import (
	"fmt"
	"os"
	"project/BasedGoScript/repl"
)

func main() {

	fmt.Println("Hello! This is a language based on Golang")
	fmt.Println("Now use Start Game to begin your journey")
	fmt.Println("Have a good time!")
	repl.Start(os.Stdin, os.Stdout)
}
