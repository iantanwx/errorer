package errorer

import "fmt"

const errStrToValueMap = `func %[1]sString(s string) ($[1]s, error) {
	if val, ok := _%[1]sNameToValue_map[s]; ok {
		return val, nil
	}

	return 0, fmt.Errorf("%%s is not the name of type %[1]s")
}
`

// adapted from github.com/alvaroloes/enumer
func (g *Generator) buildErrStrToValueMap(runs [][]Value, typeName string) {
	var n int
	var runID string
	// called after Stringer and Error are in the buffer
	g.Printf("\nvar _%sNameToValue_map = map[string]%s{\n", typeName, typeName)
	hasRuns := len(runs) > 1 && len(runs) <= 10

	for i, values := range runs {
		if hasRuns {
			runID = fmt.Sprintf("_%d", i)
		}

		for _, value := range values {
			g.Printf("\t_%s_name%s[%d:%d]: %s,\n", typeName, runID, n, n+len(value.name), &value)
			n += len(value.name)
		}
	}

	g.Printf("}\n\n")
	// now build our function
	g.Printf(errStrToValueMap, typeName)
}

// Arguments:
//	[1]: type name
const jsonMethods = `
func (i %[1]s) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	msg, err := json.Marshal(i.Error())
	if err != nil {
		return b, err
	}
	name, err := json.Marshal(i.String())
	if err != nil {
		return b, err
	}
	json := fmt.Sprintf("{\"type\":%%s,\"message\":%%s}", name, msg)
	b.WriteString(json)
	return b.Bytes(), nil
}
`

func (g *Generator) buildJsonMethods(typeName string) {
	g.Printf(jsonMethods, typeName)
}
