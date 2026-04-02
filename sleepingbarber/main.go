package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1 * time.Second
var timeOpen = 10 * time.Second

func main() {

	randomiser := rand.New(rand.NewSource(time.Now().UnixNano()))

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("--------------------------------")

	clientChan := make(chan string, seatingCapacity) //channel buffered with seating capacity
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientChan:      clientChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	shop.AddBarber("Frank")
	shop.AddBarber("John")
	shop.AddBarber("Jane")
	shop.AddBarber("Jim")
	shop.AddBarber("Jill")
	shop.AddBarber("Jack")
	shop.AddBarber("Bob")

	shopClosing := make(chan bool)
	closedShop := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.CloseShopForDay()
		closedShop <- true
	}()

	//add clients
	i := 1

	go func() {
		for {
			randomMilliSeconds := randomiser.Intn(2000) % (2 * arrivalRate)
			select {
			case <-closedShop:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliSeconds)):
				shop.AddClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()

	<-shopClosing
	color.Green("The shop is closed for the day. Bye!")
}
