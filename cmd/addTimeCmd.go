package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/KlatterAB/klatter-burton/cfg"
	"github.com/KlatterAB/klatter-burton/db"
	"github.com/KlatterAB/klatter-burton/util"
)

func AddTime(minutes, projectId string) error {
	m, err := strconv.Atoi(minutes)
	err = db.AddTimeToProject(m, projectId, cfg.Cfg.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	addedTime := fmt.Sprintf("Added %v minutes to project with id %s\n", minutes, projectId)
	fmt.Printf(addedTime)

	if cfg.Cfg.Notifications {
		util.Notify("Added time", addedTime)
	}

	return nil
}
