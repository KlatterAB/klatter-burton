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
	d := (15 * time.Minute)
	rounded := time.Now().Local().Round(d)
	checkinUnix := time.Now().Unix()
	db.Store.Put("checkinUnix", checkinUnix)
	db.Store.Put("checkinRounded", rounded)
	db.Store.Put(strconv.FormatInt(checkinUnix, 10), opts.Project)

	de := time.Unix(checkinUnix, 0).Local().Format("15:04:05")
	dr := rounded.Format("15:04:05")
	checkinMsg := fmt.Sprintf("Ok, checked in at %s (%s)\n", de, dr)
	fmt.Println(checkinMsg)

	if cfg.Cfg.Notifications {
		util.Notify("Checking in \n", checkinMsg)
	}

	return nil
}
