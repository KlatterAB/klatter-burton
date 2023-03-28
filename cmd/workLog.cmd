package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/KlatterAB/klatter-burton/cfg"
	"github.com/KlatterAB/klatter-burton/db"
	"github.com/KlatterAB/klatter-burton/util"
)

type WorkLogParams struct {
	ProjectID string
	WorkerID  string
	FromDate  string
	ToDate    string
}

func GetWorkLog(params WorkLogParams) error {
	res, err := db.GetWorkLog(params.ProjectID, params.WorkerID, params.FromDate, params.ToDate)
	if err != nil {
		log.Fatal(err)
		return err
	}

	var workLog strings.Builder
	if params.WorkerID != "" {
		workLog.WriteString(fmt.Sprintf("Hours worked on project %s by %s: %s hours\n", params.ProjectID, params.WorkerID, res))
	} else {
		workLog.WriteString(fmt.Sprintf("Hours worked on project %s: %s hours\n", params.ProjectID, res))
	}

	if params.ToDate != "" && params.FromDate != "" {
		workLog.WriteString(fmt.Sprintf("between %s and %s\n", params.FromDate, params.ToDate))
	}

	fmt.Printf(workLog.String())

	if cfg.Cfg.Notifications == true {
		util.Notify("Work log", workLog.String())
	}

	return nil
}
