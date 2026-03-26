package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w // hiijack os.Stdout so that it saves the data into the pipe instead of into actual Std out.

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	testResult := string(result)

	fmt.Println(testResult)

	os.Stdout = stdOut

	if !strings.Contains(testResult, "$40300.00") {
		t.Error("Expected $40300.00 balance")
	}

}
