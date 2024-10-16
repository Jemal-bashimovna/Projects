package main

import (
	"fmt"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}

func parseTime() time.Time {
	date := "Tue, 09/22/1995, 13:00"
	layout := "Mon, 01/02/2006, 15:04"

	t, _ := time.Parse(layout, date)
	fmt.Println(t)
	return t
}

func main() {
	parseTime()
	fmt.Println(AddGigasecond(time.Time{}))
}
