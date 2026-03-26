package main

import (
	"fmt"
	"time"
)

func serv1(ch chan string) {

	for {
		time.Sleep(time.Second * 3)
		ch <- "Ping server 1"
	}
}
func serv2(ch chan string) {
	for {
		time.Sleep(time.Second * 1)
		ch <- "Ping server 2"
	}
}

func main() {

	fmt.Println("Select with channels")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go serv1(channel1)
	go serv2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("Case 1", s1)
		case s2 := <-channel1:
			fmt.Println("Case 2", s2)
		case s3 := <-channel2:
			fmt.Println("Case 3", s3)
		case s4 := <-channel2:
			fmt.Println("Case 3", s4)

		}

	}

}
