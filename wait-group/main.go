package main

import (
	"fmt"
	"sync"
	"time"
)

func task(valor string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, valor)
		time.Sleep(time.Second / 1)
		wg.Done()
	}
}

func main() {
	wg := sync.WaitGroup{}
	// Cada task roda 10x + as 5x do meu principal, com isso foi sinalizado para o
	// WaitGroup que deve aguardar as 25 operações serem executadas
	wg.Add(25)

	go task("A", &wg)
	go task("B", &wg)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(time.Second / 1)
			// Cada vez que passa informa para o wg que uma operação foi executada
			wg.Done()
		}
	}()
	// Aguarda até que não tenham mais operações para serem executadas
	wg.Wait()
}
