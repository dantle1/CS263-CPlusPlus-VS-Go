package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"time"
	"flag"
)

var (
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
	infile     = flag.String("infile", "../data/array/array.in", "Matrices to multiply")
)

func partition(arr []int64, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int64, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
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
	quickSort(arr, 0, n-1)
	//end := time.Now()

	for _, element := range arr {
		fmt.Println(element)
	}

	//fmt.Println("Mergesort time for", n, "elements:", end.Sub(start))
}
