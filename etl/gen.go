package etl

import (
	"flag"
	"os"
	"path"
	"time"

	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/sebnyberg/etl-go/avro"
)

type generator struct {
	filename        string
	numFakes        int
	recordsPerblock int
}

func newGeneratorFromArgs(args []string) (*generator, error) {
	gen := &generator{}

	fs := flag.NewFlagSet("generate", flag.ContinueOnError)

	fs.StringVar(&gen.filename, "gen-path", "tmp/purchases.avro", "generated avro file path")
	fs.IntVar(&gen.numFakes, "n", 1e3, "number of generated fakes")
	fs.IntVar(&gen.recordsPerblock, "r", 1e3, "records per write-block")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return gen, nil
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
	aw, err := avro.NewPurchaseWriter(f, container.Null, 1e3)
	for i := 0; i < g.numFakes; i++ {
		aw.WriteRecord(&avro.Purchase{})
	}
	time.Sleep(1 * time.Second)

	return nil
}
