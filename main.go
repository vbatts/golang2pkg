package main

import (
  "flag"
)

func main() {
  flag.Parse()

  /*
  $> golang2pkg labix.org/v2/mgo
  - check if arg is in the current GOROOT or GOPATH or build.IsLocalImport
   -- no:
      ** mktmpdir and set GOPATH env
      ** invoke `go get -d arg`
   -- yes:
    ** cp to mktmpdir
  - scan for deps to cp to mktmpdir
  - scan tmpdir for root import to package
  - deduce the type of build artificat to generate
  - for each root import
    ** render a relavant version string (git, hg, bzr, etc)
    ** make a *.tar.gz like `filepath.Base(path)-Version().tar.gz` 
    ** is the source-only flag was passed
     -- no: render the build artifact.
     -- yes: then print source name

  - rendering the build artifact
    ** *build.Package.IsCommand() wiil show whether the import results in an executable.
    ** this means for a tree with source directories, that multiple rpms/debs may need to be generated
    ** _any_ final executable build ought to be packaged with no runtime deps (unless linked to C libs)
      -- thankfully there is a post rpm job that will append runtime requires based on ldd linkage
  */
}
