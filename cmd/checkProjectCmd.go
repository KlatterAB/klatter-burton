package cmd

import (
	"fmt"
	"log"

	"github.com/PatrikOlin/skvs"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
)

func CheckProject() error {
	var valUnix int64
	if err := db.Store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	}
	name, err := GetProject(valUnix)
	if err != nil {
		log.Fatal(err)
		return err
	}

	checkedInProject := fmt.Sprintf("Currently checked in on project %s\n", name)

	fmt.Printf(checkedInProject)
	if cfg.Cfg.Notifications {
		util.Notify("Project", checkedInProject)
	}

	return nil
}
