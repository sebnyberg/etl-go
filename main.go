package main

import (
	"os"

	"github.com/sebnyberg/etl-go/etl"
)

func main() {
	os.Exit(etl.CLI(os.Args[1:]))
}
