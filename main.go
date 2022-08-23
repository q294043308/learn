package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var A []int
	var B []int
	B = make([]int, 0)
	a, _ := json.Marshal(A)
	b, _ := json.Marshal(B)
	fmt.Println(string(a))
	fmt.Println(string(b))
}
