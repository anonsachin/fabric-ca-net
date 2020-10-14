package main

import (
	"fmt"
)

type exec interface{
	Run() error
}

func execute(cmd exec) error{
	defer recov()
	err  := cmd.Run()
	return err
}


func recov() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}