package main

import "fmt"

func Print(args ...interface{}) {
	if len(args) != 1 {
		return
	}
	if v, ok := args[0].(map[string]string); ok {
		fmt.Println(v)
	}
}
