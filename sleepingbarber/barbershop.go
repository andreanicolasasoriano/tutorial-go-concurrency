package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientChan      chan string
	Open            bool
}

func (shop *BarberShop) AddBarber(name string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s is ready to work", name)

		for {
			if len(shop.ClientChan) == 0 {
				color.Yellow("No clients, so %s is going to sleep", name)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientChan //shop open will return false if channel is closed

			if shopOpen {
				if isSleeping {
					color.Yellow("%s is waking up", name)
					isSleeping = false
				}
				shop.CutHair(name, client)

			} else {
				shop.SendBarberHome(name)
				return
			}
		}
	}()
}

func (shop *BarberShop) CutHair(barber string, client string) {
	color.Green("%s is cutting hair for %s", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is done cutting hair for %s", barber, client)
}

func (shop *BarberShop) SendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.NumberOfBarbers--
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) CloseShopForDay() {
	color.Cyan("Closing shop for the day")
	close(shop.ClientChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	color.Green("All barbers are done for the day. Closing shop")
	close(shop.BarbersDoneChan)
}

func (shop *BarberShop) AddClient(client string) {
	color.Green("Client %s has entered the shop", client)
	if shop.Open {
		select {
		case shop.ClientChan <- client:
			color.Yellow("Client %s takes seat in waiting room", client)
		default:
			color.Red("Shop is full, so %s has left the shop", client)
		}
	} else {
		color.Red("Shop is closed, so %s has left the shop", client)
	}
}
