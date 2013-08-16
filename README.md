golang2pkg
==========

A utility to manage the packaging of a golang binary or source library into
a distribution package.

WARNING: This is presently a work-in-progress and is not functional.
         expect this warning to drop off and a stable tag once usable

Use Cases
---------
![UseCases](https://raw.github.com/vbatts/golang2pkg/master/UseCases.png)

* to end up with an binary, regardless of dependencies
* to package an source library that is fetched
* to package an already checkouted source library
* to package local source directory

To achieve this, there may likely need to be flags to provide the version
number (default will make assumptions on the %version-%release based on VCS
revision and datetime), provide the packages import path (in the case of local
source directory). 

Another nicety of rendering the build time dependencies of a package, is that
the golang compiler wholly handles the *.c and *.s files (with 6c and 6a or
similar), so apart from having a "BuildRequires" of 'golang', it ought to only
have "BuildRequires" of packages that "Provide" the meta
'golang("some.com/fqdn/library")', and not gcc and friends.

Though there is the case of linking with cgo, to a shared object library, which
would require that library and its headers to be present on the system.


Package type assumptions
------------------------

For the generation of the build artifact, an assumption is made based on the
presence of dpkg, rpm or neither.  Though the construction of the artifact is
not platform specific (you can generate RPMs from debian, etc).


Layout of RPMs
--------------

Since a single source code repository can have its subdirectories referenced
individually, then either the rpm spec could generate a list of RPMs for these
subdirectories, each with their own:

	Provides: golang('some.com/fqdn/library')

and the package name appends the sbdirectory name.

or

Just land the whole source tree, and enumerate all the 'Provides' of the
subdirectories in this single package.


Naming of RPMs
--------------

Since package naming guidelines do not allow for '.' and '/' in the %{name},
how then should the FQDN import path be accounted for in the Name?

Versioning
----------

The .git/.hg/.bzr will be used to collect the revision hash for a package
versioning, but that meta data directory will not be included in the end system
package.

bzr:

	$> bzr revno

git:

	$> git rev-parse --short HEAD


hg (this is ugly):

	$> hg tip |grep ^changeset: | awk '{ print $2 }' | tr ':' '-'


Executables
-----------

Often, executables are at the root of a repo, but that is not always the case.
There might ought to be a search for ELF executables to put them into bindir.
This is a little muddied by demos or example apps, that may be alright to
compile, but not to land in /usr/bin. This own-ness would be left up to the
packager to review %files (or similar).


spectemplate-golang
-------------------

First effort at hand jamming an RPM for a golang source library, for the
'launchpad.net/goyaml' project:
http://www.hashbangbash.com/~vbatts/golang/golang-goyaml-bzr47-1.fc19.src.rpm

