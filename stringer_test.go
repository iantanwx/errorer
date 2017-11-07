package errorer

import (
	"strings"
	"testing"
)

// Simple, no offset, single-run enum of errors
var basic_in = `type Error int
const (
	NotFound         Error = iota //User could not be found
	AlreadyExists                 //User already exists
	NotSure                       //Not sure what happened
	BadRequestData                //You didn't send a good request
	WorksOnMyMachine              //Works on my machine
)
`

var basic_out = `
const _Error_name = "NotFoundAlreadyExistsNotSureBadRequestDataWorksOnMyMachine"

var _Error_index = [...]uint8{0, 8, 21, 28, 42, 58}

func (i Error) String() string {
	if i < 0 || i >= Error(len(_Error_index)-1) {
		return fmt.Sprintf("Error(%d)", i)
	}
	return _Error_name[_Error_index[i]:_Error_index[i+1]]
}

const _Error_msg = "User could not be foundUser already existsNot sure what happenedYou didn't send a good requestWorks on my machine"

var _Error_msg_index = [...]uint8{0, 23, 42, 64, 94, 113}

func (i Error) Error() string {
	if i < 0 || i >= Error(len(_Error_index)-1) {
		return fmt.Sprintf("Error(%d)", i)
	}
	return _Error_msg[_Error_msg_index[i]:_Error_msg_index[i+1]]
}

func (i Error) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	msg, err := json.Marshal(i.Error())
	if err != nil {
		return b, err
	}
	name, err := json.Marshal(i.String())
	if err != nil {
		return b, err
	}
	json := fmt.Sprintf("{\"type\":%s,\"message\":%s}", name, msg)
	b.WriteString(json)
	return b.Bytes(), nil
}
`

type Golden struct {
	name   string
	input  string
	output string
}

var golden = []Golden{
	{"basic", basic_in, basic_out},
}

func TestGolden(t *testing.T) {
	for _, test := range golden {
		var g Generator
		in := "package test\n" + test.input

		file := test.name + ".go"

		g.parsePackage(".", []string{file}, in)

		tokens := strings.SplitN(test.input, " ", 3)

		if len(tokens) != 3 {
			t.Fatalf("Need type declaration on first line")
		}

		g.Generate(tokens[1])

		out := string(g.Format())

		if out != test.output {
			t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, out, test.output)
		}
	}
}
