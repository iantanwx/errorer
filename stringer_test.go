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

const basic_out = `
const _Error_name = "NotFoundAlreadyExistsNotSureBadRequestDataWorksOnMyMachine"

var _Error_name_index = [...]uint8{0, 8, 21, 28, 42, 58}

func (i Error) String() string {
	if i < 0 || i >= Error(len(_Error_name_index)-1) {
		return fmt.Sprintf("Error(%d)", i)
	}
	return _Error_name[_Error_name_index[i]:_Error_name_index[i+1]]
}

const _Error_msg = "User could not be foundUser already existsNot sure what happenedYou didn't send a good requestWorks on my machine"

var _Error_msg_index = [...]uint8{0, 23, 42, 64, 94, 113}

func (i Error) Error() string {
	if i < 0 || i >= Error(len(_Error_msg_index)-1) {
		return fmt.Sprintf("Error(%d)", i)
	}
	return _Error_msg[_Error_msg_index[i]:_Error_msg_index[i+1]]
}

var _ErrorNameToValue_map = map[string]Error{
	_Error_name[0:8]:   0,
	_Error_name[8:21]:  1,
	_Error_name[21:28]: 2,
	_Error_name[28:42]: 3,
	_Error_name[42:58]: 4,
}

func ErrorString(s string) (Error, error) {
	if val, ok := _ErrorNameToValue_map[s]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("%s is not the name of type Error")
}

func (i Error) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	msg, err := json.Marshal(i.Error())
	if err != nil {
		return b.Bytes(), err
	}
	name, err := json.Marshal(i.String())
	if err != nil {
		return b.Bytes(), err
	}
	json := fmt.Sprintf("{\"type\":%s,\"message\":%s}", name, msg)
	b.WriteString(json)
	return b.Bytes(), nil
}

type errStruct struct {
	name    string
	message string
}

func (i *Error) UnmarshalJSON(data []byte) error {
	var errData errStruct

	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("Expecting a string, got %s", data)
	}

	val, err := ErrorString(errData.name)

	if err != nil {
		return err
	}

	*i = val

	return nil
}
`

const offset_in = `type Error int
const (
	NotFound         Error = iota + 100 //User could not be found
	AlreadyExists						//User already exists
	NotSure								//Not sure what happened
	BadRequestData						//You didn't send a good request
	WorksOnMyMachine					//Works on my machine
)
`

const offset_out = `
const _Error_name = "NotFoundAlreadyExistsNotSureBadRequestDataWorksOnMyMachine"

var _Error_name_index = [...]uint8{0, 8, 21, 28, 42, 58}

func (i Error) String() string {
	i -= 100
	if i < 0 || i >= Error(len(_Error_name_index)-1) {
		return fmt.Sprintf("Error(%d)", i+100)
	}
	return _Error_name[_Error_name_index[i]:_Error_name_index[i+1]]
}

const _Error_msg = "User could not be foundUser already existsNot sure what happenedYou didn't send a good requestWorks on my machine"

var _Error_msg_index = [...]uint8{0, 23, 42, 64, 94, 113}

func (i Error) Error() string {
	i -= 100
	if i < 0 || i >= Error(len(_Error_msg_index)-1) {
		return fmt.Sprintf("Error(%d)", i+100)
	}
	return _Error_msg[_Error_msg_index[i]:_Error_msg_index[i+1]]
}

var _ErrorNameToValue_map = map[string]Error{
	_Error_name[0:8]:   100,
	_Error_name[8:21]:  101,
	_Error_name[21:28]: 102,
	_Error_name[28:42]: 103,
	_Error_name[42:58]: 104,
}

func ErrorString(s string) (Error, error) {
	if val, ok := _ErrorNameToValue_map[s]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("%s is not the name of type Error")
}

func (i Error) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	msg, err := json.Marshal(i.Error())
	if err != nil {
		return b.Bytes(), err
	}
	name, err := json.Marshal(i.String())
	if err != nil {
		return b.Bytes(), err
	}
	json := fmt.Sprintf("{\"type\":%s,\"message\":%s}", name, msg)
	b.WriteString(json)
	return b.Bytes(), nil
}

type errStruct struct {
	name    string
	message string
}

func (i *Error) UnmarshalJSON(data []byte) error {
	var errData errStruct

	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("Expecting a string, got %s", data)
	}

	val, err := ErrorString(errData.name)

	if err != nil {
		return err
	}

	*i = val

	return nil
}
`

const multiple_in = `type Error int
const (
	NotFound         Error = 100		//User could not be found
	AlreadyExists	 Error = 101				//User already exists
	NotSure			 Error = 103					//Not sure what happened
	BadRequestData	Error = 104				//You didn't send a good request
	WorksOnMyMachine Error = 105					//Works on my machine
)
`

const multiple_out = `
const (
	_Error_name_0 = "NotFoundAlreadyExists"
	_Error_name_1 = "NotSureBadRequestDataWorksOnMyMachine"
)

var (
	_Error_name_index_0 = [...]uint8{0, 8, 21}
	_Error_name_index_1 = [...]uint8{0, 7, 21, 37}
)

func (i Error) String() string {
	switch {
	case 100 <= i && i <= 101:
		i -= 100
		return _Error_name_0[_Error_name_index_0[i]:_Error_name_index_0[i+1]]
	case 103 <= i && i <= 105:
		i -= 103
		return _Error_name_1[_Error_name_index_1[i]:_Error_name_index_1[i+1]]
	default:
		return fmt.Sprintf("Error(%d)", i)
	}
}

const (
	_Error_msg_0 = "User could not be foundUser already exists"
	_Error_msg_1 = "Not sure what happenedYou didn't send a good requestWorks on my machine"
)

var (
	_Error_msg_index_0 = [...]uint8{0, 23, 42}
	_Error_msg_index_1 = [...]uint8{0, 22, 52, 71}
)

func (i Error) Error() string {
	switch {
	case 100 <= i && i <= 101:
		i -= 100
		return _Error_msg_0[_Error_msg_index_0[i]:_Error_msg_index_0[i+1]]
	case 103 <= i && i <= 105:
		i -= 103
		return _Error_msg_1[_Error_msg_index_1[i]:_Error_msg_index_1[i+1]]
	default:
		return fmt.Sprintf("Error(%d)", i)
	}
}

var _ErrorNameToValue_map = map[string]Error{
	_Error_name_0[0:8]:   100,
	_Error_name_0[8:21]:  101,
	_Error_name_1[0:7]:   103,
	_Error_name_1[7:21]:  104,
	_Error_name_1[21:37]: 105,
}

func ErrorString(s string) (Error, error) {
	if val, ok := _ErrorNameToValue_map[s]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("%s is not the name of type Error")
}

func (i Error) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	msg, err := json.Marshal(i.Error())
	if err != nil {
		return b.Bytes(), err
	}
	name, err := json.Marshal(i.String())
	if err != nil {
		return b.Bytes(), err
	}
	json := fmt.Sprintf("{\"type\":%s,\"message\":%s}", name, msg)
	b.WriteString(json)
	return b.Bytes(), nil
}

type errStruct struct {
	name    string
	message string
}

func (i *Error) UnmarshalJSON(data []byte) error {
	var errData errStruct

	if err := json.Unmarshal(data, &errData); err != nil {
		return fmt.Errorf("Expecting a string, got %s", data)
	}

	val, err := ErrorString(errData.name)

	if err != nil {
		return err
	}

	*i = val

	return nil
}
`

type Golden struct {
	name   string
	input  string
	output string
}

var golden = []Golden{
	{"basic", basic_in, basic_out},
	{"offset", offset_in, offset_out},
	{"multiple", multiple_in, multiple_out},
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
