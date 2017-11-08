package main

import (
	"encoding/json"
	"fmt"
)

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
	verify(NotFound, "NotFound", "User could not be found", "{\"type\":\"NotFound\",\"message\":\"User could not be found\"}")
	verify(AlreadyExists, "AlreadyExists", "User already exists", "{\"type\":\"AlreadyExists\",\"message\":\"User already exists\"}")
	verify(NotSure, "NotSure", "Not sure what happened", "{\"type\":\"NotSure\",\"message\":\"Not sure what happened\"}")
	verify(BadRequestData, "BadRequestData", "You didn't send a good request", "{\"type\":\"BadRequestData\",\"message\":\"You didn't send a good request\"}")
	verify(WorksOnMyMachine, "WorksOnMyMachine", "Works on my machine", "{\"type\":\"WorksOnMyMachine\",\"message\":\"Works on my machine\"}")
}

func toString(val interface{}) string {
	stringer, _ := val.(fmt.Stringer)
	return stringer.String()
}

func toJSON(val interface{}) []byte {
	encoded, _ := json.Marshal(val)

	return encoded
}

func verify(err Error, name string, message string, response string) {
	if fmt.Sprint(err) != message {
		panic("Didn't get the write value")
	}

	if toString(err) != name {
		panic("Wrong name")
	}

	if string(toJSON(err)) != response {
		panic("Wrong JSON")
	}
}
