package imports

import (
	"fmt"
	"go/build"
	"testing"
)

func TestFind(t *testing.T) {
	for _, src := range build.Default.SrcDirs() {
		i, e := FindImports(src)
		if e != nil {
			t.Fatal(e)
		}
		//fmt.Printf("%#v\n", i)
		fmt.Println(i)
	}
}
