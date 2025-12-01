package pkg

import (
	"fmt"
	"os"
	"time"
)

type Information struct {
	command_name []string
	time_start   [6]string
	time_end     [6]string
}

func Month2Num(month string) int {
	return map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}[month]
}

func GetCommandName(information *Information) {
	information.command_name = os.Args[1:]
}

func GetTime(information *Information, flag string) {
	now := time.Now()
	year, raw_month, day := now.Date()
	hour, minute, second := now.Clock()

	month := Month2Num(raw_month.String())

	if flag == "start" {
		information.time_start = [6]string{fmt.Sprint(year), fmt.Sprint(month), fmt.Sprint(day), fmt.Sprint(hour), fmt.Sprint(minute), fmt.Sprint(second)}
	} else if flag == "end" {
		information.time_end = [6]string{fmt.Sprint(year), fmt.Sprint(month), fmt.Sprint(day), fmt.Sprint(hour), fmt.Sprint(minute), fmt.Sprint(second)}
	} else {
		panic("Invalid Flag")
	}
}

func InformationToString(information Information) string {
	return fmt.Sprintf(
		"The command name is %v\nThe start time is %s-%s-%s %s:%s:%s\nThe end time is %s-%s-%s %s:%s:%s\n",
		information.command_name,
		information.time_start[0], information.time_start[1], information.time_start[2], information.time_start[3], information.time_start[4], information.time_start[5],
		information.time_end[0], information.time_end[1], information.time_end[2], information.time_end[3], information.time_end[4], information.time_end[5],
	)
}
