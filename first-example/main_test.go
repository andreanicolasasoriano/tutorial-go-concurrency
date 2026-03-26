package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w // redirect stdout to the pipe

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("test", &wg)

	wg.Wait()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	fmt.Println("output: ", output)
	// fmt.Println("stdOut: ", stdOut)

	os.Stdout = stdOut

	if !strings.Contains(output, "test") {
		t.Errorf("expected to find 'test' in the output")
	}
}
