package main

import (
	"fmt"
	"sync"
	"time"
)

func weekReport(fn string, s Settings) {
	db := readParseSort(fn, s)
	weekNumbers := weekList(db)
	for i := range weekNumbers {
		for j := range db {
			if db[j].kw == weekNumbers[i] {
				date := db[j].date
				duration := db[j].duration
				dateF := date.Format(s.datefmt)
				durationF := duration.String()
				fmt.Printf("%s: %s\n", dateF, durationF)
			}
		}
		kW := weekNumbers[i]
		sumOverWeek := weekSum(db, kW)
		overtime := calculateOvertime(settings.weeklyHours, sumOverWeek)
		fmt.Printf("=========================\nSummary KW%d:\n-------------------------\nWeekly sum: %s\n", kW, sumOverWeek)
		fmt.Printf("Weekly overtime: %s\n=========================\n\n\n", overtime)
	}
}

func formatData(timestr time.Time, formatstr string) string {
	return timestr.Format(formatstr)
}

func prettyPrint(fn string, s Settings) {
	dS := readParseSort(fn, s)
	var wg sync.WaitGroup
	for i, _ := range dS {
		wg.Add(3) // add 3 calls to stack
		date := dS[i].date
		var dateF string
		go func() {
			dateF = formatData(date, s.datefmt)
			wg.Done()
		}()
		startTime := dS[i].start
		var startTimeF string
		go func() {
			startTimeF = formatData(startTime, s.timefmt)
			wg.Done()
		}()
		endTime := dS[i].end
		var endTimeF string
		go func() {
			endTimeF = formatData(endTime, s.timefmt)
			wg.Done()
		}()
		jobName := dS[i].job
		wg.Wait()
		fmt.Printf("%s %s %s %s\n", dateF, startTimeF, endTimeF, jobName)
	}
}