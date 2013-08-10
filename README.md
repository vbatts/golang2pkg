golang2pkg
==========

A utility to manage the packaging of a golang binary or source library into
a distribution package.

Use Cases
---------
![UseCases](../blob/master/UseCases.png?raw=true)

* to end up with an binary, regardless of dependencies
* to package an source library that is fetched
* to package an already checkouted source library
* to package local source directory

To achieve this, there may likely need to be flags to provide the version
number (default will make assumptions on the %version-%release based on VCS
revision and datetime), provide the packages import path (in the case of local
source directory). 

