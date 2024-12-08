package main

type Status string

var s Status = "A" // want "restricted type Status cannot be used outside .enum.go files"
