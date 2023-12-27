package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"time"
)

func initCli() {
	var filename string
	app := &cli.App{
		// app definition
		Name:     "ttt",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name: "Stefan Rohrbacher (thefenriswolf)",
			},
		},
		Usage: "CLI time tracker tool",
		//UsageText:            "ttt [global options] command [subcommand]",
		EnableBashCompletion: true,
		Suggest:              true,
		HideVersion:          false,
		// flags
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "database .csv file to use",
				Destination: &filename,
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
			// #todo: create helper function for monthly, weekly, and yearly reports
			&cli.Command{
				Name:        "graph",
				Aliases:     []string{"g"},
				Usage:       "graph worktime",
				Description: "If no command is specified ttt will report in weekly mode",
				Action: func(cCtx *cli.Context) error {
					graphHours(filename, 1)
					return nil
				},
				// subcommands for: graph
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:    "month",
						Usage:   "generate monthly graph",
						Aliases: []string{"m"},
						Action: func(cCtx *cli.Context) error {
							graphHours(filename, 1)
							return nil
						},
					},
					&cli.Command{
						Name:    "year",
						Usage:   "generate yearly graph",
						Aliases: []string{"y"},
						Action: func(cCtx *cli.Context) error {
							graphHours(filename, 2)
							return nil
						},
					},
				},
			},
			&cli.Command{
				Name:        "report",
				Aliases:     []string{"rep", "r"},
				Usage:       "report worktime",
				Description: "If no command is specified ttt will report in weekly mode",
				Action: func(cCtx *cli.Context) error {
					reportWorktime(filename, 1)
					return nil
				},
				// subcommands for: report
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:    "weekly",
						Usage:   "generate weekly report",
						Aliases: []string{"w"},
						Action: func(cCtx *cli.Context) error {
							reportWorktime(filename, 1)
							return nil
						},
					},
					&cli.Command{
						Name:    "monthly",
						Usage:   "generate monthly report",
						Aliases: []string{"m"},
						Action: func(cCtx *cli.Context) error {
							reportWorktime(filename, 2)
							return nil
						},
					},
					&cli.Command{
						Name:    "yearly",
						Usage:   "generate monthly report",
						Aliases: []string{"y"},
						Action: func(cCtx *cli.Context) error {
							reportWorktime(filename, 3)
							return nil
						},
					},
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

// helper report function
func reportWorktime(fn string, timeframe int) {
	switch {
	case timeframe == 1:
		weekReport(fn)
	case timeframe == 2:
		fmt.Println("monthly not implemented yet")
	case timeframe == 3:
		fmt.Println("yearly not implemented yet")
	default:
		fmt.Println("default not implemented yet")
		return
	}
}
