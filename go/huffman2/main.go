package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/dantle1/CS263-CPlusPlus-VS-Go/go/huffman2/huffman"
)

var (
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to this file")
	infile     = flag.String("infile", "../data/string/string.in", "Input file")
)

func main() {
	flag.Parse()

	file, _ := os.Open(*infile)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	message := ""

	for scanner.Scan() {
		message += scanner.Text()
	}

	tree, _ := huffman.HuffTree(huffman.SymbolCountOrd(message))

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	codes := make(map[rune][]bool)
	huffman.HuffEncoding(tree, nil, codes)
	messageCoded := huffman.HuffEncode(codes, message)
	var b bytes.Buffer
	messageHuffDecoded := huffman.HuffDecode(tree, tree, messageCoded, b, "", 1024)
	if messageHuffDecoded != message {
		fmt.Printf("got: %q\nbut expected: %q", messageHuffDecoded, message)
	}
}