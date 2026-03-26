package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "From chan ch")

		//simulate long process
		time.Sleep(time.Second * 1)
	}
}

func main() {

	ch := make(chan int, 10) //channel with fixed size of 10 Great for rate limiting;

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending ", i, " to channel")
		ch <- i
		fmt.Println("Sent ", i, " to channel")
	}

	fmt.Println("Done!")

	close(ch)

}
