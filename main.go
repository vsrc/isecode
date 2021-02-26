package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	cli "github.com/urfave/cli/v2"
)

var path, config, matchFiles, matchString string

func main() {

	app := &cli.App{
		Name: "isecode",
		Usage: "utility cli tool to populate your code with traceable and secure error codes",
	}

	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		
		handlePath()
		loadConfig()

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


func process(filePath string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match(matchFiles, fi.Name())

	if err != nil {
		return err
	}

	if matched {
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		matchExp := regexp.MustCompile(matchString)

		for matchExp.MatchString(string(file)) == true {
			codeNumber := getCodeNumber()

			newContents := strings.Replace(string(file), matchString, codeNumber, 1)

			err = ioutil.WriteFile(filePath, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}

			file, err = ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}

		}

		log.Println("Processed file: " + filePath)

	}

	return nil
}

func getCodeNumber() string {
	isecode := viper.GetInt("LAST_NUMBER") + 1
	viper.Set("LAST_NUMBER", isecode)
	viper.WriteConfig()
	return strconv.Itoa(isecode)
}