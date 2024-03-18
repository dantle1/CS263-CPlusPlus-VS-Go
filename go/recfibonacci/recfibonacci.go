package main

import (
	"flag"
	"fmt"
	"runtime/pprof"
	"os"
	"log"
)

var (
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
	infile     = flag.String("infile", "../data/matrix/matrix.in", "Matrices to multiply")
)

func fib (n int) int {
	if n <= 1 {return n}
	return fib(n-1) + fib(n-2)
}

func main(){
	flag.Parse()
	n := 50
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	fib := fib(n)
	fmt.Println(fib)
}

