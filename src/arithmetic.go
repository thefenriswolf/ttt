package ttt

import (
	"github.com/mpvl/unique"
	"time"
)

func durationSum(data []Entry) time.Duration {
	// calulate over all durations in array
	var s time.Duration = 0
	for i := range data {
		s += data[i].duration
	}
	return s
}

func monthList(data []Entry) []int {
	var months []int
	for i := range data {
		dateString := data[i].date
		months = append(months, int(dateString.Month()))
	}
	// only keep unique months
	unique.Ints(&months)
	return months
}

func weekList(data []Entry) []int {
	// get all week numbers contained in list of structs "data"
	var kWeeks []int
	for i := range data {
		kWeeks = append(kWeeks, data[i].kw)
	}
	// only keep unique week numbers
	unique.Ints(&kWeeks)
	return kWeeks
}

func calculateOvertime(targetHours time.Duration, actualHours time.Duration) time.Duration {
	// overtime can be negative or positive
	return actualHours - targetHours
}

func monthSum(data []Entry, monthNum int) time.Duration {
	var month []Entry
	// create array of all entry corresponding to the given month
	for i := range data {
		monthN := int((data[i].date).Month())
		if monthN == monthNum {
			month = append(month, data[i])
		}
	}
	// calculate sum over all days in that week
	monthSum := durationSum(month)
	return monthSum
}

func weekSum(data []Entry, weekNumber int) time.Duration {
	var week []Entry
	// create array of all entry corresponding to the given week
	for i := range data {
		if data[i].kw == weekNumber {
			week = append(week, data[i])
		}
	}
	// calculate sum over all days in that week
	weekSum := durationSum(week)
	return weekSum
}

func lastWeekSum(data []Entry, weeks []int) time.Duration {
	var lastWeek []Entry
	// create array of all entry corresponding to the most recent week
	for i := range data {
		if data[i].kw == weeks[len(weeks)-1] {
			lastWeek = append(lastWeek, data[i])
		}
	}
	// calculate sum over all days in that week
	lastWeekSum := durationSum(lastWeek)
	return lastWeekSum
}
