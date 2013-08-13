/*
how to treat any given "import" that is attempting to be packaged
*/
package imports

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
Call out to `go get`. root will be set as GOPATH
*/
func GetPkgSource(pkg string, root string) error {
	err = os.Setenv("GOPATH", root)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "get", "-d", pkg)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	return cmd.Run()
}

// XXX this is largely already available in go/build ...
type Import struct {
	Base string // relative base path of this src
	Name string // the import name
}

func (i Import) String() string {
	return i.Name
}

/*
Scan basepath and find the import'able paths relative to it

XXX this is largely already available in go/build ...
*/
func FindImports(basepath string) ([]Import, error) {
	set := map[string]bool{} // unique keys
	findImportFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() && filepath.Ext(path) == ".go" {
			lib := strings.TrimPrefix(filepath.Dir(path), basepath+"/")
			if _, found := set[lib]; !found {
				set[lib] = true
			}
		}
		return nil
	}

	err := filepath.Walk(basepath, findImportFn)

	found_imports := []Import{}
	for lib, _ := range set {
		found_imports = append(found_imports, Import{
			Base: basepath,
			Name: lib})
	}
	if err != nil {
		return found_imports, err
	}
	return found_imports, nil
}
