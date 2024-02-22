BIN=./bin

PROG=havlakgo revcmpgo
LIST=$(addprefix $(BIN)/, $(PROG))

all: $(LIST)

$(BIN)/havlakgo:
	cd havlak && go build -o ../bin/havlakgo

$(BIN)/revcmpgo:
	cd reverse-complement && go build -o ../bin/revcmpgo

clean:
	rm -f $(BIN)/*

.PHONY: all havlakgo revcmpgo clean