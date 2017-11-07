package errorer

import (
	"strings"
	"testing"
)

var src = `
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

// func TestCommentMap(t *testing.T) {
// 	fset := token.NewFileSet()
// 	pkg := new(Package)
// 	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)

// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	pkg.check(fset, []*ast.File{f})

// 	ast.Inspect(f,
// 		func(node ast.Node) bool {
// 			decl, ok := node.(*ast.GenDecl)

// 			if !ok || decl.Tok != token.CONST {
// 				return true
// 			}

// 			for _, spec := range decl.Specs {
// 				vspec := spec.(*ast.ValueSpec)
// 				for _, name := range vspec.Names {
// 					t.Error(name)
// 					obj, _ := pkg.defs[name]

// 					val := obj.(*types.Const).Val()
// 					u64, _ := exact.Uint64Val(val)
// 					t.Error(u64)
// 					t.Error(vspec.Comment.Text())
// 				}
// 			}

// 			return false
// 		})
// }

func TestTryTry(t *testing.T) {
	var g Generator
	input := "package test\n" + src
	g.parsePackage(".", []string{"error_test.go"}, input)

	tokens := strings.SplitN(src, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("Need type declaration on first line")
	}
	g.Generate(tokens[1])

	t.Error(string(g.Format()))
}
