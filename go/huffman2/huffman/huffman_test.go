// huffmancoding_test.go
// description: Tests the huffman, encoding and decoding algorithms of huffmancoding.go.
// author(s) [pedromsrocha](https://github.com/pedromsrocha)
// see huffmancoding.go

package huffman_test

import (
	"bytes"
	"testing"

	"github.com/dantle1/CS263-CPlusPlus-VS-Go/go/huffman2/huffman"
)

func TestHuffman(t *testing.T) {
	messages := []string{
		"hello world \U0001F600",
		"colorless green ideas sleep furiously",
		"the quick brown fox jumps over the lazy dog",
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit,
		sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
		Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
		aliquip ex ea commodo consequat.`,
	}

	for _, message := range messages {
		t.Run("huffman: "+message, func(t *testing.T) {
			tree, _ := huffman.HuffTree(huffman.SymbolCountOrd(message))
			codes := make(map[rune][]bool)
			huffman.HuffEncoding(tree, nil, codes)
			messageCoded := huffman.HuffEncode(codes, message)
			var b bytes.Buffer
			messageHuffDecoded := huffman.HuffDecode(tree, tree, messageCoded, b, "", 1024)
			if messageHuffDecoded != message {
				t.Errorf("got: %q\nbut expected: %q", messageHuffDecoded, message)

			}
		})
	}
}