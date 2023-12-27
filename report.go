package main

import (
	"fmt"
	"time"
)

func weekReport(fn string) {
	db := readParseSort(fn)
	weekNumbers := weekList(db)
	for i := range weekNumbers {
		for j := range db {
			if db[j].kw == weekNumbers[i] {
				date := db[j].date
				duration := db[j].duration
				dateF := date.Format(datefmt_ddmmyyyy)
				durationF := duration.String()
				fmt.Printf("%s: %s\n", dateF, durationF)
			}
		}
		kW := weekNumbers[i]
		weekSum := weekSum(db, kW)
		overtime := calculateOvertime(weeklyHours, weekSum)
		fmt.Printf("=========================\nSummary KW%d:\n-------------------------\nWeekly sum: %s\n", kW, weekSum)
		fmt.Printf("Weekly overtime: %s\n=========================\n\n\n", overtime)
	}
}

func formatData(ch chan string, timestr time.Time, formatstr string) {
	ch <- timestr.Format(formatstr)
}

func prettyPrint(fn string) {
	dS := readParseSort(fn)
	ic := make(chan string)
	defer close(ic)
	for i, _ := range dS {
		date := dS[i].date
		go formatData(ic, date, datefmt_ddmmyyyy)
		dateF := <-ic

		startTime := dS[i].start
		go formatData(ic, startTime, timefmt_hhmm)
		startTimeF := <-ic

		endTime := dS[i].end
		go formatData(ic, endTime, timefmt_hhmm)
		endTimeF := <-ic

		jobName := dS[i].job
		fmt.Printf("%s %s %s %s\n", dateF, startTimeF, endTimeF, jobName)
	}
}
