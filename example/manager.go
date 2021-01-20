package main

import (
	"fmt"
	"github.com/derit/crontab" 
	"log" 
)

func main() {
	var job = crontab.NewTaskManager()

	cron := crontab.NewCronTab(crontab.CT_Second).SetSecond(1)
	cron2 := crontab.NewCronTab(crontab.CT_Second).SetSecond(1)
	job.Add("xxx", cron, teststr)
	job.Add("xxx222", cron2, teststrs)

	log.Println("start...")
	job.Start()
	job.Wait()
}

func teststr(args ...interface{}) {
	fmt.Println("task test 1")
}
func teststrs(args ...interface{}) {
	fmt.Println("task test 2")
}