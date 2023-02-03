package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/cmd"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
)

var Version string

func init() {
	db.InitDB()
	cfg.InitConfig()
}

func main() {
	opts := util.Options{
		Verbose:    false,
		ShowStatus: false,
		Project:    "",
	}

	app := &cli.App{
		Name:     "Klatter Burton",
		Usage:    "a smartish utility for reporting time spent working on Klatter projects",
		Version:  Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Patrik Olin",
				Email: "olin@klatter.se",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Value:       false,
				Usage:       "turn on verbose mode",
				Destination: &opts.Verbose,
			},
			// &cli.BoolFlag{
			// 	Name:        "loud",
			// 	Aliases:     []string{"l"},
			// 	Value:       false,
			// 	Usage:       "turn on loud mode for this command (send Teams message)",
			// 	Destination: &opts.Loud,
			// },
		},
		Commands: []*cli.Command{
			{
				Name:    "check in",
				Aliases: []string{"ci"},
				Usage:   "trigger check in sequence",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{"p"},
						Value:       "",
						Usage:       "which project to check in time to",
						Destination: &opts.Project,
					},
				},
				Action: func(c *cli.Context) error {
					return cmd.Checkin(opts)
				},
			},
			{
				Name:    "check out",
				Aliases: []string{"co"},
				Usage:   "trigger check out sequence",
				Flags:   []cli.Flag{
					// &cli.BoolFlag{
					// 	Name:        "catered",
					// 	Aliases:     []string{"c"},
					// 	Value:       false,
					// 	Usage:       "check BL-lunch field in time sheet for todays shift",
					// 	Destination: &opts.Catered,
					// },
					// &cli.BoolFlag{
					// 	Name:        "overtime",
					// 	Aliases:     []string{"o"},
					// 	Value:       false,
					// 	Usage:       "write overtime to overtime column",
					// 	Destination: &opts.Overtime,
					// },
				},
				Action: func(c *cli.Context) error {
					return cmd.Checkout(opts)
				},
			},
			{
				Name:    "check time",
				Aliases: []string{"ct"},
				Usage:   "get time spent checked in",
				Action: func(c *cli.Context) error {
					return cmd.CheckTime()
				},
			},
			// {
			// 	Name:     "time sheet",
			// 	Aliases:  []string{"ts"},
			// 	Usage:    "commands directly related to time sheet",
			// 	Category: "time sheet",
			// 	Subcommands: []*cli.Command{
			// 		{
			// 			Name:     "create",
			// 			Aliases:  []string{"c"},
			// 			Usage:    "download original time sheet and name it according to month and username",
			// 			Category: "time sheet",
			// 			Action: func(c *cli.Context) error {
			// 				return cmd.CreateNewReport()
			// 			},
			// 		},
			// 		{
			// 			Name:     "get",
			// 			Aliases:  []string{"g"},
			// 			Usage:    "get current time sheet filename",
			// 			Category: "time sheet",
			// 			Action: func(c *cli.Context) error {
			// 				return cmd.GetReportFilename()
			// 			},
			// 		},
			// 		{
			// 			Name:     "set",
			// 			Aliases:  []string{"s"},
			// 			Usage:    "set new time sheet filename",
			// 			Category: "time sheet",
			// 			Action: func(c *cli.Context) error {
			// 				return cmd.SetReportFilename(c.Args().First())
			// 			},
			// 		},
			// 		{
			// 			Name:     "upload",
			// 			Aliases:  []string{"u"},
			// 			Usage:    "upload the time sheet to sharepoint",
			// 			Category: "time sheet",
			// 			Action: func(c *cli.Context) error {
			// 				return cmd.UploadReport()
			// 			},
			// 		},
			// 		{
			// 			Name:     "download",
			// 			Aliases:  []string{"d"},
			// 			Usage:    "download the time sheet from sharepoint",
			// 			Category: "time sheet",
			// 			Action: func(c *cli.Context) error {
			// 				return cmd.DownloadReport()
			// 			},
			// 		},
			// 	},
			// },
			{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "commands directly related to config",
				Category: "config",
				Subcommands: []*cli.Command{
					{
						Name:     "edit",
						Aliases:  []string{"e"},
						Usage:    "edit config-file",
						Category: "config",
						Action: func(c *cli.Context) error {
							return cmd.EditConfig()
						},
					},
					{
						Name:     "print",
						Aliases:  []string{"p"},
						Usage:    "print config-file",
						Category: "config",
						Action: func(c *cli.Context) error {
							return cmd.PrintConfig()
						},
					},
				},
			},
			// {
			// 	Name:    "afk",
			// 	Aliases: []string{"a"},
			// 	Usage:   "set afk status",
			// 	Action: func(c *cli.Context) error {
			// 		return cmd.ToggleAFK(c.Args().Get(0), opts)
			// 	},
			// },
			// {
			// 	Name:    "exercise",
			// 	Aliases: []string{"ex"},
			// 	Usage:   "set exercise status",
			// 	Action: func(c *cli.Context) error {
			// 		return cmd.ToggleExercise(c.Args().Get(0), opts)
			// 	},
			// },
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
