cmake_minimum_required(VERSION 3.27.6)
project(263 LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 23)
SET(CMAKE_OSX_ARCHITECTURES "arm64")

# havlakcc 
set(HAVLAK_SOURCES
    havlak/LoopTesterApp.cc
    havlak/mao-loops.h
    havlak/mao-loops.cc
)
add_executable(havlakcc ${HAVLAK_SOURCES})
target_link_libraries(havlakcc -lprofiler)