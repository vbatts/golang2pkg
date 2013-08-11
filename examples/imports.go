package main

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"strings"
)

var src = `
package hurp

import _ "fmt"
import (
  _ "io"
  _ "bytes"
  _ "github.com/vbatts/go-httplog"
)

var c = 0
`

func main() {
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}
	for _, i_spec := range f.Imports {
		path := strings.Trim(i_spec.Path.Value, "\"")

		p, err := build.Import(path, build.Default.GOROOT, build.FindOnly)
		if err != nil {
			fmt.Printf("%s -- Need to fetch\n", path)
			return
		}
		fmt.Printf("%s -- %#v\n", path, p)
	}
}
