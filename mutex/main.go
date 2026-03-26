package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	//bank balance variable
	var bankBalance int
	var balance sync.Mutex

	//print starting values
	fmt.Printf("Initial account balance: $%d.00 /n", bankBalance)
	//define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Part-time job", Amount: 50},
		{Source: "Investments", Amount: 100},
		{Source: "Side hustle", Amount: 25},
		{Source: "Other", Amount: 100},
	}

	wg.Add(len(incomes))
	//loop 52 times
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
				balance.Unlock()
			}
		}(i, income)
	}

	wg.Wait()

	//print final balance
	fmt.Printf("Final account balance: $%d.00 /n", bankBalance)
}
