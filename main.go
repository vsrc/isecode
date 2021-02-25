package main

import (
	cli "github.com/urfave/cli/v2"
)

func main() {

	_ = &cli.App{
		Name: "isecode",
		Usage: "utility cli tool to populate your code with traceable and secure error codes",
	}


}