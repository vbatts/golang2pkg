package imports

import (
	"go/build"
	"testing"
)

func TestFind(t *testing.T) {
	def_pkgs := Imports{}
	for _, src := range build.Default.SrcDirs() {
		i, e := FindImports(src)
		if e != nil {
			t.Fatal(e)
		}
		def_pkgs = append(def_pkgs, i...)
	}

	i, e := FindImportsDefault()
	if e != nil {
		t.Fatal(e)
	}
	if len(i) != len(def_pkgs) {
		t.Errorf("FindImports of SrcDirs returned [%d], but FindImportsDefault returned [%d]",
			len(def_pkgs),
			len(i))
	}
}
