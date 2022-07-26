package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		usage()
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s \"sentence to add\"\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}
