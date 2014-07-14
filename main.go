package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func FromBTC(value string) string {
	return fmt.Sprintf("TBD\n")
}

func ToBTC(value string) string {
	return fmt.Sprintf("TBD\n")
}

func main() {
	app := cli.NewApp()
	app.Name = "gobit"
	app.Usage = "Perform currency conversions between major currencies and Bitcoin."
	app.Version = "0.1"
	app.Action = func(c *cli.Context) {
		fmt.Printf("Args: %v", c.Args())
		return
	}
	app.Run(os.Args)
}
