package main

import "fmt"

type Error int

const (
	NotFound         Error = iota // User could not be found
	AlreadyExists                 // User already exists
	NotSure                       // Not sure what happened
	BadRequestData                // You didn't send a good request
	WorksOnMyMachine              // Works on my machine
)

type TestError struct {
	Name    string
	Message string
}

func main() {
	verify(NotFound, "User could not be found")
	verify(AlreadyExists, "User already exists")
	verify(NotSure, "Not sure what happened")
	verify(BadRequestData, "You didn't send a good request")
	verify(WorksOnMyMachine, "Works on my machine")
}

func verify(err Error, message string) {
	if fmt.Sprint(err) != message {
		panic("Didn't get the write value")
	}
}
