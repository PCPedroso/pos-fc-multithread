package main

import "fmt"

// Routine Principal
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

		// Se não for preenchido vai imprimir os 10 valores mas vai continuar dando deadlock!
		forever <- true
	}()

	<-forever
}
