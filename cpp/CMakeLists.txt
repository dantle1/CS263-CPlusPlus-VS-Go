cmake_minimum_required(VERSION 3.22.2)
project(263 LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 23)

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

set(CMAKE_BUILD_TYPE RelWithDebInfo)

set(LPROFILER
    /usr/lib/aarch64-linux-gnu/libtcmalloc_and_profiler.so.4
)

# For havlak loop finding algorithm
set(HAVLAK_SOURCES
    havlak/LoopTesterApp.cc
    havlak/mao-loops.h
    havlak/mao-loops.cc
)
add_executable(havlakcc ${HAVLAK_SOURCES})
# So that libprofiler is actually linked if not invoked
target_link_options(havlakcc PRIVATE -Wl,--no-as-needed ${LPROFILER} -Wl,--as-needed)

# For reverse complement from 
# https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/revcomp-gpp-1.html
set(REVCMP_SOURCES
    reverse-complement/reverse-complement.cc
)
add_executable(revcmpcc ${REVCMP_SOURCES})
target_link_options(revcmpcc PRIVATE -Wl,--no-as-needed ${LPROFILER} -Wl,--as-needed)

# For three coloring DEBUG (sample solution)
set(THREECOLORDEBUG_SOURCES
    three-coloring/three-coloring-debug.cc
)
add_executable(threecolor_debug ${THREECOLORDEBUG_SOURCES})

# For three coloring (my implementation)
set(THREECOLOR_SOURCES 
    three-coloring/three-coloring.cc
)
add_executable(threecolorcc ${THREECOLOR_SOURCES})
target_link_options(threecolorcc PRIVATE -Wl,--no-as-needed ${LPROFILER} -Wl,--as-needed)