package main

import (
	"fmt"
)

func main() {
	sample := []int{3, 3, 2, 1, 3}
	fmt.Println(equalizeArray(sample))
	sample2 := []int{1, 2, 2, 3}
	fmt.Println(equalizeArray(sample2))
	sample3 := []int{1, 2, 3, 1, 2, 3, 3, 3}
	fmt.Println(equalizeArray(sample3))
}

//Find the minimum number of deletions to equalize an array
//e.g. []int{1, 2, 2, 3} = 2 deletions we remove 1 and 3
//e.g []int{3, 3, 2, 1, 3} = 2 deletions we remove 2 and 1
func equalizeArray(n []int) int { //O(n)

	dict := make(map[int]int)

	for i := 0; i < len(n); i++ {
		dict[n[i]]++
	}

	var maxFreq int

	//sort this dictionary by value
	for _, v := range dict {
		if v > maxFreq {
			maxFreq = v
		}
	}

	return int(len(n)) - maxFreq
}
