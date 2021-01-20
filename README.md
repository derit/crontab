# crontab

a simple and powerful crontab written in golang

### Simple Cron

```go
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
```


### Custom Logger

```go
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
```
 

### Task Manager

```go
package main

import (
	"fmt"
	"github.com/derit/crontab" 
	"log"
	"time"
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
```
