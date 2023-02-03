package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
)

func Checkin(opts util.Options) error {
	// d := (15 * time.Minute)
	checkinUnix := time.Now().Unix()
	db.Store.Put("checkinUnix", checkinUnix)
	db.Store.Put(strconv.FormatInt(checkinUnix, 10), opts.Project.ID)

	de := time.Unix(checkinUnix, 0).Local().Format("15:04:05")
	checkinMsg := fmt.Sprintf("Ok, checked in at %s\n", de)
	fmt.Println(checkinMsg)

	if cfg.Cfg.Notifications {
		util.Notify("Checking in \n", checkinMsg)
	}

	return nil
}
