package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() //when the function exists, decrement the wait group counter by 1
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	list := []string{"a", "b", "c", "d", "e", "f"}

	wg.Add(len(list))
	for index, item := range list {
		go printSomething(fmt.Sprintf("%d%s", index, item), &wg)
	}

	wg.Wait()

}
