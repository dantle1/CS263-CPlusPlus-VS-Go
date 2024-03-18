package main

import (
	"flag"
	"fmt"
	"runtime/pprof"
	"time"
	"os"
	"log"
)

var (
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
	infile     = flag.String("infile", "../data/matrix/matrix.in", "Matrices to multiply")
)

func isPrime(n int64) bool {
	if n==1 || n==0 { return false }
	for i:=int64(2); i < n; i++ {
		if (n%i) == 0 {return false}
	}
	return true
}

func main(){
	flag.Parse();
	N := int64(400000);
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	start := time.Now()
	for i := int64(1); i<= N; i++ { 
		if isPrime(i) {
			//fmt.Println(i)}
		}
	}
	end := time.Now()
	fmt.Println("primenum time for ", N, ":", end.Sub(start))
}