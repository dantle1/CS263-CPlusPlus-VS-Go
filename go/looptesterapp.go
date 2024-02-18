// Copyright 2011 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Test Program for the Havlak loop finder.
//
// This program constructs a fairly large control flow
// graph and performs loop recognition. This is the Go
// version.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/dantle1/CS263-CPlusPlus-VS-Go/go/havlakloopfinder"
)

//======================================================
// Testing Code
//======================================================

func buildDiamond(cfgraph *havlakloopfinder.CFG, start int) int {
	bb0 := start
	havlakloopfinder.NewBasicBlockEdge(cfgraph, bb0, bb0+1)
	havlakloopfinder.NewBasicBlockEdge(cfgraph, bb0, bb0+2)
	havlakloopfinder.NewBasicBlockEdge(cfgraph, bb0+1, bb0+3)
	havlakloopfinder.NewBasicBlockEdge(cfgraph, bb0+2, bb0+3)

	return bb0 + 3
}

func buildConnect(cfgraph *havlakloopfinder.CFG, start int, end int) {
	havlakloopfinder.NewBasicBlockEdge(cfgraph, start, end)
}

func buildStraight(cfgraph *havlakloopfinder.CFG, start int, n int) int {
	for i := 0; i < n; i++ {
		buildConnect(cfgraph, start+i, start+i+1)
	}
	return start + n
}

func buildBaseLoop(cfgraph *havlakloopfinder.CFG, from int) int {
	header := buildStraight(cfgraph, from, 1)
	diamond1 := buildDiamond(cfgraph, header)
	d11 := buildStraight(cfgraph, diamond1, 1)
	diamond2 := buildDiamond(cfgraph, d11)
	footer := buildStraight(cfgraph, diamond2, 1)
	buildConnect(cfgraph, diamond2, d11)
	buildConnect(cfgraph, diamond1, header)

	buildConnect(cfgraph, footer, from)
	footer = buildStraight(cfgraph, footer, 1)
	return footer
}

var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")

func main() {
	/* This section starts a CPU profile and 'defer's the execution of Stop until
	main exits */
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	lsgraph := havlakloopfinder.NewLSG()
	cfgraph := havlakloopfinder.NewCFG()

	cfgraph.CreateNode(0) // top
	buildBaseLoop(cfgraph, 0)
	cfgraph.CreateNode(1) // bottom
	havlakloopfinder.NewBasicBlockEdge(cfgraph, 0, 2)

	for dummyloop := 0; dummyloop < 15000; dummyloop++ {
		havlakloopfinder.FindHavlakLoops(cfgraph, havlakloopfinder.NewLSG())
	}

	n := 2

	for parlooptrees := 0; parlooptrees < 10; parlooptrees++ {
		cfgraph.CreateNode(n + 1)
		buildConnect(cfgraph, 2, n+1)
		n = n + 1

		for i := 0; i < 100; i++ {
			top := n
			n = buildStraight(cfgraph, n, 1)
			for j := 0; j < 25; j++ {
				n = buildBaseLoop(cfgraph, n)
			}
			bottom := buildStraight(cfgraph, n, 1)
			buildConnect(cfgraph, n, top)
			n = bottom
		}
		buildConnect(cfgraph, n, 1)
	}

	havlakloopfinder.FindHavlakLoops(cfgraph, lsgraph)

	/* This section starts a heap profile and stops after one iteration
	of loop finding */
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 50; i++ {
		havlakloopfinder.FindHavlakLoops(cfgraph, havlakloopfinder.NewLSG())
	}

	fmt.Printf("# of loops: %d (including 1 artificial root node)\n", lsgraph.NumLoops())
	lsgraph.CalculateNestingLevel()
}
