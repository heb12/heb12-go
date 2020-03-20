package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"code.heb12.com/heb12/bref"
	"code.heb12.com/heb12/heb12/bible"
)

// printVerses prints all the verses with verse numbers
func printVerses(verses []string, ref bref.Reference) {
	for i, verse := range verses {
		fmt.Println(i+ref.From, verse)
	}
}

func main() {
	var translation string

	app := &cli.App{
		Name:    "heb12",
		Version: "v0.1.0",
		Usage:   "Read the Bible",
		Commands: []*cli.Command{
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "Read a passage of scripture",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "translation",
						Value:       "web",
						Aliases:     []string{"t"},
						Usage:       "Bible version to use",
						Destination: &translation,
					},
				},
				Action: func(c *cli.Context) error {
					if c.Args().First() == "" {
						fmt.Println("Usage example: aiodl read \"John 3:16\"")
						return nil
					}
					ref := c.Args().First()
					text, err := bible.Get(ref, translation)
					if err != nil {
						return err
					}

					reference, err := bref.Parse(ref)
					if err != nil {
						return err
					}

					printVerses(text, reference)

					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List downloaded Bible translations",
				Action: func(c *cli.Context) error {
					translations, err := bible.List()
					if err != nil {
						return err
					}
					languages, err := bible.ListLanguages()
					if err != nil {
						return err
					}

					for _, lang := range languages {
						fmt.Println(lang + ":")
						for _, version := range translations[lang] {
							fmt.Println("\t" + version)
						}
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
