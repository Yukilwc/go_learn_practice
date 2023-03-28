package timer

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Timer interface {
	AddTaskByFunc(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error)
}

type timer struct {
	taskList map[string]*cron.Cron
	sync.Mutex
}

func (t *timer) AddTaskByFunc(taskName string, spec string, task func(), option ...cron.Option) (cron.EntryID, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New(option...)
	}
	id, err := t.taskList[taskName].AddFunc(spec, task)
	t.taskList[taskName].Start()
	return id, err
}

func NewTimerTask() Timer {
	return &timer{taskList: make(map[string]*cron.Cron)}
}
