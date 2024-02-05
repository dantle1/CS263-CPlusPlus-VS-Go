cmake_minimum_required(VERSION 3.22.2)
project(263 LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 23)

# havlakcc 
set(HAVLAK_SOURCES
    havlak/LoopTesterApp.cc
    havlak/mao-loops.h
    havlak/mao-loops.cc
)
set(LPROFILER
    /usr/lib/aarch64-linux-gnu/libprofiler.so.0
)
add_executable(havlakcc ${HAVLAK_SOURCES})
# So that libprofiler is actually linked if not invoked
target_link_options(havlakcc PRIVATE -Wl,--no-as-needed ${LPROFILER} -Wl,--as-needed)