package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"github.com/dantle1/CS263-CPlusPlus-VS-Go/go/huffman/huffman"
)

var (
	memprofile  = flag.String("memprofile", "", "write memory profile to this file")
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to this file")
	infile 		= flag.String("infile", "../data/string/string.in", "Input file")
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

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	tree, _ := huffman.HuffTree(huffman.SymbolCountOrd(message))
	codes := make(map[rune][]bool)
	huffman.HuffEncoding(tree, nil, codes)
	messageCoded := huffman.HuffEncode(codes, message)
	messageHuffDecoded := huffman.HuffDecode(tree, tree, messageCoded, "")
	if messageHuffDecoded != message {
		fmt.Printf("got: %q\nbut expected: %q", messageHuffDecoded, message)
	}
}