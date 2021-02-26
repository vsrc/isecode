package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)


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