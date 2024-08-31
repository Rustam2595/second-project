package main

import (
	"fmt"
)

func findMaxValueIndex(data map[int]int) int {
	var n, r int
	for ind, i := range data {
		fmt.Println(ind, i, data[ind])
		if n < i {
			n = i
			r = ind
		}
	}
	return r
}
func main() {
	mapa := map[int]int{
		1: 10, 2: 20, 3: 30, 4: 5,
	}
	fmt.Println(findMaxValueIndex(mapa))
}
