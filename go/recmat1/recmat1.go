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

// Function to add two matrices
func add(mat1 [][]float64, mat2 [][]float64) [][]float64 {
	rows := len(mat1)
	cols := len(mat1[0])
	result := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = mat1[i][j] + mat2[i][j]
		}
	}
	return result
}

// Function to subtract two matrices
func sub(mat1 [][]float64, mat2 [][]float64) [][]float64 {
	rows := len(mat1)
	cols := len(mat1[0])
	result := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = mat1[i][j] - mat2[i][j]
		}
	}
	return result
}

// Function to multiply two matrices using Strassen's algorithm
func multiplyMatrices(mat1, mat2 [][]float64) [][]float64 {
	rows1 := len(mat1)
	cols1 := len(mat1[0])
	cols2 := len(mat2[0])

	// Base case: if matrices are 1x1, perform simple multiplication
	if rows1 == 1 && cols1 == 1 && cols2 == 1 {
		result := make([][]float64, 1)
		result[0] = make([]float64, 1)
		result[0][0] = mat1[0][0] * mat2[0][0]
		return result
	}

	// Split matrices into quadrants
	newRows := rows1 / 2
	newCols := cols1 / 2

	A11 := make([][]float64, newRows)
	A12 := make([][]float64, newRows)
	A21 := make([][]float64, newRows)
	A22 := make([][]float64, newRows)

	B11 := make([][]float64, newRows)
	B12 := make([][]float64, newRows)
	B21 := make([][]float64, newRows)
	B22 := make([][]float64, newRows)

	C11 := make([][]float64, newRows)
	C12 := make([][]float64, newRows)
	C21 := make([][]float64, newRows)
	C22 := make([][]float64, newRows)

	for i := 0; i < newRows; i++ {
		A11[i] = make([]float64, newCols)
		A12[i] = make([]float64, newCols)
		A21[i] = make([]float64, newCols)
		A22[i] = make([]float64, newCols)

		B11[i] = make([]float64, newCols)
		B12[i] = make([]float64, newCols)
		B21[i] = make([]float64, newCols)
		B22[i] = make([]float64, newCols)

		C11[i] = make([]float64, newCols)
		C12[i] = make([]float64, newCols)
		C21[i] = make([]float64, newCols)
		C22[i] = make([]float64, newCols)

		for j := 0; j < newCols; j++ {
			A11[i][j] = mat1[i][j]
			A12[i][j] = mat1[i][j+newCols]
			A21[i][j] = mat1[i+newRows][j]
			A22[i][j] = mat1[i+newRows][j+newCols]

			B11[i][j] = mat2[i][j]
			B12[i][j] = mat2[i][j+newCols]
			B21[i][j] = mat2[i+newRows][j]
			B22[i][j] = mat2[i+newRows][j+newCols]
		}
	}

	// Recursive steps for multiplication
	P1 := multiplyMatrices(A11, sub(B12, B22))
	P2 := multiplyMatrices(add(A11, A12), B22)
	P3 := multiplyMatrices(add(A21, A22), B11)
	P4 := multiplyMatrices(A22, sub(B21, B11))
	P5 := multiplyMatrices(add(A11, A22), add(B11, B22))
	P6 := multiplyMatrices(sub(A12, A22), add(B21, B22))
	P7 := multiplyMatrices(sub(A11, A21), add(B11, B12))

	// Computing result quadrants
	C11 = add(sub(add(P5, P4), P2), P6)
	C12 = add(P1, P2)
	C21 = add(P3, P4)
	C22 = sub(sub(add(P1, P5), P3), P7)

	// Combining result quadrants into the final result matrix
	result := make([][]float64, rows1)
	for i := 0; i < newRows; i++ {
		result[i] = make([]float64, cols2)
		result[i+newRows] = make([]float64, cols2)
		for j := 0; j < newCols; j++ {
			result[i][j] = C11[i][j]
			result[i][j+newCols] = C12[i][j]
			result[i+newRows][j] = C21[i][j]
			result[i+newRows][j+newCols] = C22[i][j]
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

	if m != n || n != l || m&(m-1) != 0 {
		log.Fatal("Must be square matrices of size power of 2")
	}

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
