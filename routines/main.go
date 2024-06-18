package main

import (
	"fmt"
	"time"
)

func task(valor string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, valor)
		time.Sleep(time.Second / 1)
	}
}

func main() {
	go task("A")
	go task("B")

	//Gambiara pra esperar as routines acabarem
	time.Sleep(time.Second * 6)
}
