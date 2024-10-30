package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// create
	app := cli.NewApp()

	// add flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name",
			Value: "stranger",
			Usage: "your beautyful name",
		},
		cli.IntFlag{
			Name:  "age",
			Value: 0,
			Usage: "your graceful age",
		},
		cli.StringFlag{
			Name:  "city",
			Value: "any",
			Usage: "your city pleasse",
		},
	}

	// This function parses and brings data in cli.Context struct
	app.Action = func(c *cli.Context) error {
		log.Printf("Hello %s (%d years), from %s city Welcome to the command line with CLI app", c.String("name"), c.Int("age"), c.String("city"))
		return nil
	}

	// Pass os.Args to cli app to parse content
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

	// go run . --name Victoria --age 371 --city 'the sky'
}
