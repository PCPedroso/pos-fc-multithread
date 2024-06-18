package main

import (
	"fmt"
)

// Routine Principal
func main() {
	canal := make(chan string)

	// Segunda routine
	go func() {
		canal <- "Olá mundo"
	}()

	// Routine Principal
	fmt.Println(<-canal)
}
