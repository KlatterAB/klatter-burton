package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/KlatterAB/klatter-burton/db"
	"github.com/PatrikOlin/skvs"
)

func GetMonthFolderReplacer() *strings.Replacer {
	r := strings.NewReplacer(
		"January", "01 Januari",
		"February", "02 Februari",
		"March", "03 Mars",
		"April", "04 April",
		"May", "05 Maj",
		"June", "06 Juni",
		"July", "07 Juli",
		"August", "08 Augusti",
		"September", "09 September",
		"October", "10 Oktober",
		"November", "11 November",
		"December", "12 December",
	)
	return r
}

func GetMonthFileReplacer() *strings.Replacer {
	r := strings.NewReplacer(
		"January", "Jan",
		"February", "Feb",
		"March", "Mars",
		"April", "Apr",
		"May", "Maj",
		"June", "Juni",
		"July", "Juli",
		"August", "Aug",
		"September", "Sep",
		"October", "Okt",
		"November", "Nov",
		"December", "Dec",
	)
	return r
}

func GetProject(checkinTime int64) (string, error) {
	var name string
	if err := db.Store.Get(strconv.FormatInt(checkinTime, 10), &name); err == skvs.ErrNotFound {
		fmt.Println("not found")
		return "", err
	} else if err != nil || name == "" {
		return "", err
	}

	return name, nil
}
