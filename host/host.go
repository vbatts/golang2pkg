/*
For host type queries, like HasRpm, HasDpkg, HasDhMake, and HasRpmBuild
(or similar).

This will be used for determining how far you can go with making a build
artifact on a box that does not have all the tools needed.
*/
package host

import (
	"os/exec"
)

//var DefaultRoot = "/"

/*
Searches PATH for cmd, and returns the path to it.
Empty string is lack of presence.
*/
func Command(cmd string) string {
	path, err := exec.LookPath(cmd)
	if err != nil {
		return ""
	}
	return path
}

// Whether rpmbuild is in PATH
func HasRpmBuild() bool {
	return len(Command("rpmbuild")) > 0
}

// Path to rpmbuild (if present)
func RpmBuild() string {
	return Command("rpmbuild")
}

// Whether rpm is in PATH
func HasRpm() bool {
	return len(Command("rpm")) > 0
}

// Path to rpm (if present)
func Rpm() string {
	return Command("rpm")
}

// Whether dpkg is in PATH
func HasDpkg() bool {
	return len(Command("dpkg")) > 0
}

// Path to dpkg (if present)
func Dpkg() string {
	return Command("dpkg")
}

// Whether dh_make is in PATH
func HasDhMake() bool {
	return len(Command("dh_make")) > 0
}

// Path to dh_make (if present)
func DhMake() string {
	return Command("dh_make")
}

// Path to hg (if present)
func Hg() string {
	return Command("hg")
}

// Path to bzr (if present)
func Bzr() string {
	return Command("bzr")
}

// Path to git (if present)
func Git() string {
	return Command("git")
}
