package permutation

import (
	"strings"
)

func swap(str1, str2 *string) {
	*str1, *str2 = *str2, *str1
}

func fact(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

// Heap's Algorithm for generating all permutations of n objects
func Heaps(out chan []string, n int) {
	elementSetCh := make(chan []string)
	go GenerateElementSet(elementSetCh, n)
	elementSet := <-elementSetCh

	var recursiveGenerate func([]string, int, []string)
	permutations := make([]string, fact(n))
	last := 0
	// var permutations []string
	recursiveGenerate = func(previousIteration []string, n int, elements []string) {
		if n == 1 {
			// permutations = append(permutations, strings.Join(elements, ""))
			permutations[last] = strings.Join(elements, "")
			last++
		} else {
			for i := 0; i < n; i++ {
				recursiveGenerate(previousIteration, n-1, elements)
				if n%2 == 1 {
					swap(&elements[i], &elements[n-1])
					// tmp := elements[i]
					// elements[i] = elements[n-1]
					// elements[n-1] = tmp
				} else {
					swap(&elements[0], &elements[n-1])
					// tmp := elements[0]
					// elements[0] = elements[n-1]
					// elements[n-1] = tmp
				}
			}
		}
	}
	recursiveGenerate(permutations, n, elementSet)
	out <- permutations
}

func GenerateElementSet(out chan []string, n int) {
	elementSet := make([]string, n)
	for i := range elementSet {
		elementSet[i] = string(rune(i + 49)) // Adjust this if you want to change your charset
	}
	out <- elementSet
}
