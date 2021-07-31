package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/celsoRodrigues/dec/pkg/cmd"
)

func main() {

	flags := pflag.NewFlagSet("kubectl-dec", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdDec()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}

}
