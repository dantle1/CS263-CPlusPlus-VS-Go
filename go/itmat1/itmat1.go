package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

var (
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
	infile     = flag.String("infile", "../data/matrix/matrix.in", "Matrices to multiply")
)

// Function to multiply two matrices
func multiplyMatrices(mat1, mat2 [][]float64) [][]float64 {
	rows1 := len(mat1)
	cols1 := len(mat1[0])
	cols2 := len(mat2[0])

	result := make([][]float64, rows1)
	for i := range result {
		result[i] = make([]float64, cols2)
	}

	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	return result
}

func main() {
	flag.Parse()

	file, _ := os.Open(*infile)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	dimensions := strings.Fields(scanner.Text())

	m, _ := strconv.Atoi(dimensions[0])
	n, _ := strconv.Atoi(dimensions[1])
	l, _ := strconv.Atoi(dimensions[2])

	A := make([][]float64, m)
	for i := range A {
		scanner.Scan()
		line := scanner.Text()
		row := strings.Fields(line)

		A[i] = make([]float64, n)
		for j := range A[i] {
			A[i][j], _ = strconv.ParseFloat(row[j], 64)
		}
	}

	B := make([][]float64, n)
	for i := range B {
		scanner.Scan()
		line := scanner.Text()
		row := strings.Fields(line)

		B[i] = make([]float64, l)
		for j := range B[i] {
			B[i][j], _ = strconv.ParseFloat(row[j], 64)
		}
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	C := multiplyMatrices(A, B)

	for r := range C {
		for _, val := range C[r] {
			fmt.Printf("%d\n", int(val))
		}
	}
}
