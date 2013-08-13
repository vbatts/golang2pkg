package main

import (
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: provide package\n")
		os.Exit(1)
	}
	t_path := path.Join(os.TempDir(), "golang2pkg-asdfasdf")
	err := os.Mkdir(t_path, 0755)
	if err != nil && os.IsExist(err) {
		fmt.Fprintf(os.Stderr, "Re-using directory '%s'\n", t_path)
  } else if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
	//defer os.RemoveAll(t_path)

	err = os.Setenv("GOPATH", t_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("go", "get", "-d", os.Args[1])
	fmt.Printf("%#v\n", cmd)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
  if len(buf) > 0 {
	  fmt.Printf("%s\n", buf)
  }

}

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

func demo() {
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
