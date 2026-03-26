package main

import (
	"fmt"
	"strings"
	"time"
)

func shout(ping chan string, pong chan string) {
	for {
		s := <-ping                                   //sending value from ping
		pong <- fmt.Sprintf("%s", strings.ToUpper(s)) //sending value to pong
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	time.Sleep(time.Second * 10)

	fmt.Println("Type something then hit ENTER. Type Q to quit")

	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		response := <-pong //assigning value from pong

		fmt.Println("Pong: ", response)
	}

	fmt.Println("Done. Closing channels")
	close(ping) //must always close channels
	close(pong)

}
