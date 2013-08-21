/*
helpers to get the version information for a checked out library
*/
package version

import (
	"bytes"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Version struct {
	Path  string // path of the import that this Version is for
	Value string // Manually setting a version string
}

/*
Stringer formating of a Version, regardless of which VCS or an explicit Value
*/
func (v Version) String() string {
	if len(v.Value) > 0 {
		return v.Value
	}

	// if there is a path and a VCS for that path, then call the corresponding func
	if len(v.Path) > 0 && len(hasVcsDir(v.Path)) > 0 {
		return VCSDirs[hasVcsDir(v.Path)](v.Path)
	}

	return "0.0"
}

/*
Version control versioner

Takes an argument of the path to check, and returns a string of the version
*/
type VcsFn func(string) string

/*
map of Version control directory names, and the func to call on
a path of that type, in order to get a version.
*/
var VCSDirs = map[string]VcsFn{
	".git": GitHead,
	".bzr": nothing,
	".hg":  nothing,
}

// XXX placeholder for actual vcs functions
func nothing(p string) string { return p }

/*
Given a build.Package, check the Dir for a directory matching CVSDirRe,
if nothing check parents until SrcRoot
*/
func FromPackage(pkg *build.Package) Version {
	return FromDir(pkg.Dir, pkg.SrcRoot)
}

/*
Check for VCS in pth file path, not to exceed root.
*/
func FromDir(pth, root string) Version {
	for {
		if len(hasVcsDir(pth)) > 0 {
			return Version{Path: pth}
		}
		pth = path.Dir(pth)
		if pth == root {
			break
		}
	}
	return Version{}
}

func hasVcsDir(pth string) string {
	for vcs, _ := range VCSDirs {
		if fi, err := os.Stat(path.Join(pth, vcs)); err == nil && fi.IsDir() {
			return vcs
		}
	}
	return ""
}

/*
functionally equivalent to `git rev-parse --short HEAD`
*/
func GitHead(pth string) (hash string) {
	file, err := os.Open(path.Join(pth, ".git", "HEAD"))
	if err != nil {
		return ""
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	ref_path := strings.TrimPrefix(strings.Trim(bytes.NewBuffer(buf).String(), "\n "), "ref: ")
	ref_file, err := os.Open(path.Join(pth, ".git", ref_path))
	if err != nil {
		return ""
	}
	ref_buf, err := ioutil.ReadAll(ref_file)
	if err != nil {
		return ""
	}
	return strings.Trim(bytes.NewBuffer(ref_buf).String(), "\n ")[0:7]
}
