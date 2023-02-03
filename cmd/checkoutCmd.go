package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/PatrikOlin/skvs"

	"github.com/PatrikOlin/butler-burton/cfg"
	"github.com/PatrikOlin/butler-burton/db"
	"github.com/PatrikOlin/butler-burton/util"
)

func Checkout(opts util.Options) error {
	var valUnix int64
	if err := db.Store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	}
	var valRound time.Time
	if err := db.Store.Get("checkinRounded", &valRound); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return err
	} else if err != nil {
		log.Fatal(err)
		return err
	} else {
		tci := CalculateTimeCheckedIn(valUnix)
		fmt.Println("Ok, checking out.")
		checkedInMsg := fmt.Sprintf("Time spent checked in: %s\n", tci)
		fmt.Println(checkedInMsg)

		de := time.Unix(valUnix, 0).Local().Format("15:04:05")
		dr := valRound.Local().Format("15:04:05")

		d := (15 * time.Minute)

		checkedInDurMsg := fmt.Sprintf("You checked in at: %s (%s)\n", de, dr)
		fmt.Println(checkedInDurMsg)

		if cfg.Cfg.Notifications {
			n := fmt.Sprintf("%s%s \n", checkedInMsg, checkedInDurMsg)
			util.Notify("Checking out \n", n)
		}

	}
	return nil
}

func CalculateTimeCheckedIn(checkin int64) time.Duration {
	t1 := time.Unix(checkin, 0)
	t2 := time.Since(t1)

	d := (1000 * time.Millisecond)
	trunc := t2.Truncate(d)
	return trunc
}
