package imports

import (
  "testing"
  "go/build"
  "fmt"
)

func TestFind(t *testing.T) {
  i, e := FindImports(build.Default.GOROOT + "/src/pkg")
  if e != nil {
    t.Fatal(e)
  }
  //fmt.Printf("%#v\n", i)
  fmt.Println(i)
}
