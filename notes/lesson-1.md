# Goroutines
- Runs concurrently with primary program
- Every main function has a goroutine
- Goroutine is a super lightweight thread;

## Problems with goroutines
- sometimes your main routine finishes before your child routines do

## Wait groups
- lets you wait for a goroutine to finish before terminating parent routine;

```go
func main() {

	var wg sync.WaitGroup

	list := []string{"a", "b", "c", "d", "e", "f"}

	wg.Add(len(list))
	for index, item := range list {
		go printSomething(fmt.Sprintf("%d%s", index, item), &wg)
	}

	wg.Wait()

}
```
