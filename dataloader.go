package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

const csvComment = '#'

var csvDelim rune = ' '
var datefmt_ddmmyyyy string = "02.01.2006" // dd.mm.yyyy
var timefmt_hhmm string = "1504"           // hhmm
var timefmt_hhmmss string = "150405"       // hhmmss
var hours time.Duration = 30               // hours
var weeklyHours time.Duration = time.Duration(hours * time.Hour)

type Entry struct {
	date     time.Time
	kw       int
	start    time.Time
	end      time.Time
	duration time.Duration
	job      string
}

func readParseSort(filename string) []Entry {
	settingsReader(filename)
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
	// #todo: this assumes:
	// date, start time, end time, job
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

//  hours: 30
// delimiter: ' '
// datefmt: "02.01.2006"
// timefmt: "1504"

func settingsParser(settings []string) {
	for _, line := range settings {
		if strings.Contains(line, "hours") {
			fmt.Println("not implemented yet: ", line)
		} else if strings.Contains(line, "delimiter") {
			fmt.Println("not implemented yet: ", line)
		} else if strings.Contains(line, "datefmt") {
			fmt.Println("not implemented yet: ", line)
		} else if strings.Contains(line, "timefmt") {
			fmt.Println("not implemented yet: ", line)
		} else {
			continue
		}
	}
}

func settingsReader(fileName string) {
	// file open
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var settings []string
	for scanner.Scan() {
		// omit empty lines
		if scanner.Text() == "" {
			continue
		} else if scanner.Text()[0] != csvComment {
			continue
		} else {
			settings = append(settings, scanner.Text())
		}
	}
	// if no settings were applied we use the defaults
	if len(settings) == 0 {
		fmt.Printf("\nINFO No configuration found:\nUsing default settings\n\n")
		csvDelim = ' '
		datefmt_ddmmyyyy = "02.01.2006"
		timefmt_hhmm = "1504"
		timefmt_hhmmss = "150405"
		hours = 30
	} else {
		settingsParser(settings)
	}
}

func csvReader(fileName string) [][]string {
	// file open
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// csv reader init
	reader := csv.NewReader(file)
	// csv reader settings
	reader.TrimLeadingSpace = true
	reader.Comma = csvDelim
	reader.Comment = csvComment
	// csv parsing
	// read entire file is fine because we deal with sub MB file sizes
	content, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return content
}
