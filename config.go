package main

import cli "github.com/urfave/cli/v2"

func getFlags() []cli.Flag {
	return []cli.Flag {
		&cli.StringFlag{
			Name: "path",
			Aliases: []string{"p"},
			Usage: "path where blah blah blah",
			Destination: &path,
		},
		&cli.StringFlag{
			Name: "config",
			Aliases: []string{"c"},
			Usage: "path to the config file",
			Destination: &config,
			DefaultText: "isecode.json",
		},
	}
}