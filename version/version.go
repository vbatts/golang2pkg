/*
helpers to get the version information for a checked out library
*/
package version

import (
	"github.com/vbatts/golang2pkg/host"
	"log"
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
