package main

import (
	"fmt"
)

// Routine Principal
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// Ao fechar o canal sinaliza que nada mais vai ser feito e evita o deadlock
	close(ch)
}
