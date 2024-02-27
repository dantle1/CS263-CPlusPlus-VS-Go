cmake_minimum_required(VERSION 3.22.2)
project(263 LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 23)

set(GPERFTOOLS_INCLUDE_DIR "/usr/include/gperftools")
include_directories(${GPERFTOOLS_INCLUDE_DIR})

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

set(PROF_DIRECTORY ${CMAKE_BINARY_DIR}/prof)

add_custom_target(clean-prof
    COMMAND ${CMAKE_COMMAND} -E remove ${PROF_DIRECTORY}/*
    COMMENT "Cleaning profiling files"
)

set(CMAKE_BUILD_TYPE RelWithDebInfo)

function(cpuprofile PROGRAM_NAME)
    target_link_libraries(${PROGRAM_NAME} tcmalloc_and_profiler)
    add_custom_target(${PROGRAM_NAME}.cprof
        COMMAND CPUPROFILE=${PROF_DIRECTORY}/${PROGRAM_NAME}.cprof ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}/${PROGRAM_NAME}
        DEPENDS ${PROGRAM_NAME}
        WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
    )
endfunction()

# The base version, no optimizations 
# For havlak loop finding algorithm
set(HAVLAK1_SOURCES
    havlak1/LoopTesterApp.cc
    havlak1/mao-loops.h
    havlak1/mao-loops.cc
)
add_executable(havlakcc1 ${HAVLAK1_SOURCES})
cpuprofile(havlakcc1)

# Optimized by changing BasicBlockMap to a vector...we get 2X speedup!!
set(HAVLAK2_SOURCES
    havlak2/LoopTesterApp.cc
    havlak2/mao-loops.h
    havlak2/mao-loops.cc
)
add_executable(havlakcc2 ${HAVLAK2_SOURCES})
cpuprofile(havlakcc2)

# Optimized by changing BasicBlockMap to an int array...similar speedup to a vector
set(HAVLAK2.5_SOURCES
    havlak2.5/LoopTesterApp.cc
    havlak2.5/mao-loops.h
    havlak2.5/mao-loops.cc
)
add_executable(havlakcc2.5 ${HAVLAK2.5_SOURCES})
cpuprofile(havlakcc2.5)

# For reverse complement from 
# https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/revcomp-gpp-1.html
set(REVCMP_SOURCES
    reverse-complement/reverse-complement.cc
)
add_executable(revcmpcc ${REVCMP_SOURCES})
cpuprofile(revcmpcc)

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
cpuprofile(threecolorcc)