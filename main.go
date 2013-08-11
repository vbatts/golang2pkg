package main

import (
  "flag"
)

func main() {
  flag.Parse()

  /*
  $> golang2pkg labix.org/v2/mgo
  - check if arg is in the current GOROOT or GOPATH
   -- no:
      ** mktmpdir and set GOPATH env
      ** invoke `go get -d arg`
   -- yes:
    ** cp to mktmpdir
  - scan for deps to cp to mktmpdir
  - make a list of libraries to pkg
  */
}
