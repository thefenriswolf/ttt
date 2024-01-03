package main

import (
	"fmt"
	"sync"
	"time"
)

func inttoMonth(month int) string {
	var monthMap = make(map[int]string)
	monthMap[1] = "January"
	monthMap[2] = "February"
	monthMap[3] = "March"
	monthMap[4] = "April"
	monthMap[5] = "May"
	monthMap[6] = "June"
	monthMap[7] = "July"
	monthMap[8] = "August"
	monthMap[9] = "September"
	monthMap[10] = "October"
	monthMap[11] = "November"
	monthMap[12] = "December"
	return monthMap[month]
}

func monthReport(fn string, s Settings) {
	db := readParseSort(fn, s)
	months := monthList(db)
	for i := range months {
		for j := range db {
			mn := (db[j].date).Month()
			monthNum := int(mn)
			if monthNum == months[i] {
				date := db[j].date
				duration := db[j].duration
				dateF := date.Format(s.datefmt)
				durationF := duration.String()
				fmt.Printf("%s: %s\n", dateF, durationF)
			}
		}
		m := months[i] // current month
		var wg sync.WaitGroup
		wg.Add(2) // add 2 calls to wg stack
		var sumOverMonth time.Duration
		//overtime := calculateOvertime(settings.weeklyHours*4, sumOverMonth) // #todo: fix
		var mString string
		go func() {
			sumOverMonth = monthSum(db, m)
			wg.Done()
		}()
		go func() {
			mString = inttoMonth(m)
			wg.Done()
		}()
		wg.Wait() // wait for goroutines to finish before printing report
		fmt.Printf("=========================\nSummary of %s:\n-------------------------\nMonthly sum: %s\n\n\n", mString, sumOverMonth)
		//fmt.Printf("Monthly overtime: %s\n=========================\n\n\n", overtime)
	}
	fmt.Printf("\nOvertime calculation is currently only supported for weekly reports!\n")
}

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
	for i := range dS {
		wg.Add(3) // add 3 calls to wg stack
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
