/*
helpers to get the version information for a checked out library
*/
package version

import (
	"go/build"
	"os"
	"path"
)

type Version struct {
	Path  string // path of the import that this Version is for
	Value string // Manually setting a version string
}

func (v Version) String() string {
	if len(v.Value) > 0 {
		return v.Value
	}

	if len(v.Path) > 0 {
	}

	return "0.0"
}

var VCSDirs = []string{".git", ".bzr", ".hg"}

/*
Given a build.Package, check the Dir for a directory matching CVSDirRe,
if nothing check parents until SrcRoot
*/
func FromPackage(pkg *build.Package) Version {
	pth := pkg.Dir
	for {
		if hasVcsDir(pth) {
			return Version{Path: pth}
		}
		pth = path.Dir(pth)
		if pth == pkg.SrcRoot {
			break
		}
	}
	return Version{}
}

func hasVcsDir(pth string) bool {
	for _, vcs := range VCSDirs {
		if fi, err := os.Stat(path.Join(pth, vcs)); err == nil && fi.IsDir() {
			return true
		}
	}
	return false
}
