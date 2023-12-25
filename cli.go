package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"time"
)

func argParse() {
	var filename string
	app := &cli.App{
		// app definition
		Name:     "ttt",
		Version:  "0.1-alpha",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name: "Stefan Rohrbacher (thefenriswolf)",
			},
		},
		Usage:                "CLI time tracker tool",
		EnableBashCompletion: true,
		Suggest:              true,
		HideVersion:          true,
		// flags
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "file path to .csv file",
				Destination: &filename,
			},
			&cli.BoolFlag{
				Name:    "potato",
				Aliases: []string{"pot"},
				Value:   false,
				Usage:   "run in performance mode (not a potato mode)",
				Action: func(ctx *cli.Context, t bool) error {
					if t == true {
						fmt.Printf("not a potato mode: %t\n", t)
					}
					return nil
				},
			},
		},
		//commands
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "print cleaned up source file to stdout",
				Action: func(cCtx *cli.Context) error {
					prettyPrint(filename)
					return nil
				},
			},
			&cli.Command{
				Name:    "report",
				Aliases: []string{"rep", "r"},
				Usage:   "report worktime",
				Action: func(cCtx *cli.Context) error {
					reportWorktime(filename)
					return nil
				},
			},
		},
	}
	// sort flags and commands
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	// run flag parser and handle errors
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func reportWorktime(fn string) {
	weekReport(fn)
}
