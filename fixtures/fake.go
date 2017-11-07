package main

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
	verify(NotFound, "NotFound", 1, "User could not be found")
	verify(AlreadyExists, "AlreadyExists", 2, "User already exists")
	verify(NotSure, "NotSure", 3, "Not sure what happened")
	verify(BadRequestData, "BadRequestData", 4, "You didn't send a good request")
	verify(WorksOnMyMachine, "WorksOnMyMachine", 5, "Works on my machine")
}

func verify(err Error, name string, code int, message string) {
	errStruct, ok := Error.Error().(TestError)

	if !ok {
		panic("Couldn't cast")
	}

	if errStruct.Name != name {
		panic("wrong name")
	}

	if errStruct.Message != message {
		panic("wrong message")
	}
}
