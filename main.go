package main

import (
	"log"
	"os"
	"sort"

	"github.com/spf13/viper"
	cli "github.com/urfave/cli/v2"
)

var path, config string

func main() {

	app := &cli.App{
		Name: "isecode",
		Usage: "utility cli tool to populate your code with traceable and secure error codes",
	}

	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		
		handlePath()
		loadConfig()

		log.Println(viper.GetString("hello"))

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