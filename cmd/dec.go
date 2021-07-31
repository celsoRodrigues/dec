package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/celsoRodrigues/dec/pkg/cmd"
	"github.com/celsoRodrigues/dec/pkg/conf"
)

var (
	home   string
	config conf.ViperConfig
)

func init() {

	var err error
	home, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("error while getting the home folder", err)
		os.Exit(1)
	}
	config, err = conf.NewConfigWithOpt("config.yaml", "yaml", home)
	if err != nil {
		fmt.Println("error creating the configuration object", err)
	}
	// err = config.ReadInConfig()
	// if err != nil {
	// 	fmt.Println("error reading the configuration", err)
	// }

	config.ReadInConfig()

}

func main() {

	flags := pflag.NewFlagSet("kubectl-dec", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdDec(config)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}

}
