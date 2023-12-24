package main

import (
	"fmt"
	"os"
)

func main() {
	fn, _ := argParse(os.Args[1:]) // #todo: swap if actions are implemented
	data := readAndParse(fn)
	var kwEntry []Entry
	for i, _ := range data {
		if data[i].kw == 52 {
			kwEntry = append(kwEntry, data[i])
		}
	}
	_ = durationSum(data)
	_ = durationSum(kwEntry)

	fmt.Println("parsed file entries:")
	for i := range data {
		fmt.Println(data[i].date.Format(datefmt_ddmmyyyy), data[i].kw, data[i].duration)
	}
}
