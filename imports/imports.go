/*
how to treat any given "import" that is attempting to be packaged
*/
package imports

import (
  "path/filepath"
  "os"
  "strings"
)

type Import struct {
  Base string // relative base path of this src
  Name string // the import name
}

func (i Import) String() string {
  return i.Name
}

/*
Scan basepath and find the import'able paths relative to it
*/
func FindImports(basepath string) ([]Import, error) {
  set := map[string]bool{} // unique keys
  findImportFn := func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if info.Mode().IsRegular() && filepath.Ext(path) == ".go" {
      lib := strings.TrimPrefix(filepath.Dir(path), basepath + "/")
      if _, found := set[lib]; !found {
        set[lib] = true
      }
    }
    return nil
  }

  err := filepath.Walk(basepath, findImportFn)

  found_imports := []Import{}
  for lib ,_ := range set {
      found_imports = append(found_imports, Import{
        Base: basepath,
        Name:  lib})
  }
  if err != nil {
    return found_imports, err
  }
  return found_imports, nil
}

