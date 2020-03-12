package main

import (
	"flag"
	"io"
	"os"

	"github.com/crhntr/uniqio"
)

func main() {
	var watch bool
	flag.BoolVar(&watch, "w", false, "watch")
	flag.Parse()

	fn := uniqio.Lines

	if watch {
		fn = uniqio.WatchLines
	}

	if err := fn(os.Stdout, os.Stdin); err != nil {
		_, _ = io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}
}
