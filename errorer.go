package errorer

import (
	"bytes"
	"fmt"
)

const errorMethod = `
		
`

// buildErrorMethod ensures that our struct satisfies the error interface
func (g *Generator) buildErrorMethod(run [][]Value, typeName string, runsThreshold int) {
	// g.Printf(jsonMethods, typeName)
}

// buildNameToMsgMap allows the Error() method to retrieve the error message
// quickly. we do need to generate the slice, though.
func (g *Generator) buildNameToMsgMap(runs [][]Value, typeName string, runsThreshold int) {
	// TODO: account for runs. for now we can assume we don't have any.
	// n tracks where in the 'name' slice we are now
	var n int
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t_^s_name%s[%d:%d]: %s,\n", typeName, "", n, n+len(value.name), &value)
			n += len(value.name)
		}
	}
}

// TODO: implement multiple-run version
// single-run fn for generating msg string and corresponding index
func (g *Generator) declareIndexAndMsgVar(run []Value, typeName string) {
	// pass empty string since this is only for a single run
	index, msg := g.createIndexAndMsgDecl(run, typeName, "")
	g.Printf("const %s\n", msg)
	g.Printf("var %s\n", index)
}

// this generates the rhs of const _TypeName_msg and var _TypeName_msg_index
// var and const are added by the caller (declareIndexAndMsgVar for now)
func (g *Generator) createIndexAndMsgDecl(run []Value, typeName string, suffix string) (string, string) {
	// this buffer holds all the messages. they get concatenated together
	b := new(bytes.Buffer)
	indexes := make([]int, len(run))
	for i := range run {
		b.WriteString(run[i].msg)
		indexes[i] = b.Len()
	}
	msgConst := fmt.Sprintf("_%s_msg%s = %q", typeName, suffix, b.String())
	msgLen := b.Len()
	b.Reset()
	fmt.Fprintf(b, "_%s_index%s = [...]uint%d{0, ", typeName, suffix, usize(msgLen))
	for i, v := range indexes {
		if i > 0 {
			fmt.Fprintf(b, ", ")
		}
		fmt.Fprintf(b, "%d", v)
	}
	fmt.Fprintf(b, "}")
	return b.String(), msgConst
}
