# CS263-C++ VS Go

## C++ Profiling:
We used [Google Pprof](https://github.com/google/pprof) for analyzing profiles, and [gperftools](https://github.com/gperftools/gperftools) to generate profile. Based on [gperftools documentation](https://github.com/gperftools/gperftools/blob/master/INSTALL), the tools get best performance on modern GNU/Linux systems -- we performed all of the analysis inside of a Debian 12 Docker container.  

### Setup:
As Docker newbies, we opted to use the VS Code [Dev Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers), which requires [Docker Desktop](https://www.docker.com/products/docker-desktop/). For a demo on how to set up a Dev Container in VS Code, look [here](https://www.youtube.com/watch?v=b1RavPr_878). All of the dependencies should be installed when you open the Dev Container (see `.devcontainer/` for more info).

#### CPU Profiling
To generate a cpuprofile for one of the programs, simply run `make `

#### Memory Profiling
More to come!
