package main

import (
	"learn/function"
)

var aaa int64 = 6

func main() {
	r := function.CanFinish(5, [][]int{{1, 0}, {0, 1}, {3, 1}, {3, 2}})
	println(r)
}
