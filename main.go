package main

import (
	"fmt"
	"os"
	//"reflect"
)

func main() {
	fn, _ := argParse(os.Args[1:]) // #todo: swap if actions are implemented
	data := readAndParse(fn)
	//	fmt.Printf("%+v\n", data[0:])
	for i := range data {
		fmt.Println(data[i].date.Format(datefmt_ddmmyyyy), data[i].kw, data[i].duration)
	}
}
