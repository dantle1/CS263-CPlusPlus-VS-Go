BIN=./bin
PROF=./prof
# This is how we change the Go compiler we are using
GO=go # 1.19
# GO=go1.18
#GO=go1.22.0

PROG=havlakgo1 havlakgo2 havlakgo3 havlakgo4 revcmpgo gcgo1 gcgo2 iterativematrix1

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

# naive iterative matrix mult
iterativematrix1:
	cd iterativematrix1 && $(GO) build -o ../bin/iterativematrix1

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

.PHONY: all havlakgo revcmpgo iterativematrix1 clean