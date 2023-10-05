package etl

import (
	"fmt"
	"os"
)

// CLI runs the command and returns the exit status
func CLI(args []string) int {
	if len(args) < 1 {
		return 2
	}

	defer profileStop(args[1:])()

	var err error
	switch args[0] {
	case "gen":
		err = cliGen(args[1:])
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
	}
	return 0
}

func cliGen(args []string) error {
	gen := newGeneratorFromArgs(args)
	if err := gen.run(); err != nil {
		return err
	}
	return nil
}
