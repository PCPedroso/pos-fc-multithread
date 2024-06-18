package main

import "time"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		time.Sleep(time.Second)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 2
	}()

	for range 3 {
		select {
		case msg1 := <-ch1:
			println("receives", msg1)
		case msg2 := <-ch2:
			println("receives", msg2)
		case <-time.After(time.Second * 3):
			print("timeout")
			// default:
			// 	println("default")
		}

		//		time.Sleep(time.Second)
	}
}
