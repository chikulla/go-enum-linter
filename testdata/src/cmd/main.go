package main

import (
	"testdata/invalid"
	"testdata/valid"
)

func main() {
	valid.Run()
	invalid.Run()
}
