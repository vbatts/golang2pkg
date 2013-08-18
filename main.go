package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/vbatts/golang2pkg/imports"
	"github.com/vbatts/golang2pkg/osutil"
)

func main() {
	var (
		output string
		root   string
		debug  bool
		//importname string
	)

	cwd, err := os.Getwd()
	if err != nil {
		cwd = os.TempDir()
	}

	root = fmt.Sprintf("%s/golang2pkg-%d", os.TempDir(), os.Getpid())

	//flag.StringVar(&importname, "importname", "", "overide the qualified import path (for a relative path packaging)")
	flag.StringVar(&root, "root", root, "root directory to treat as GOPATH for imports collected")
	flag.StringVar(&output, "output", cwd, "output directory for artifacts")
	flag.BoolVar(&debug, "debug", false, "debugging output")
	flag.Parse()

	if err = os.MkdirAll(path.Join(root, "src"), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}

	// if we can get the collection of current pkgs,
	// then symlink those present to our root, unless they're already there
	if pkgs, err := imports.FindImportsDefault(); err == nil {
		for _, pkg := range pkgs {
			for _, arg := range flag.Args() {
				if pkg.ImportPath == arg {
					if _, err = os.Stat(path.Join(root, "src", arg)); os.IsNotExist(err) {
						dest := path.Join(root, "src", arg)
						if err = os.MkdirAll(path.Dir(dest), 0755); err != nil {
							fmt.Fprintf(os.Stderr, "WARN: %s\n", err)
							continue
						}
						err = osutil.Copy(path.Join(pkg.SrcRoot, pkg.ImportPath), path.Dir(dest))
						if err != nil {
							fmt.Fprintf(os.Stderr, "WARN: %s\n", err)
						}
					}
				}
			}
		}
	}

	for _, arg := range flag.Args() {
		err = imports.GetPkgSource(arg, root)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Println(root)
	pkgs, err := imports.FindImports(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
	for _, pkg := range pkgs {
		//fmt.Println(pkg.SrcRoot, pkg.ImportPath)
		fmt.Printf("%#v\n", pkg)
		//fmt.Println(path.Join(pkg.SrcRoot, pkg.ImportPath))
	}

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
