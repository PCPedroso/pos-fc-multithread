package main

import (
	"fmt"
)

// Routine Principal
func main() {
	ch := make(chan string)
	go enviar("Hello", ch)
	receber(ch)
}

// chan<- : Indica um canal que só pode ser usado para enviar valores
func enviar(valor string, ch chan<- string) {
	ch <- valor
}

// <-chan : Indica um canal que só pode ser usado para receber valores
func receber(ch <-chan string) {
	fmt.Println(<-ch)
}
