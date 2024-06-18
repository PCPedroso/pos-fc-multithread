package main

import (
	"fmt"
	"sync"
	"time"
)

// Routine Principal
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)
	reader(ch, &wg)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
		wg.Done()
		time.Sleep(time.Second / 4)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// Ao fechar o canal sinaliza que nada mais vai ser feito e evita o deadlock
	close(ch)
}
