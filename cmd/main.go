package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"code.heb12.com/heb12/bref"
	"code.heb12.com/heb12/heb12/config"
	"code.heb12.com/heb12/heb12/osis"
)

// printVerses prints all the verses with verse numbers
func printVerses(verses []string, ref bref.Reference) {
	for i, verse := range verses {
		fmt.Println(i+ref.From, verse)
	}
}

func main() {
	// Initialize osistool
	//osismanage, err := osistool.New(config.DataDirs.GratisSplit)

	app := &cli.App{
		Name:  "heb12",
		Usage: "Read the Bible",
		Commands: []*cli.Command{
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "Read a passage of scripture",
				Action: func(c *cli.Context) error {
					if c.Args().First() == "" {
						fmt.Println("Usage example: aiodl read \"John 3:16\"")
						return nil
					}
					ref := c.Args().First()
					reference, err := bref.Parse(ref)
					if err != nil {
						return err
					}
					scope := config.GetScope()
					dataDir, err := scope.DataDir()
					osisData, err := osis.Load(dataDir + "/" + config.DataDirs.GratisSplit + "/en/asv.xml")
					if err != nil {
						return err
					}
					text, err := osisData.GetVerses(osis.Reference{
						ID:      reference.ID,
						Chapter: reference.Chapter,
						From:    reference.From,
						To:      reference.To,
					})
					if err != nil {
						return err
					}
					printVerses(text, reference)

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
