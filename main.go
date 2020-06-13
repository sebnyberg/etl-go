package main

import (
	"os"

	"github.com/pkg/profile"
	"github.com/sebnyberg/etl-go/etl"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// p := profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	// p := profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	// defer p.Stop()
	os.Exit(etl.CLI(os.Args[1:]))
}
