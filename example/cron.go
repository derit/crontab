package main

import (
	"github.com/derit/crontab"
	"log"
)

func main() {
	crontab.NewCronTab(crontab.CT_Second).
		SetSecond(1).
		Run(func(args ...interface{}) {
			log.Println("Task Running")
		})
}