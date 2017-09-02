package cli

import (
	"time"

	"github.com/urfave/cli"
)

var (
	// RootCommands is a list of subcommands for the application.
	RootCommands = []cli.Command{}

	// RootFlags is a list of flags for the application.
	RootFlags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "application configuration file",
			EnvVar: "CONFIG",
			Value:  "config.json",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "add this flag to enable debug mode",
		},
	}
)

// RootAction is executing when program called without any subcommand.
func RootAction(c *cli.Context) error {
	for {
		log.Print("hello")
		time.Sleep(1 * time.Second)
	}
}
