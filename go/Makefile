BIN=./bin
PROF=./prof
# This is how we change the Go compiler we are using
# GO=go # 1.19
# GO=go1.21.0
# GO=go1.18
GO=go1.22.0

PROG=havlakgo1 havlakgo2 havlakgo3 havlakgo4 revcmpgo gcgo1 gcgo2 recmat1 itmat1 recmatpgo itmatpgo recmat2 graph huffman1 huffman2 permutation1 permutation2 huffman0  md mdpgo primenum iterfibonacci recfibonacci quicksort1 mergesort1

all: $(PROG)

# Initial implementation
havlakgo1:
	cd havlak1 && $(GO) build -o ../bin/havlakgo1

# Optimized with slicing, as opposed to using maps
havlakgo2:
	cd havlak2 && $(GO) build -o ../bin/havlakgo2

# Optimized with additional simpler data structures (again), changing nonBackPreds to a slice
# (changed backPreds as well! Used about half as much space on line 176 in havlakloopfinder.go!!)
havlakgo3:
	cd havlak3 && $(GO) build -o ../bin/havlakgo3

# Used caching structure to alleviate some of the cpu load from the garbage collector
havlakgo4:
	cd havlak4 && $(GO) build -o ../bin/havlakgo4

revcmpgo:
	cd reverse-complement && $(GO) build -o ../bin/revcmpgo

# Modified https://www.linkedin.com/pulse/understanding-tricolor-mark-sweep-algorithm-gos-sourav-choudhary/
# The Tricolor Mark and Sweep garbage collector algorithm in Go's gc
gcgo1:
	cd gc1 && $(GO) build -o ../bin/gcgo1

# Now multithreaded 
gcgo2:
	cd gc2 && $(GO) build -o ../bin/gcgo2

# Recursive and iterative matrix multiplication 
itmat1:
	cd itmat1 && $(GO) build -o ../bin/itmat1 

recmat1:
	cd recmat1 && $(GO) build -o ../bin/recmat1 

md:
	cd markdown && $(GO) build -o ../bin/md 

mdpgo:
	@if [ "$(GO)" != "go1.22.0" ]; then \
		echo "Error: change Go compiler to version 1.22.0"; \
		exit 1; \
	fi
	cd markdown && $(GO) build -pgo=auto -o ../bin/mdpgo
	
recmat2:
	cd recmat2 && $(GO) build -o ../bin/recmat2 

# quicksort and mergesort 
quicksort1:
	cd quicksort1 && $(GO) build -o ../bin/quicksort1

mergesort1:
	cd mergesort1 && $(GO) build -o ../bin/mergesort1

# primenum
primenum:
	cd primenum && $(GO) build -o ../bin/primenum

# fibonacci
iterfibonacci:
	cd iterfibonacci && $(GO) build -o ../bin/iterfibonacci

recfibonacci:
	cd recfibonacci && $(GO) build -o ../bin/recfibonacci

# profiler guide optimization

gcpgo:
	cd gc1 && $(GO) build -pgo=auto -o ../bin/gcpgo

itmatpgo:
	@if [ "$(GO)" != "go1.22.0" ]; then \
			echo "Error: change Go compiler to version 1.22.0"; \
			exit 1; \
	fi
	cd itmat1 && $(GO) build -pgo=auto -o ../bin/itmatpgo

recmatpgo:
	@if [ "$(GO)" != "go1.22.0" ]; then \
			echo "Error: change Go compiler to version 1.22.0"; \
			exit 1; \
	fi
	cd recmat1 && $(GO) build -pgo=auto -o ../bin/recmatpgo

graph:
	cd graph && $(GO) build -o ../bin/graph 

huffman0:
	cd huffman0 && $(GO) build -o ../bin/huffman0

huffman1:
	cd huffman1 && $(GO) build -o ../bin/huffman1

huffman2:
	cd huffman2 && $(GO) build -o ../bin/huffman2

permutation1:
	cd permutation1 && $(GO) build -o ../bin/permutation1

permutation2:
	cd permutation2 && $(GO) build -o ../bin/permutation2

recmat2pgo:
	@if [ "$(GO)" != "go1.22.0" ]; then \
			echo "Error: change Go compiler to version 1.22.0"; \
			exit 1; \
	fi
	cd recmat2 && $(GO) build -pgo=auto -o ../bin/recmat2pgo

%.cprof: 
	$(MAKE) $*
	$(BIN)/$* -cpuprofile=$(PROF)/$@

%.mprof:
	$(MAKE) $*
	$(BIN)/$* -memprofile=$(PROF)/$@

# These give us interactive mode, but there are also other interesting options we can add with flags
# For example --nodefraction=0.1 will filter out noise 
cprof-%: 
	$(MAKE) $*.cprof
	$(GO) tool pprof $(PROF)/$*.cprof

mprof-%: 
	$(MAKE) $*.mprof 
	$(GO) tool pprof $(PROF)/$*.mprof

xtime-%:
	$(MAKE) $*
	../xtime $(BIN)/$*

clean:
	rm -f $(BIN)/*

clean-prof:
	rm -f $(PROF)/*

.PHONY: all havlakgo revcmpgo itmat1 recmat1 recmat2 recmatpgo itmatpgo graph huffman1 huffman2 permutation1 permutation2 huffman0 md mdpgo quicksort1 mergesort1 primenum iterfibonacci recfibonacci clean