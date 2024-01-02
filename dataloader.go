package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readParseSort(filename string, s Settings) []Entry {
	settingsReader(filename)
	csvContent := csvReader(filename, s)
	data := extractCSVData(csvContent, s)
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

func extractCSVData(input [][]string, s Settings) []Entry {
	var tmpEntry Entry
	var allEntries []Entry
	// #todo: this assumes:
	// date, start time, end time, job
	for _, entry := range input {
		tmpEntry.date = timeParse(entry[0], s.datefmt)
		_, kw := tmpEntry.date.ISOWeek() // returns year and week as int
		tmpEntry.kw = kw
		tmpEntry.start = timeParse(entry[1], s.timefmt)
		tmpEntry.end = timeParse(entry[2], s.timefmt)
		tmpEntry.duration = timeDiff(entry[1], entry[2], s.timefmt)
		tmpEntry.job = entry[3]
		allEntries = append(allEntries, tmpEntry)
	}
	return allEntries
}

//  hours: 30
// delimiter: ' '
// datefmt: "02.01.2006"
// timefmt: "1504"

func settingsParser(set []string) {
	const sep string = ":"
	for _, line := range set {
		if strings.Contains(line, "hours") {
			Hours := hoursParser(line, sep)
			fmt.Printf("hours: %d\n", Hours)
			settings.weeklyHours = time.Duration(Hours * time.Hour)
		} else if strings.Contains(line, "delimiter") {
			delimiter := delimiterParser(line, sep)
			fmt.Printf("delimiter: %s\n", string(delimiter))
			settings.csvDelim = delimiter
		} else if strings.Contains(line, "datefmt") {
			dateformat := dateTimeFormat(line, sep)
			fmt.Printf("date format: %s\n", dateformat)
			settings.datefmt = dateformat
		} else if strings.Contains(line, "timefmt") {
			timeformat := dateTimeFormat(line, sep)
			fmt.Printf("time format: %s\n\n", timeformat)
			settings.timefmt = timeformat
		} else {
			continue
		}
	}
}

func dateTimeFormat(line string, sep string) string {
	_, value, valid := strings.Cut(line, sep)
	if valid == false {
		log.Fatal("not a valid config: \n", line)
	}
	dateTime := strings.TrimSpace(value)
	return dateTime
}

func delimiterParser(line string, sep string) rune {
	_, value, valid := strings.Cut(line, sep)
	if valid == false {
		log.Fatal("not a valid config: \n", line)
	}
	value = strings.TrimSpace(value)
	if value == "" {
		val := []rune(" ")
		return val[0]
	} else {
		val := []rune(value)
		return val[0]
	}
}

func hoursParser(line string, sep string) time.Duration {
	_, value, valid := strings.Cut(line, sep)
	if valid == false {
		log.Fatal("not a valid config: \n", line)
	}
	value = strings.TrimSpace(value)
	hours, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return time.Duration(hours)
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
	} else {
		settingsParser(settings)
	}
}

func csvReader(fileName string, s Settings) [][]string {
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
	reader.Comma = s.csvDelim
	reader.Comment = csvComment
	// csv parsing
	// read entire file is fine because we deal with sub MB file sizes
	content, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return content
}
