package main

import (
	"fmt"
	"github.com/CzaOrz/AGScheduler"
	"github.com/CzaOrz/AGScheduler/schedulers"
	"github.com/CzaOrz/AGScheduler/stores"
	"github.com/CzaOrz/AGScheduler/tasks"
	"github.com/CzaOrz/AGScheduler/triggers"
	"os"
	"time"
)

func main() {
	now := time.Now()
	scheduler := schedulers.NewScheduler(AGScheduler.WorksMap{}, stores.NewMemoryStore())

	trigger1, _ := triggers.NewIntervalTrigger(now.Add(time.Second*1), AGScheduler.EmptyDateTime, time.Second*6)
	task1 := tasks.NewTask("task1", func(args []interface{}) {
		fmt.Println(args, time.Now())
	}, []interface{}{"this", "is", "task1"}, trigger1)
	_ = scheduler.AddTask(task1)

	go func() {
		time.Sleep(time.Second * 20)
		task, _ := scheduler.GetTask("task1")

		trigger2, _ := triggers.NewIntervalTrigger(now, AGScheduler.EmptyDateTime, time.Second*2)
		task.UpdateTrigger(trigger2)

		time.Sleep(time.Second * 30)
		os.Exit(1)
	}()

	scheduler.Start()
}
