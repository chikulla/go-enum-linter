package valid

type Status int

const (
	Done       Status = iota // OK
	InProgress               // OK
	ToDo                     // OK
)
