package main

import (
	"fmt"
	"time"
)

func worker(ch chan int, id int) {
	for i := range ch {
		fmt.Printf("Worker ID %d received %d\n", id, i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	ch := make(chan int)

	// // Disparado dois workers para trabalhar simultaneamente
	// go worker(ch, 1)
	// go worker(ch, 2)

	// Definido a quandidade de workers simultaneos
	workers := 1000
	for i := range workers {
		go worker(ch, i)
	}

	for i := range 100000 {
		ch <- i
	}
}
