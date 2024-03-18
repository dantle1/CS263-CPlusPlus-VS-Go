# CS263-C++ VS Go

## C++ Profiling:
We used [Google Pprof](https://github.com/google/pprof) for analyzing profiles, and [gperftools](https://github.com/gperftools/gperftools) to generate profiles. Based on [gperftools documentation](https://github.com/gperftools/gperftools/blob/master/INSTALL), the tools get best performance on modern GNU/Linux systems -- we performed all of the analysis inside of a Debian 12 Docker container.  

## Go Profiling:
We used [Go pprof command](https://go.dev/blog/pprof) for analyzing profiles, and [Go Pprof](https://pkg.go.dev/runtime/pprof) to generate profiles. These should work well independent of architecture to our knowledge.

## Setup:
As Docker newbies, we opted to use the VS Code [Dev Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers), which requires [Docker Desktop](https://www.docker.com/products/docker-desktop/). For a demo on how to set up a Dev Container in VS Code, look [here](https://www.youtube.com/watch?v=b1RavPr_878). All of the dependencies should be installed when you open the Dev Container (see `.devcontainer/` for more info).

## Profiling

### C++

To generate a profile for one of the programs, simply run `make program.cprof` for a CPU profile or `make program.mprof` for a memory profile (this profiles the whole program execution, where the memory profile will contain multiple profiles taken throughout the program). You may also import `<profiler.h>` from gperftools and explicitly start and stop a profile with a call to StartProfiler and StopProfiler.

### Go
To generate a profile for one of the programs, simple run `make program.cprof` for a CPU profile or `make program.mprof` for a memory profile. For more information, see [this](https://go.dev/blog/pprof). Some of the programs may not yet be configured for memory profiling, so this resource may be extra useful in those cases.
