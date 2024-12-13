package valid

import "fmt"

type task struct {
	status   Status
	category Category
}

var AnotherStatus Status = 1111 // want "some"

const AnotherCategory Category = "3929" // want "some"

func Run() {
	var status Status // want "some"
	fmt.Println(status)
	printStatus(238) // want "some"

	var category Category // want "some"
	fmt.Println(category)
	printCategory("pppppp") // want "some"

	_ = task{
		status:   348123893,   // want "some"
		category: "asvieijie", // want "some"
	}
}

func printCategory(task Category) {
	fmt.Println(task)
}

func printStatus(status Status) {
	fmt.Println(status)
}
