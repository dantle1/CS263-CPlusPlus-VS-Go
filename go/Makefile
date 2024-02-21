# Updated Makefile for standard Go toolchain

BINARY := havlak1
SOURCES := basicblock.go lsg.go looptesterapp.go havlakloopfinder.go

all: $(BINARY)

$(BINARY): $(SOURCES)
	go build -o $(BINARY) $(SOURCES)

run: $(BINARY)
	./$(BINARY)

clean:
	rm -f $(BINARY)