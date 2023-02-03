package cmd

import (
	"fmt"
	"log"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
)

func AddProject(name, id string) error {
	err := db.AddProject(name, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	addedProject := fmt.Sprintf("Added new project %s with id %s\n", name, id)
	fmt.Printf(addedProject)

	if cfg.Cfg.Notifications {
		util.Notify("Added project", addedProject)
	}

	return nil
}
