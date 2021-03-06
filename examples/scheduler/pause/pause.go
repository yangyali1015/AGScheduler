package main

import (
	"fmt"
	"github.com/CzaOrz/AGScheduler"
	"github.com/CzaOrz/AGScheduler/schedulers"
	"github.com/CzaOrz/AGScheduler/stores"
	"github.com/CzaOrz/AGScheduler/tasks"
	"github.com/CzaOrz/AGScheduler/triggers"
	"time"
)

func main() {
	now := time.Now()
	scheduler := schedulers.NewScheduler(AGScheduler.WorksMap{}, stores.NewMemoryStore())

	trigger1, _ := triggers.NewIntervalTrigger(now.Add(time.Second*1), AGScheduler.EmptyDateTime, time.Second*5)
	task1 := tasks.NewTask("task1", func(args []interface{}) {
		fmt.Println(args)
	}, []interface{}{"this", "is", "task1"}, trigger1)
	_ = scheduler.AddTask(task1)

	go func() {
		time.Sleep(time.Second * 10)
		task1.Pause()
		fmt.Println("Pause", time.Now())
		time.Sleep(time.Second * 20)
		fmt.Println("Resume", time.Now())
		task1.Resume()
	}()

	scheduler.Start()
}
