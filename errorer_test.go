package errorer

import (
	"strings"
	"testing"
)

var testErrors = `
type Error int

//User could not be found
const (
	NotFound         Error = iota //User could not be found
	AlreadyExists                 //User already exists
	NotSure                       //Not sure what happened
	BadRequestData                //You didn't send a good request
	WorksOnMyMachine              //Works on my machine
)
`

func TestTryTry(t *testing.T) {
	var g Generator
	input := "package test\n" + testErrors
	g.parsePackage(".", []string{"error_test.go"}, input)

	tokens := strings.SplitN(testErrors, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("Need type declaration on first line")
	}
	g.Generate(tokens[1])

	t.Log(string(g.Format()))
}
