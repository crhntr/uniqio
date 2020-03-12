package main

import (
	"io"
	"os"

	"github.com/crhntr/uniqio"
)

func main() {
	if err := uniqio.Lines(os.Stdout, os.Stdin); err != nil {
		_, _ = io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}
