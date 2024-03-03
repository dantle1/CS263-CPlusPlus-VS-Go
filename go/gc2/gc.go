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
	"sync"
)

type Object struct {
	Marked  int // Color of Object
	Left    int
	Right   int
	Present bool
}

func NewObject() *Object {
	return &Object{
		Marked:  white, // Default value for Marked
		Left:    -1,    // Default value for Left
		Right:   -1,    // Default value for Right
		Present: true,  // Default value for Present
	}
}

const (
	white = iota // Unmarked and unreachable
	gray
	black // Marked and reachable
)

var (
	memprofile  = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to this file")
	heapprop    = flag.Float64("heapprop", 0.85, "Proportion of heap capacity allocated")
	numcycles   = flag.Int("numcycles", 1000000, "Number of Garbage collection cycles")
	rootsetsize = flag.Int("rootsetsize", 200, "Number of variables in the root set")
	infile      = flag.String("infile", "../python/reg-graph-data/big-graph.txt", "Graph to model heap structure")
	heap        []*Object // Global heap
	rootSet     []int     // Global root set of objects (program variables)
	mu          sync.Mutex
)

// Initializes the heap of size `size`
func initHeap(size int) {
	heap = make([]*Object, size)
	for i := 0; i < size; i++ {
		heap[i] = NewObject()
	}
}

// Unmarks all of the objects (colors them white)
func resetColors() {
	mu.Lock()
	defer mu.Unlock()

	for i := range heap {
		heap[i].Marked = white
		heap[i].Present = true
	}
}

// Marks anything directly reachable from rootset as gray
func initialMark() {
	mu.Lock()
	defer mu.Unlock()

	for _, i := range rootSet {
		obj := heap[i]
		obj.Marked = gray
	}
}

// Recursively mark reachable objects using DFS
func mark(i int) {
	obj := heap[i]
	if obj.Marked == black {
		return
	}

	// Mark the current object as reachable.
	obj.Marked = gray

	// Recursively mark its references.
	if heap[obj.Left].Marked != gray {
		mark(obj.Left)
	}
	if heap[obj.Right].Marked != gray {
		mark(obj.Right)
	}

	// After marking references, mark the object as black. (postorder)
	obj.Marked = black
}

// Marks objects as gray concurrently
func concurrentMark() {
	var wg sync.WaitGroup
	mu.Lock()

	for _, i := range rootSet {
		obj := heap[i]
		if obj.Marked == gray {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				mark(j)
			}(i)
		}
	}
	mu.Unlock()
	wg.Wait()
}

func sweep() {
	mu.Lock()
	defer mu.Unlock()

	for i := range heap {
		obj := heap[i]
		if obj != nil && obj.Marked == white {
			// Free the object
			heap[i].Present = false
		}
	}
}

// A complete tricolor mark & sweep gc cycle
func gcCycle() {
	resetColors()

	initialMark()

	concurrentMark()

	sweep()

	// Reclaim() -- handled by the Go runtime, not the gc typically
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	file, _ := os.Open(*infile)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	initHeap(int(float64(n) / *heapprop))

	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Fields(line)
		u, _ := strconv.Atoi(pair[0])
		v, _ := strconv.Atoi(pair[1])
		if heap[u].Left == -1 {
			heap[u].Left = v
		} else {
			heap[u].Right = v
		}
	}

	// Set the roots evenly between 0,...,n - 1
	interval := n / *rootsetsize
	for i := 0; i < *rootsetsize; i++ {
		rootSet = append(rootSet, interval*i)
	}

	for i := 0; i < *numcycles; i++ {
		gcCycle()
		// fmt.Printf("After GC cycle %d:\n", i+1)
		// printObjectStatus()
	}
}

// printObjectStatus prints the status of each object (black, gray, or white).
func printObjectStatus() {
	mu.Lock()
	defer mu.Unlock()

	for i, obj := range heap {
		if !obj.Present {
			fmt.Printf("Object %d: Freed\n", i)
		} else if obj.Marked == black {
			fmt.Printf("Object %d: Marked (Black)\n", i)
		}
	}
}
