package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"

	cli "github.com/urfave/cli/v2"
)

var path, config, matchFiles, matchString, prefix string

func main() {

	app := &cli.App{
		Name: "isecode",
		Usage: "utility cli tool to populate your code with traceable and secure error codes",
	}

	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		
		handlePath()
		loadConfig()
		getPrefix()

		// run process() fn from process.go for each file in this path
		err := filepath.Walk(path, process)
		if err != nil {
			panic(err)
		}

		return nil
	}

	app.Flags = getFlags()

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
