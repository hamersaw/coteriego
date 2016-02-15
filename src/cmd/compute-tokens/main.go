package main

import (
	"flag"
	"fmt"
	"math"
)

var tokenCount int

func init() {
	flag.IntVar(&tokenCount, "tokenCount", 8, "Number of tokens to create")
}

func main() {
	flag.Parse()

	delta := int64(math.MaxUint64 / uint64(tokenCount))
	token := int64(0)
	tokens := []int64{}
	for i := 0; i < tokenCount; i++ {
		tokens = append(tokens, token)
		token += delta
	}

	fmt.Printf("tokens: %v\n", tokens)
}
