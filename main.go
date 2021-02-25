package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	cli "github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name: "isecode",
		Usage: "utility cli tool to populate your code with traceable and secure error codes",
	}

	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		fmt.Println("Run with -h to see list of available commands")
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}