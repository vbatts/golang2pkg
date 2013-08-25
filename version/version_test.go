package version

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestBasics(t *testing.T) {
	v := Version{}

	if v.String() != "0.0" {
		t.Errorf("Empty version did not match '0.0', but returned %s", v.String())
	}

	b, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("error getting Abs path: %s", err)
	}
	v = FromDir(b, filepath.Dir(b))
	fmt.Println(v)

	pth := "/home/vbatts/opt/go/src/launchpad.net/goyaml"
	s := BzrRevno(pth)
	v.Path = pth
	if s != v.String() {
		t.Errorf("versions do not match for `bzr`. Expected [%s], got [%s]", s, v.String())
	}

	pth = "/home/vbatts/opt/go/src/cgl.tideland.biz"
	s = HgTip(pth)
	v.Path = pth
	if s != v.String() {
		t.Errorf("versions do not match for `hg`. Expected [%s], got [%s]", s, v.String())
	}
}
