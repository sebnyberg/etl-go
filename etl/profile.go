package etl

import (
	"flag"

	"github.com/pkg/profile"
)

func profileStop(args []string) func() {
	fs := flag.NewFlagSet("profile", flag.ContinueOnError)

	var (
		cpu bool
		mem bool
	)

	profiles := []interface{ Stop() }{}

	fs.BoolVar(&cpu, "cpu", false, "run cpu profile")
	fs.BoolVar(&mem, "mem", false, "run memory profile")

	if err := fs.Parse(args); err != nil {
		panic("failed to parse flags")
	}

	if mem {
		p := profile.Start(profile.MemProfile, profile.ProfilePath("."))
		profiles = append(profiles, p)
	}

	if cpu {
		p := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
		profiles = append(profiles, p)
	}

	return func() {
		for _, p := range profiles {
			p.Stop()
		}
	}
}
