package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const csvDelim = ' '
const csvComment = '#'
const datefmt_ddmmyyyy = "02.01.2006" // dd.mm.yyyy
const timefmt_hhmm = "1504"           // hhmm
const timefmt_hhmmss = "150405"       // hhmmss
const weeklyHours = time.Duration(30 * time.Hour)

type Entry struct {
	date     time.Time
	kw       int
	start    time.Time
	end      time.Time
	duration time.Duration
	job      string
}

func weekReport(fn string) {
	dS := readParseSort(fn)
	weekNumbers := weekList(dS)
	for i := range weekNumbers {
		for j := range dS {
			if dS[j].kw == weekNumbers[i] {
				date := dS[j].date
				duration := dS[j].duration
				dateF := date.Format(datefmt_ddmmyyyy)
				durationF := duration.String()
				fmt.Printf("%s: %s\n", dateF, durationF)
			}
		}
		weekSum := weekSum(dS, weekNumbers[i])
		overtime := calculateOvertime(weeklyHours, weekSum)
		fmt.Printf("Weekly sum: %s\n", weekSum)
		fmt.Printf("Weekly overtime: %s\n\n", overtime)
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

func readParseSort(filename string) []Entry {
	csvContent := csvReader(filename)
	data := extractCSVData(csvContent)
	// sort by calendar week
	sort.Slice(data, func(i, j int) bool {
		return data[i].kw < data[j].kw
	})
	return data
}

func timeDiff(start string, end string, format string) time.Duration {
	s, err := time.Parse(format, start)
	if err != nil {
		log.Print(err)
	}
	e, err := time.Parse(format, end)
	if err != nil {
		log.Print(err)
	}
	diff := e.Sub(s)
	return diff
}

func timeParse(timestring string, format string) time.Time {
	time, err := time.Parse(format, timestring) //-> time.Time
	if err != nil {
		log.Print(err)
	}
	return time
}

func extractCSVData(input [][]string) []Entry {
	var tmpEntry Entry
	var allEntries []Entry
	for _, entry := range input {
		tmpEntry.date = timeParse(entry[0], datefmt_ddmmyyyy)
		_, kw := tmpEntry.date.ISOWeek() // returns year and week as int
		tmpEntry.kw = kw
		tmpEntry.start = timeParse(entry[1], timefmt_hhmm)
		tmpEntry.end = timeParse(entry[2], timefmt_hhmm)
		tmpEntry.duration = timeDiff(entry[1], entry[2], timefmt_hhmm)
		tmpEntry.job = entry[3]
		allEntries = append(allEntries, tmpEntry)
	}
	return allEntries
}

func csvReader(fileName string) [][]string {
	// file open
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// csv reader init
	r := csv.NewReader(file)
	r.Comma = csvDelim
	r.Comment = csvComment
	// csv parsing
	content, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return content
}
