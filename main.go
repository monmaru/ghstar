package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var lang, sort, direction string
	app := cli.NewApp()
	app.Version = "1.0"
	app.Author = "monmaru"
	app.UsageText = "$ ghstar <GitHub User Name>"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Usage:       "Filter the repository by the name of the programming language",
			Destination: &lang,
		},
		cli.StringFlag{
			Name:        "sort, s",
			Usage:       "Sort. You can specify either created, updated or pushed",
			Value:       "created",
			Destination: &sort,
		},
		cli.StringFlag{
			Name:        "direction, d",
			Usage:       "Sorting direction. You can specify either desc or asc.",
			Value:       "desc",
			Destination: &direction,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			fmt.Print(c.App.UsageText)
			return nil
		}

		params := &params{
			lang:      lang,
			sort:      sort,
			direction: direction,
		}
		return listRepositories(c.Args().First(), params)
	}
	app.Run(os.Args)
}
