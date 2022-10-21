package main

import (
	"fmt"
	"learn/function"
)

func main() {
	board := [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}
	words := []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}
	fmt.Println(function.FindWords(board, words))

}
