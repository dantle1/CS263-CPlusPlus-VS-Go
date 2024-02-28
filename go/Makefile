BIN=./bin
PROF=./prof

PROG=havlakgo1 havlakgo2 revcmpgo

all: $(PROG)

# Initial implementation
havlakgo1:
	cd havlak1 && go build -o ../bin/havlakgo1

# Optimized with slicing, as opposed to using maps
havlakgo2:
	cd havlak2 && go build -o ../bin/havlakgo2

# Optimized with additional simpler data structures (again), changing nonBackPreds to a slice
# (changed backPreds as well! Used about half as much space on line 176 in havlakloopfinder.go!!)
havlakgo3:
	cd havlak3 && go build -o ../bin/havlakgo3

# Used caching structure to alleviate some of the cpu load from the garbage collector
havlakgo4:
	cd havlak4 && go build -o ../bin/havlakgo4

revcmpgo:
	cd reverse-complement && go build -o ../bin/revcmpgo

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
	go tool pprof $(PROF)/$*.cprof

mprof-%: 
	$(MAKE) $*.mprof 
	go tool pprof $(PROF)/$*.mprof

xtime-%:
	$(MAKE) $*
	../xtime $(BIN)/$*

clean:
	rm -f $(BIN)/*
	rm -f $(PROF)/*

.PHONY: all havlakgo revcmpgo clean