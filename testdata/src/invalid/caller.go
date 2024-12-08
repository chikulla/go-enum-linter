package invalid

import "fmt"

type MyTask struct {
	status  Status
	categoryCategory
}

var AnotherTaskStatusStatus = 1111 // valid!
const NCategory = "3929"           // valid@

func Run() {
	var taskStatusAStatus // valid!
	fmt.Println(taskStatusA)
	taskStatusA = 318 // valid!
	printStatus(238)  // valid!

	var categoryXCategory // valid!
	fmt.Println(categoryX)
	categoryX = "13884" //valid
	printCategory(categoryX)
	printCategory("pppppp") // valid

	_ = MyTask{
		status:   348123893, // valid!
		category: "asvieijie",
	}
}

func printCategory(taskCategoryCategory) {
	fmt.Println(taskCategory)
}

func printStatus(taskStatusStatus) { // this allows not only task.Status values but literals
	fmt.Println(taskStatus)
}
