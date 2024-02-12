# CS263-C++ VS Go

## C++ Profiling:
We used [Google Pprof](https://github.com/gperftools/gperftools/tree/master) to profile our C++ programs. Based on the documentation, the tools get best performance on modern GNU/Linux systems -- we performed all of the analysis inside of a Debian 12 Docker container. For more information on compatibility, look [here](https://github.com/gperftools/gperftools/blob/master/INSTALL).  

### Setup:
As mentioned, our experiments were run inside of a Docker container. As Docker newbies, we opted to use the VS Code [Dev Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers), which requires [Docker Desktop](https://www.docker.com/products/docker-desktop/). For a demo on how to set up a Dev Container in VS Code, look [here](https://www.youtube.com/watch?v=b1RavPr_878).  

Besides `cmake`, and `build-essential`, the only other package you will need is `google-perftools`.

#### CPU Profiling
More to come!

#### Memory Profiling
More to come!
