package errorer

// buildNameToMsgMap allows the Error() method to retrieve the error message
// quickly. we do need to generate the slice, though.
func (g *Generator) buildNameToMsgMap(runs [][]Value, typeName string, runsThreshold int) {
	// TODO: account for runs. for now we can assume we don't have any.
	// n tracks where in the 'name' slice we are now
	var n int
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t_^s_msg%s[%d:%d]: %s,\n", typeName, "", n, n+len(value.name), &value)
			n += len(value.name)
		}
	}
}
