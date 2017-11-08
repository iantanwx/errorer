# Installation
```
go get -u github.com/iantanwx/errorer
```

# Usage
Input:
```
error.go

// this is a custom error type
// we add the go:generate annotation
//go:generate errorer -type=Error
type Error int

// error message is a comment annotation
// generated output can be passed to fmt.Sprint*, json.Marshal/Unmarshal
const (
	NotFound         Error = iota //User could not be found
	AlreadyExists                 //User already exists
	NotSure                       //Not sure what happened
	BadRequestData                //You didn't send a good request
	WorksOnMyMachine              //Works on my machine
)

```

1. `go generate`
2. ...
3. Profit

Output:
```
error_string.go

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
```

The resulting output satisfies:

- `fmt.Stringer`
- `error.Error`
- `json.Marshaler`
- `json.Unmarshaler`

# Inspiration

This package is heavily inspired by and adapted from Rob Pike's stringer and github.com/alvaroloes/enumer

It is still in a very raw state. Use at your own risk. PRs are welcome.
