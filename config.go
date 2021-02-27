package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
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

func loadConfig() {

	var dir, file string

	if config != "" {
		dir, file = filepath.Split(config)
		file = strings.TrimSuffix(file, filepath.Ext(file))
	} else {
		file = "isecode"
	}

	if dir == "" {
		dir = "."
	}

    viper.SetConfigName(file) // name of config file (without extension)
    viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
    viper.AddConfigPath(dir)    // optionally look for config in the working directory
    err := viper.ReadInConfig() // Find and read the config file

	if err != nil {
		log.Println(err)
	}

	matchFiles = viper.GetString("MATCH_FILES")
	matchString = viper.GetString("MATCH_STRING")
}

func getPrefix() {
	prefix = viper.GetString("CODE_PREFIX")
}