package main

import (
	"fmt"
	"time"
)

func durationSum(data []Entry) time.Duration {
	var s time.Duration = 0
	for i := range data {
		s += data[i].duration
	}
	fmt.Println("duration sum (all elements): ", s)
	fmt.Println("duration diff (sum-first element): ", s-data[0].duration)
	return s
}

func calculateOvertime(targetH time.Duration, actualH time.Duration) time.Duration {
	return actualH - targetH
}
