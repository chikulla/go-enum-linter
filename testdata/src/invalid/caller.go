package invalid

import "fmt"

type task struct {
	status   Status
	category Category
}

var AnotherStatus Status = 1111         // OK (Status not defined at *.enum.go)
const AnotherCategory Category = "3929" // OK (Category not defined at *.enum.go)

func Run() {
	var status Status // OK
	fmt.Println(status)
	status = 318     // OK
	printStatus(238) // OK

	var category Category // OK
	fmt.Println(category)
	category = "13884" // OK
	printCategory(category)
	printCategory("pppppp") // OK

	_ = task{
		status:   348123893,   // OK
		category: "asvieijie", // OK
	}
}

func printCategory(task Category) {
	fmt.Println(task)
}

func printStatus(status Status) {
	fmt.Println(status)
}
