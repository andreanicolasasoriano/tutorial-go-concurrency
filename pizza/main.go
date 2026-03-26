package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received an order number %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++
		fmt.Printf("making pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("We're out of ingredients for pizza #%d\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("Cook rage quit while making pizza #%d\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {

	//keep track of pizza we are making
	var i = 0

	//run forever or until we receive a quit notification
	//try to make pizzas
	for {
		currentPizza := makePizza(i)
		i = currentPizza.pizzaNumber
		select {
		//we sent something to the .data channel
		case pizzaMaker.data <- *currentPizza:

		case quitChan := <-pizzaMaker.quit:
			close(pizzaMaker.data)
			close(quitChan)
			return
		}
		//try to make pizza
		//decision
	}

}

func main() {

	//seed random number generator
	rand.NewSource(time.Now().UnixNano())

	//print out a message
	color.Cyan("Pizzeria Andrea is open!!")
	color.Cyan("-------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run producer in the background
	go pizzeria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.message)
			} else {
				color.Red(i.message)
				color.Red("Customer is like wtf")
			}
		} else {
			color.Cyan("Pizzeria Andrea is closed!!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error closing channel", err)
			}
		}
	}

	color.Cyan("Done for the day!")
	color.Cyan("Pizzas made: %d, Pizzas failed: %d, Pizzas attempted: %d")

}
