package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/KlatterAB/klatter-burton/cfg"
	"github.com/KlatterAB/klatter-burton/cmd"
	"github.com/KlatterAB/klatter-burton/db"
	"github.com/KlatterAB/klatter-burton/util"
)

var Version string

func init() {
	db.InitDB()
	db.InitStore()
	cfg.InitConfig()
}

func main() {
	opts := util.Options{
		Verbose:    false,
		ShowStatus: false,
		Project: util.Project{
			Name: "",
			ID:   "",
		},
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
		},
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "check in",
				Aliases: []string{"ci"},
				Usage:   "trigger check in sequence",
				Flags:   []cli.Flag{},
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 1 {
						fmt.Println("You need to supply project id as an argument")
						return nil
					}

					opts.Project.ID = c.Args().Get(0)
					return cmd.Checkin(opts)
				},
			},
			{
				Name:    "check out",
				Aliases: []string{"co"},
				Usage:   "trigger check out sequence",
				Flags:   []cli.Flag{},
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
			{
				Name:    "add time",
				Aliases: []string{"at"},
				Usage:   "add time in minutes to a project",
				Flags:   []cli.Flag{},
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 2 {
						fmt.Println("You need to supply two arguments: number of minutes and project id")
						return nil
					}

					minutes := c.Args().Get(0)
					projectId := c.Args().Get(1)
					return cmd.AddTime(minutes, projectId)
				},
			},
			{
				Name:    "add project",
				Aliases: []string{"ap"},
				Usage:   "add a new project to the database with project name and id",
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 2 {
						fmt.Println("You need to supply two arguments: project name and project id")
						return nil
					}

					name := c.Args().Get(0)
					id := c.Args().Get(1)
					return cmd.AddProject(name, id)
				},
			},
			{
				Name:    "check project",
				Aliases: []string{"cp"},
				Usage:   "show the checked in project name",
				Action: func(c *cli.Context) error {
					return cmd.CheckProject()
				},
			},
			{
				Name:    "work log",
				Aliases: []string{"wl"},
				Usage:   "get the work log for a project, optionally for a person and/or a specific date",
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 1 || c.Args().Len() == 3 {
						fmt.Println("You need to supply at least one arguments: project id (and optionally worker id and/or from date and to date)")
						return nil
					}
					params := cmd.WorkLogParams{
						ProjectID: "",
						WorkerID:  "",
						FromDate:  "",
						ToDate:    "",
					}

					params.ProjectID = c.Args().Get(0)
					params.WorkerID = c.Args().Get(1)
					params.FromDate = c.Args().Get(2)
					params.ToDate = c.Args().Get(3)

					return cmd.GetWorkLog(params)
				},
			},
			{
				Name:    "pair programming timer",
				Aliases: []string{"ppt"},
				Usage:   "start a pair programming timer with a set duration so you know when it's time to switch 'positions'",
				Action: func(c *cli.Context) error {
					defaultDuration, err := time.ParseDuration("30m")
					if err != nil {
						log.Panic(err)
					}
					params := cmd.TimerParams{
						Duration:      defaultDuration,
						SoundFilePath: "",
					}

					if c.Args().Len() == 2 {
						dur, err := time.ParseDuration(c.Args().Get(0))
						if err != nil {
							fmt.Println("That is not a correctly formatted duration. Use {n}m format, i.e 15m for 15 minutes.")
							log.Fatal(err)
						}
						params.Duration = dur
						params.SoundFilePath = c.Args().Get(1)
					}

					if c.Args().Len() == 1 {
						dur, err := time.ParseDuration(c.Args().Get(0))
						if err != nil {
							fmt.Println("That is not a correctly formatted duration. Use {n}m format, i.e 15m for 15 minutes.")
							log.Fatal(err)
						}
						params.Duration = dur
					}

					return cmd.StartTimer(params)
				},
			},
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
