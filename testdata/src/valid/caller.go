package valid

import (
	"fmt"
)

const (
	statusConst Status = 32818
)

var statusVar Status = 191949

func Run() {
	valid_main()
}

func valid_main() {
	status := Done
	run(status)
	run(statusConst)
	run(statusVar)
}

func run(s Status) {
	fmt.Println(s)
}
