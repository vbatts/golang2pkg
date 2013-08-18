package version

import (
  //"fmt"
  "testing"
)

func TestBasics(t *testing.T) {
  v := Version{}

  if v.String() != "0.0" {
    t.Errorf("Empty version did not match '0.0', but returned %s", v.String())
  }
}
