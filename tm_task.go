package crontab

import "context"

type TaskObject struct {
	cron     *CronTab
	callback HandleFunc
	args     []interface{}
	cancel   context.CancelFunc
	title    string
	taskId   string
}

func (to *TaskObject) start() {
	if to.cron.running == true {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	to.cron.ctx = ctx
	to.cancel = cancel
	to.cron.Run(to.callback, to.args...)
}
func (to *TaskObject) stop() {
	if to.cron.running == false {
		return
	}
	(to.cancel)()
}
func (to *TaskObject) IsRunning() bool {
	return to.cron.IsRunning()
}
func (to *TaskObject) Title() string {
	return to.title
}
