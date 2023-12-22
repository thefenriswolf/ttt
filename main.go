package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Entry struct {
	date  string
	start string
	end   string
	job   string
}

func main() {
	fn := os.Args[1]
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ' '
	r.Comment = '#'

	content, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	var tmpEntry Entry
	var allEntries []Entry
	for _, entry := range content {
		tmpEntry.date = entry[0]
		tmpEntry.start = entry[1]
		tmpEntry.end = entry[2]
		tmpEntry.job = entry[3]
		allEntries = append(allEntries, tmpEntry)
	}
	fmt.Println(allEntries[1].start)
}
