package main

import "fmt"

//Given 2 arrays, find whether they contain any common items.
//{'a', 'b', 'c', 'x'}
//{'z', 'y', 'x'}
func main() {
	input1 := []int32{'a', 'b', 'c', 'x'}
	input2 := []int32{'z', 'y', 'x'}
	fmt.Println(commonItems(input1, input2))
	fmt.Println(commonItems2(input1, input2))
}

//the first solution that comes to mind is the naive brute force approach with a quadratic rumtime complexity O(n^2)
func commonItems(arr1, arr2 []rune) bool { //O(n^2)
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}
	return false
}

// a rune is an alias of an int32. in golang characters are represented in memory as int32 values. rune is for convenience.
//lets improve our algorithm's runtime, how about we use a map to store the first array and lookup the elements of the second array in the map
//remember to check for errors wherever necessary. or sanity checks like the length of the inputs.
func commonItems2(arr1, arr2 []rune) bool { //O(n)
	dict := make(map[rune]bool) //lets use a bool (which is only a byte) as our value..We are also being mindful of space complexity :)
	for i := 0; i < len(arr1); i++ {
		if ok := dict[arr1[i]]; !ok {
			dict[arr1[i]] = true
		}
	}

	//lookup
	for j := 0; j < len(arr2); j++ {
		if ok := dict[arr2[j]]; ok {
			return ok
		}
	}

	return false
}
