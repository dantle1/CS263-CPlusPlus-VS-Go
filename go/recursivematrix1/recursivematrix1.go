package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random matrix
func generateRandomMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(2) + 1 // Generate random number between 1 and 2
		}
	}
	return matrix
}

// Function to add two matrices
func addMatrices(mat1, mat2 [][]int) [][]int {
	rows := len(mat1)
	cols := len(mat1[0])
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
		for j := range result[i] {
			result[i][j] = mat1[i][j] + mat2[i][j]
		}
	}
	return result
}

// Function to subtract two matrices
func subtractMatrices(mat1, mat2 [][]int) [][]int {
	rows := len(mat1)
	cols := len(mat1[0])
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
		for j := range result[i] {
			result[i][j] = mat1[i][j] - mat2[i][j]
		}
	}
	return result
}

// Function to multiply two matrices using Strassen's algorithm
func multiplyMatrices(mat1, mat2 [][]int) [][]int {
	rows1 := len(mat1)
	cols1 := len(mat1[0])
	cols2 := len(mat2[0])

	// Base case: if matrices are 1x1, perform simple multiplication
	if rows1 == 1 && cols1 == 1 && cols2 == 1 {
		result := make([][]int, 1)
		result[0] = make([]int, 1)
		result[0][0] = mat1[0][0] * mat2[0][0]
		return result
	}

	// Split matrices into quadrants
	newRows := rows1 / 2
	newCols := cols1 / 2

	A11 := make([][]int, newRows)
	A12 := make([][]int, newRows)
	A21 := make([][]int, newRows)
	A22 := make([][]int, newRows)

	B11 := make([][]int, newRows)
	B12 := make([][]int, newRows)
	B21 := make([][]int, newRows)
	B22 := make([][]int, newRows)

	C11 := make([][]int, newRows)
	C12 := make([][]int, newRows)
	C21 := make([][]int, newRows)
	C22 := make([][]int, newRows)

	for i := 0; i < newRows; i++ {
		A11[i] = make([]int, newCols)
		A12[i] = make([]int, newCols)
		A21[i] = make([]int, newCols)
		A22[i] = make([]int, newCols)

		B11[i] = make([]int, newCols)
		B12[i] = make([]int, newCols)
		B21[i] = make([]int, newCols)
		B22[i] = make([]int, newCols)

		C11[i] = make([]int, newCols)
		C12[i] = make([]int, newCols)
		C21[i] = make([]int, newCols)
		C22[i] = make([]int, newCols)

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
	P1 := multiplyMatrices(A11, subtractMatrices(B12, B22))
	P2 := multiplyMatrices(addMatrices(A11, A12), B22)
	P3 := multiplyMatrices(addMatrices(A21, A22), B11)
	P4 := multiplyMatrices(A22, subtractMatrices(B21, B11))
	P5 := multiplyMatrices(addMatrices(A11, A22), addMatrices(B11, B22))
	P6 := multiplyMatrices(subtractMatrices(A12, A22), addMatrices(B21, B22))
	P7 := multiplyMatrices(subtractMatrices(A11, A21), addMatrices(B11, B12))

	// Computing result quadrants
	C11 = addMatrices(subtractMatrices(addMatrices(P5, P4), P2), P6)
	C12 = addMatrices(P1, P2)
	C21 = addMatrices(P3, P4)
	C22 = subtractMatrices(subtractMatrices(addMatrices(P1, P5), P3), P7)

	// Combining result quadrants into the final result matrix
	result := make([][]int, rows1)
	for i := 0; i < newRows; i++ {
		result[i] = make([]int, cols2)
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
	rand.Seed(time.Now().Unix()) // Seed for random number generator using current time
	sizes := []int{1024, 2048, 4096}
	for _, size := range sizes {
		// Generate random matrices
		matrix1 := generateRandomMatrix(size, size)
		matrix2 := generateRandomMatrix(size, size)
		start := time.Now()
		// Multiply matrices using Strassen's algorithm
		_ = multiplyMatrices(matrix1, matrix2)
		elapsedTime := time.Since(start).Seconds()
		fmt.Printf("Elapsed time for %d x %d matrix multiplication: %.2f seconds.\n", size, size, elapsedTime)
	}
}