package main

import (
	"fmt"

	"github.com/chikulla/go-enum-linter/testdata/src/valid/task"
)

type MyTask struct {
	status task.Status
}

const (
	AnotherTaskStatus task.Status = 1111 // valid!
)

func main() {
	var taskStatusA task.Status // valid!
	fmt.Println(taskStatusA)
	taskStatusA = 238 // valid!
	printStatus(238)  // valid!

	_ = MyTask{
		status: 348123893, // valid!
	}
}

func printStatus(taskStatus task.Status) { // this allows not only task.Status values but literals
	fmt.Println(taskStatus)
}
