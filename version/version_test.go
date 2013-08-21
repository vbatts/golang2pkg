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

}
