package main

import (
	"github.com/derit/crontab" 
	"fmt"
)

type aLogger struct {
}

func (l *aLogger) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func main() {
	crontab.NewCronTab(crontab.CT_Second, crontab.SetLogger(&aLogger{})).
		SetSecond(1).
		Run(func(args ...interface{}) {
			fmt.Println("Task Running")
		}, "hello", "world")
}
