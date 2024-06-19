package main

func main() {
	// Possibilta armazenar 2 valores ao mesmo tempo, não é algo recomendado
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
