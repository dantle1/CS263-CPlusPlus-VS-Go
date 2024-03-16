package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"log"
	"github.com/dantle1/CS263-CPlusPlus-VS-Go/go/graph/graph"
)

var (
	cpuprofile = ""
	memprofile = ""
	alg        = flag.String("alg", "floyd", "Algorithm we are profiling/optimizing")
	infile     = flag.String("infile", "../data/weighted/weighted.in", "Input graph")
)

func main() {
	flag.Parse()

	file, _ := os.Open(*infile)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	n_m := strings.Fields(line)

	n, _ := strconv.Atoi(n_m[0])
	m, _ := strconv.Atoi(n_m[1])

	G := make([][]float64, n)
	for i := 0; i < n; i++ {
		G[i] = make([]float64, n)
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		line := scanner.Text()
		edge := strings.Fields(line)
		a, _ := strconv.Atoi(edge[0])
		b, _ := strconv.Atoi(edge[1])
		w, _ := strconv.ParseFloat(edge[2], 64)
		G[a][b] = w
	}

	f, err := os.Create("prof/" + *alg + ".cprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	if *alg == "floyd" {
		fmt.Printf("Running floyd\n")
		_ = graph.FloydWarshall(G)
	} else {
		fmt.Printf("other branch\n")
	}
}
