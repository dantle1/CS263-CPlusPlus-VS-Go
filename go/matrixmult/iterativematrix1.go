package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random matrix
func generateRandomMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Intn(2) + 1 // Generate random number between 1 and 2
		}
	}
	return matrix
}

// Function to multiply two matrices
func multiplyMatrices(mat1, mat2 [][]int) [][]int {
	rows1 := len(mat1)
	cols1 := len(mat1[0])
	cols2 := len(mat2[0])

	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
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
	rand.Seed(time.Now().UnixNano()) // Seed for random number generator using current time
	sizes := []int{1024, 2048, 4096}
	for _, size := range sizes {
		matrix1 := generateRandomMatrix(size, size)
		matrix2 := generateRandomMatrix(size, size)

		start := time.Now()
		// Multiply matrices
		_ = multiplyMatrices(matrix1, matrix2)
		elapsedTime := time.Since(start)

		fmt.Printf("Elapsed time for %dx%d matrix multiplication: %v\n", size, size, elapsedTime)
	}
}
