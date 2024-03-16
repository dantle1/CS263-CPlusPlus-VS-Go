package main

import (
	// "fmt"
	"github.com/dantle1/CS263-CPlusPlus-VS-Go/go/permutation1/permutation"
	"log"
	"os"
	"runtime/pprof"
	"flag"
)

var (
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
)

func main() {
	flag.Parse()
	permutationsCh := make(chan []string)
	// var value []string
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go permutation.Heaps(permutationsCh, 11)
	_ = <-permutationsCh
	// for _, str := range value {
	// 	fmt.Printf("%s\n", str)
	// }
}
