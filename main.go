package main

import (
	"fmt"
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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{"p"},
						Value:       "",
						Usage:       "id of the project to check in time for",
						Destination: &opts.Project.ID,
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
