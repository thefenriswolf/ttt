package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

const RELEASE = "v20240307"
const ARCH = string(runtime.GOARCH)
const OS = string(runtime.GOOS + " ")
const VERSION = RELEASE + ", built for: " + OS + ARCH

func initCli() {
	var filename string
	app := &cli.App{
		// app definition
		Name:     "ttt",
		Version:  VERSION,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name: "Stefan Rohrbacher (thefenriswolf)",
			},
		},
		Usage:                "CLI time tracker tool",
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
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "create a template file in the current directory",
				Action: func(cCtx *cli.Context) error {
					writeTemplate(filename)
					return nil
				},
			},
			&cli.Command{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "print cleaned up source file to stdout",
				Action: func(cCtx *cli.Context) error {
					prettyPrint(filename, settings)
					return nil
				},
			},
			&cli.Command{
				Name:        "export",
				Aliases:     []string{"e"},
				Usage:       "export worktime report as PDF",
				Description: "If no command is specified ttt will report in weekly mode",
				Action: func(cCtx *cli.Context) error {
					exportWorktime(filename, 2)
					return nil
				},
				// subcommands for: graph
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:    "month",
						Usage:   "generate monthly PDF report",
						Aliases: []string{"m"},
						Action: func(cCtx *cli.Context) error {
							exportWorktime(filename, 1)
							return nil
						},
					},
					&cli.Command{
						Name:    "week",
						Usage:   "generate weekly graph",
						Aliases: []string{"y"},
						Action: func(cCtx *cli.Context) error {
							exportWorktime(filename, 2)
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
						Name:    "week",
						Usage:   "generate weekly report",
						Aliases: []string{"w"},
						Action: func(cCtx *cli.Context) error {
							reportWorktime(filename, 1)
							return nil
						},
					},
					&cli.Command{
						Name:    "month",
						Usage:   "generate monthly report",
						Aliases: []string{"m"},
						Action: func(cCtx *cli.Context) error {
							reportWorktime(filename, 2)
							return nil
						},
					},
					&cli.Command{
						Name:    "year",
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
		weekReport(fn, settings)
	case timeframe == 2:
		monthReport(fn, settings)
	case timeframe == 3:
		fmt.Println("yearly not implemented yet")
	default:
		fmt.Println("default not implemented yet")
		return
	}
}

func exportWorktime(fn string, timeframe int) {
	switch {
	case timeframe == 1:
		createPDF(fn, "month")
	case timeframe == 2:
		createPDF(fn, "week")
	default:
		createPDF(fn, "week")
		return
	}
}
