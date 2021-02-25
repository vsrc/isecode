package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func getFlags() []cli.Flag {
	return []cli.Flag {
		&cli.StringFlag{
			Name: "path",
			Aliases: []string{"p"},
			Usage: "relative path to folder where is your code which needs to be scanned and injected with isecode numbers",
			Destination: &path,
			DefaultText: "current directory",
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


func handlePath() {
	if path == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		path = wd

		log.Println("path set to current working directory: " + path)
	}
}