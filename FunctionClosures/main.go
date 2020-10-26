package main

import (
	"fmt"
	"time"
)

// Tasks is the data collection format
type Tasks struct {
	Name     string
	ExecFunc func()
}

// CallFunc is one of the method for the Tasks struct
func (t *Tasks) CallFunc(f func()) *Tasks {
	t.ExecFunc()
	return t
}

func runHello() func() {
	return func() {
		fmt.Println("Hello Graviton")
	}
}

func runProcess(k *Tasks) func() {
	return func() {
		taskName := k.Name
		doProcess(taskName)
	}
}

func doProcess(taskName string) {
	fmt.Println("task name from runProcess: ", taskName)
}

func main() {
	// Initialize the tasks struck
	t := Tasks{
		Name:     "Task Name 1",
		ExecFunc: runHello(),
	}

	fmt.Println("executing task name: ", t.Name)
	go t.CallFunc(t.ExecFunc)
	time.Sleep(1 * time.Second)

	fmt.Println("main func ended here")

	// Another common example
	t = Tasks{
		Name:     "Task Name 2",
		ExecFunc: runProcess(&t),
	}
	go t.CallFunc((t.ExecFunc))
	time.Sleep(1 * time.Second)
}
