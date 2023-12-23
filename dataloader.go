package main

import (
	"encoding/csv"
	"fmt"
	"os"
	//"reflect"
	"sort"
	"time"
)

const csvDelim = ' '
const csvComment = '#'
const datefmt_ddmmyyyy = "02.01.2006" // dd.mm.yyyy
const timefmt_hhmm = "1504"           // hhmm
const timefmt_hhmmss = "150405"       // hhmmss

type Entry struct {
	date     time.Time
	kw       int
	start    time.Time
	end      time.Time
	duration time.Duration
	job      string
}

func readAndParse(filename string) []Entry {
	csvContent := csvReader(filename)
	data := csvData(csvContent)
	sort.Slice(data, func(i, j int) bool {
		return data[i].kw < data[j].kw
	})
	return data
}

func timeDiff(start string, end string, format string) time.Duration {
	s, err := time.Parse(format, start)
	if err != nil {
		fmt.Println(err)
	}
	e, err := time.Parse(format, end)
	if err != nil {
		fmt.Println(err)
	}
	diff := e.Sub(s)
	return diff
}

func timeParse(timestring string, format string) time.Time {
	time, err := time.Parse(format, timestring) //-> time.Time
	if err != nil {
		fmt.Println(err)
	}
	return time
}

func csvData(input [][]string) []Entry {
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
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = csvDelim
	r.Comment = csvComment

	content, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return content
}
