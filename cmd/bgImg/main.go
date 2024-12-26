package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena/goImg/internal"
)

func main() {
	if err := internal.Start(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}
}
