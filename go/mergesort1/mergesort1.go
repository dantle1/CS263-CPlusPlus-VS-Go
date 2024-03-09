package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"flag"
	//"time"
)

var (
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
	infile     = flag.String("infile", "../data/array/array.in", "Matrices to multiply")
)

func merge(arr []int64, left, mid, right int) {
	subArrayOne := mid - left + 1
	subArrayTwo := right - mid

	leftArray := make([]int64, subArrayOne)
	rightArray := make([]int64, subArrayTwo)

	for i := 0; i < subArrayOne; i++ {
		leftArray[i] = arr[left+i]
	}
	for j := 0; j < subArrayTwo; j++ {
		rightArray[j] = arr[mid+1+j]
	}

	indexOfSubArrayOne := 0
	indexOfSubArrayTwo := 0
	indexOfMergedArray := left

	for indexOfSubArrayOne < subArrayOne && indexOfSubArrayTwo < subArrayTwo {
		if leftArray[indexOfSubArrayOne] <= rightArray[indexOfSubArrayTwo] {
			arr[indexOfMergedArray] = leftArray[indexOfSubArrayOne]
			indexOfSubArrayOne++
		} else {
			arr[indexOfMergedArray] = rightArray[indexOfSubArrayTwo]
			indexOfSubArrayTwo++
		}
		indexOfMergedArray++
	}

	for indexOfSubArrayOne < subArrayOne {
		arr[indexOfMergedArray] = leftArray[indexOfSubArrayOne]
		indexOfSubArrayOne++
		indexOfMergedArray++
	}

	for indexOfSubArrayTwo < subArrayTwo {
		arr[indexOfMergedArray] = rightArray[indexOfSubArrayTwo]
		indexOfSubArrayTwo++
		indexOfMergedArray++
	}
}

func mergeSort(arr []int64, begin, end int) {
	if begin >= end {
		return
	}

	mid := begin + (end-begin)/2
	mergeSort(arr, begin, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, begin, mid, end)
}

func main() {
	flag.Parse()
	
	file, _ := os.Open(*infile)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	arr_length := strings.Fields(scanner.Text())

	n, _ := strconv.Atoi(arr_length[0])

	arr := make([]int64, n)
	scanner.Scan()
	line := scanner.Text()
	iss := strings.Fields(line)
	
	for i, numStr := range iss{
		arr[i], _ = strconv.ParseInt(numStr, 10, 64)
	}

	//start := time.Now()
	mergeSort(arr, 0, n-1)
	//end := time.Now()

	for _, element := range arr {
		fmt.Println(element)
	}

	//fmt.Println("Mergesort time for", n, "elements:", end.Sub(start))
}
