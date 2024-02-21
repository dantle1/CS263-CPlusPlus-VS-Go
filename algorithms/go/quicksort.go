package main

import "fmt"

func partition(arr []int, start, end int) int {
	pivot := arr[start]
	count := 0
	for i := start + 1; i <= end; i++ {
		if arr[i] <= pivot {
			count++
		}
	}

	pivotIndex := start + count
	arr[start], arr[pivotIndex] = arr[pivotIndex], arr[start]

	i, j := start, end
	for i < pivotIndex && j > pivotIndex {
		for arr[i] <= pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i < pivotIndex && j > pivotIndex {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	return pivotIndex
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	p := partition(arr, start, end)
	quickSort(arr, start, p-1)
	quickSort(arr, p+1, end)
}

func main() {
	arr := []int{9, 3, 4, 2, 1, 8}
	n := len(arr)

	quickSort(arr, 0, n-1)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
