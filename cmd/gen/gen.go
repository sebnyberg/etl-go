package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/sebnyberg/etl-go/avro"
)

type generator struct {
	filename        string
	numFakes        int
	recordsPerblock int
}

func main() {
	filename := flag.String("gen-path", "tmp/purchases.avro", "generated avro file path")
}

func newGeneratorFromArgs(args []string) *generator {
	gen := &generator{}
	fmt.Println(args)

	fs := flag.NewFlagSet("gen", flag.ContinueOnError)

	fs.StringVar(&gen.filename, "gen-path", "tmp/purchases.avro", "generated avro file path")
	fs.IntVar(&gen.numFakes, "n", 1e3, "number of generated fakes")
	fs.IntVar(&gen.recordsPerblock, "r", 1e3, "records per write-block")

	fs.Parse(args)

	return gen
}

func (g *generator) run() error {
	// Create the output directory & file
	if err := os.MkdirAll(path.Dir(g.filename), 0744); err != nil {
		return err
	}
	f, err := os.OpenFile(g.filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// Stream fakes to the file
	fmt.Println("generating", g.numFakes, "fakes")
	aw, err := avro.NewPurchaseWriter(f, container.Null, 1e3)
	for i := 0; i < g.numFakes; i++ {
		aw.WriteRecord(&avro.Purchase{
			Id: "abc123",
		})
	}
	return aw.Flush()
}
