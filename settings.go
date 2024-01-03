package main

import (
	"time"
)

type Entry struct {
	date     time.Time
	kw       int
	start    time.Time
	end      time.Time
	duration time.Duration
	job      string
}

type Settings struct {
	csvDelim    rune
	datefmt     string // dd.mm.yyyy
	timefmt     string // hhmm
	weeklyHours time.Duration
}

const csvComment = '#' // cannot be changed by the user

const default_csvDelim rune = ' '
const default_datefmt_ddmmyyyy string = "02.01.2006" // dd.mm.yyyy
const default_timefmt_hhmm string = "1504"           // hhmm
const default_timefmt_hhmmss string = "150405"       // hhmmss
const default_hours time.Duration = 30               // hours
const default_weeklyHours time.Duration = time.Duration(default_hours * time.Hour)

var settings Settings

func initSettings() {
	settings.csvDelim = default_csvDelim
	settings.datefmt = default_datefmt_ddmmyyyy
	settings.timefmt = default_timefmt_hhmm
	settings.weeklyHours = default_weeklyHours
}
