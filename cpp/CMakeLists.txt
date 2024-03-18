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
        COMMAND mkdir -p ${PROF_DIRECTORY}
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

# For tricolor Garbage collector algorithm
set(GC1_SOURCES
    gc1/gc.cc
)
add_executable(gccc1 ${GC1_SOURCES})
cpuprofile(gccc1)

# For iterative matrix multiplication
set (ITMAT1_SOURCES
    itmat1/itmat1.cc
)
add_executable(itmat1 ${ITMAT1_SOURCES})
cpuprofile(itmat1)

# For recursive matrix multiplication 
set (RECMAT1_SOURCES
    recmat1/recmat1.cc
)
add_executable(recmat1 ${RECMAT1_SOURCES})
cpuprofile(recmat1)

# quick sort
set (QUICKSORT1_SOURCES
    quicksort1/quicksort1.cc)
add_executable(quicksort1 ${QUICKSORT1_SOURCES})
cpuprofile(quicksort1)

# merge sort
set (MERGESORT1_SOURCES
    mergesort1/mergesort1.cc)
add_executable(mergesort1 ${MERGESORT1_SOURCES})
cpuprofile(mergesort1)

# prime num
set (PRIMENUM_SOURCES
    primenum/primenum.cc)
add_executable(primenum ${PRIMENUM_SOURCES})
cpuprofile(primenum)

# fibonacci
set (ITERFIBONACCI_SOURCES
    iterfibonacci/iterfibonacci.cc)
add_executable(iterfibonacci ${ITERFIBONACCI_SOURCES})
cpuprofile(iterfbionacci)

set (RECFIBONACCI_SOURCES
    recfibonacci/recfibonacci.cc)
add_executable(recfibonacci ${RECFIBONACCI_SOURCES})
cpuprofile(recfibonacci)