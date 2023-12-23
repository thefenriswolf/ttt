package main

import (
// "fmt"
)

func argParse(args []string) (string, string) {
	action := args[0]
	file := "nil"
	//file := args[1] // todo: swap if actions are implemented
	return action, file
}
