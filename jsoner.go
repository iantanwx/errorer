package errorer

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
