package invalid

import "fmt"

type task struct {
	status   Status
	category Category
}

var AnotherStatus Status = 1111         // valid!
const AnotherCategory Category = "3929" // valid!

func Run() {
	var status Status // valid!
	fmt.Println(status)
	status = 318     // valid!
	printStatus(238) // valid!

	var category Category // valid!
	fmt.Println(category)
	category = "13884" //valid
	printCategory(category)
	printCategory("pppppp") // valid

	_ = task{
		status:   348123893, // valid!
		category: "asvieijie",
	}
}

func printCategory(task Category) {
	fmt.Println(task)
}

func printStatus(status Status) { // this allows not only task.Status values but literals
	fmt.Println(status)
}
