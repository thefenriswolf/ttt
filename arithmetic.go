package main

import (
	"fmt"
	"reflect"
	"time"
)

func durationSum(data []Entry) time.Duration {
	var s time.Duration = 0
	for i := range data {
		s += data[i].duration
	}
	fmt.Println("type of s: ", reflect.TypeOf(s))
	fmt.Println("duration sum: ", s)
	return s
}
