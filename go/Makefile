BIN=./bin
PROF=./prof

PROG=havlakgo revcmpgo
LIST=$(addprefix $(BIN)/, $(PROG))

all: $(LIST)

havlakgo:
	cd havlak && go build -o ../bin/havlakgo

revcmpgo:
	cd reverse-complement && go build -o ../bin/revcmpgo

%.cprof: 
	$(MAKE) $*
	$(BIN)/$* -cpuprofile=$(PROF)/$@

%.mprof:
	$(MAKE) $*
	$(BIN)/$* -memprofile=$(PROF)/$@

cprof-%: 
	$(MAKE) $*.cprof
	go tool pprof $(PROF)/$*.cprof

mprof-%: 
	$(MAKE) $*.mprof 
	go tool pprof $(PROF)/$*.mprof

clean:
	rm -f $(BIN)/*
	rm -f $(PROF)/*

.PHONY: all havlakgo revcmpgo clean