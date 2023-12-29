package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func writeTemplate() {
	dateTime := time.Now()
	cDate := dateTime.Format("20060102")
	const postfix string = "_ttt.csv"
	filename := cDate + postfix

	lines := []string{
		"# ttt journal file",
		"#",
		"# hours: 40",
		"# delimiter: ",
		"# datefmt: 02.01.2006",
		"# timefmt: 1504",
		"#",
		"01.01.2024 0900 1700 Some Job"}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Template created ...\nHappy time tracking!\n")
}
