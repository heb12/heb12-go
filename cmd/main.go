package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"code.heb12.com/Heb12/bref"
	"code.heb12.com/Heb12/heb12/config"
	"code.heb12.com/Heb12/heb12/osistool"
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
					osis, err := osistool.LoadOsis(dataDir + "/" + config.DataDirs.GratisSplit + "/en/asv.xml")
					if err != nil {
						return err
					}
					text, err := osis.GetVerses(reference)
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
