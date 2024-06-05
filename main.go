package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type config struct {
	hours uint64
	job   string
}

type entry struct {
	date      time.Time
	startTime time.Time
	endTime   time.Time
}

func checkNPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(f *os.File) (config []string, data []string) {
	reader := bufio.NewReader(f)
	var lines []string
	var conf []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			checkNPanic(err)
		}
		if line[0] != '#' {
			lines = append(lines, line)
		} else {
			conf = append(conf, line)
		}
	}
	return conf, lines
}

func confParser(conf []string) config {
	var config config
	for _, element := range conf {
		if strings.Contains(element, "hours") {
			h := strings.SplitAfter(element, ": ")
			hr := strings.Trim(h[1], "\n")
			hours, err := strconv.ParseUint(hr, 10, 64)
			if err != nil {
				log.Println(err)
			}
			config.hours = hours
		}
		if strings.Contains(element, "job") {
			j := strings.SplitAfter(element, ": ")
			job := strings.Trim(j[1], "\n")
			config.job = job
		}
	}
	return config
}

// const default_datefmt_ddmmyyyy string = "02.01.2006" // dd.mm.yyyy
// const default_timefmt_hhmm string = "1504"

func linesParser(lines []string) []entry {
	var data []entry
	const dateFormat = "02.01.2006"
	const timeFormat = "1504"
	const tmpDateFmt = "2006 January 1 " + timeFormat

	for _, line := range lines {
		if len(line) > 1 && !strings.Contains(line, "#") {
			var linedata entry
			fields := strings.Fields(line)

			date := fields[0]
			start := fields[1]
			end := fields[2]

			//fmt.Printf("date %v, start %v, end %v\n", date, start, end)
			t, _ := time.Parse(dateFormat, date)

			year, month, day := t.Date()
			//fmt.Printf("year %v, month %v, day %v\n", year, month, day)

			startTime := fmt.Sprintf("%i %s %i %i", year, month, day, start)
			endTime := fmt.Sprintf("%i %s %i %i", year, month, day, end)

			s, err := time.Parse(tmpDateFmt, string(startTime))
			if err != nil {
				log.Fatal(err)
			}
			e, err := time.Parse(tmpDateFmt, string(endTime))
			if err != nil {
				log.Fatal(err)
			}

			linedata.date = t
			linedata.startTime = s
			linedata.endTime = e

			fmt.Printf("date %v, start %v, end %v\n", linedata.date, linedata.startTime, linedata.endTime)
			// go func() {
			//startTime := strings.Trim(dps[1], " ")
			//fmt.Printf("start time: %s\n", startTime)
			// 	wg.Done()
			// }()
		}
	}
	return data
}

func main() {
	var filePath = flag.String("f", "", "ttt database file to use")
	// var pprint = flag.Bool("print", false, "pretty print database file")
	flag.Parse()

	f, err := os.Open(*filePath)
	checkNPanic(err)
	defer f.Close()

	conf, lines := readLines(f)

	_ = confParser(conf)
	_ = linesParser(lines)
	// fmt.Println(config)
	// fmt.Println(data)
}
