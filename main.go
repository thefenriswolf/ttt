package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fn, _ := argParse(os.Args[1:]) // #todo: swap if actions are implemented
	data := readAndParse(fn)
	k := data[0].duration
	fmt.Println("type of duration: ", reflect.TypeOf(k))
	_ = durationSum(data)

	fmt.Println("parsed file entries:")
	for i := range data {
		fmt.Println(data[i].date.Format(datefmt_ddmmyyyy), data[i].kw, data[i].duration)
	}
}
