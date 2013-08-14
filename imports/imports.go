/*
how to treat any given "import" that is attempting to be packaged
*/
package imports

import (
	"go/build"
	"io"
  "sort"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
Call out to `go get`. root will be set as GOPATH
*/
func GetPkgSource(pkg string, root string) error {
	err := os.Setenv("GOPATH", root)
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

// Convenience to see whether the import path is one that we'd be interested in
func isImportablePath(path string) bool {
	if strings.Contains(path, "/_") || strings.HasPrefix(path, "_") || strings.HasPrefix(path, ".") {
		return false
	}
	return true
}

// Convenience to see whether the file path is a possible go source type source file
func isSourceFile(path string) bool {
	if filepath.Ext(path) == ".go" ||
		filepath.Ext(path) == ".s" ||
		filepath.Ext(path) == ".c" ||
		filepath.Ext(path) == ".h" {
		return true
	}
	return false
}

type Imports []*build.Package
func (i Imports) Len() int { return len(i) }
func (i Imports) Swap(k,j int) { i[k], i[j] = i[j], i[k] }

type byImportPath struct { Imports }
func (bip byImportPath) Less(i, j int) bool { return bip.Imports[i].ImportPath < bip.Imports[j].ImportPath }

/*
Scan basepath and find the import'able paths relative to it
*/
func FindImports(basepath string) (Imports, error) {
	set := map[string]bool{} // unique keys
	findImportFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() && isSourceFile(path) {
			lib := strings.TrimPrefix(filepath.Dir(path), basepath+"/src/")
			// if the lib string is _not_ in our set and import path is sane
			if _, found := set[lib]; !found && isImportablePath(lib) {
				set[lib] = true
			}
		}
		return nil
	}

	pkgs := Imports{}
	err := filepath.Walk(basepath, findImportFn)
	if err != nil {
		return pkgs, err
	}

	for lib, _ := range set {
		if pkg, err := build.ImportDir(filepath.Join(basepath, "src", lib), 0); err == nil {
			pkgs = append(pkgs, pkg)
		}
	}
  sort.Sort(byImportPath{pkgs})
	return pkgs, nil
}
