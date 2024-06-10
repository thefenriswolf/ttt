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

func linesParser(lines []string) []entry {
	var data []entry
	const dateFormat = "02.01.2006"
	const timeFormat = "1504"
	const tmpDateFmt = "2006 1 2 " + timeFormat

	for i, line := range lines {
		fmt.Printf("\033[2K\rProcessing entry: %d", i+1)
		if len(line) > 1 && !strings.Contains(line, "#") {
			var linedata entry
			fields := strings.Fields(line)

			date := fields[0]
			start := fields[1]
			end := fields[2]

			t, err := time.Parse(dateFormat, date)
			if err != nil {
				log.Fatal(err)
			}

			year, month, day := t.Date()

			startTime := fmt.Sprintf("%v %v %v %v", year, int(month), day, start)
			endTime := fmt.Sprintf("%v %v %v %v", year, int(month), day, end)

			st := make(chan time.Time)
			go parseTime(st, tmpDateFmt, startTime)

			et := make(chan time.Time)
			go parseTime(et, tmpDateFmt, endTime)

			linedata.date = t
			linedata.startTime = <-st
			linedata.endTime = <-et

			data = append(data, linedata)
		}
	}
	return data
}

func parseTime(ch chan time.Time, fmt string, ts string) {
	s, err := time.Parse(fmt, ts)
	if err != nil {
		log.Fatal(err)
	}
	ch <- s
	close(ch)
}

func main() {
	var filePath = flag.String("f", "", "ttt database file to use")
	// var pprint = flag.Bool("print", false, "pretty print database file")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	f, err := os.Open(*filePath)
	checkNPanic(err)
	defer f.Close()

	conf, lines := readLines(f)

	_ = confParser(conf)
	data := linesParser(lines)

	var buf []string
	for i, n := range data {
		tmp := fmt.Sprintf("Date: %s\nStart: %s -- End: %s\nDiff: %s\n\n", n.date, n.startTime, n.endTime, n.endTime.Sub(n.startTime))
		buf = append(buf, tmp)
		fmt.Printf("\033[2K\rBuffering output: %d", i+1)
		//	fmt.Printf("Date: %s\nStart: %s -- End: %s\nDiff: %s\n\n", n.date, n.startTime, n.endTime, n.endTime.Sub(n.startTime))

	}
	fmt.Println(buf)
}
