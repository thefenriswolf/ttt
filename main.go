package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const csvDelim = ' '
const csvComment = '#'
const datefmt = "02.01.2006" // dd.mm.yyyy
const timefmt = "1504"       // hhmm
const timefmtfull = "150405" // hhmmss

type Entry struct {
	date     string
	start    string
	end      string
	duration string
	job      string
}

func main() {
	fn := os.Args[1]
	csvContent := csvReader(fn)
	data := csvData(csvContent)

	for i := range data {
		fmt.Println(data[i].duration)
	}
}

func diff(start string, end string, format string) string {
	s, err := time.Parse(format, start)
	if err != nil {
		fmt.Println(err)
	}
	e, err := time.Parse(format, end)
	if err != nil {
		fmt.Println(err)
	}
	diff := e.Sub(s)
	return diff.String()
}

func csvData(input [][]string) []Entry {
	var tmpEntry Entry
	var allEntries []Entry
	for _, entry := range input {
		tmpEntry.date = entry[0]
		tmpEntry.start = entry[1]
		tmpEntry.end = entry[2]
		tmpEntry.duration = diff(entry[1], entry[2], timefmt)
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
