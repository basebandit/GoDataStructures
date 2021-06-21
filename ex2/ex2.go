package main

import (
	"fmt"
)

//Find the pair of numbers whose sum is 8
//[1,2,3,9] Sum = 8 No - has no pair
//[1,2,4,4] Sum = 8 yes - has pair
//Find the indices of the pair of numbers whose sum is 8
//[1,2,3,9] Sum = 8 No - has no pair
//[1,2,4,4] Sum = 8 yes - has pair
func main() {
	dataset := [][]int{{1, 2, 4, 4}, {1, 2, 3, 9}}
	for _, data := range dataset {
		if has, pair := hasPairWithSumWithSortedInput(data, 8); has {
			fmt.Printf("input: %v, pair: %v, ", data, pair)
		}
		if pairIdx := findIndexOfPairWithSum(data, 8); pairIdx != nil {
			fmt.Println("indices: ", pairIdx)
		}
	}

	dataset2 := [][]int{{2, 4, 1, 4}, {2, 1, 3, 9}}
	for _, data := range dataset2 {
		if has, pair := hasPairWithSumWithUnsortedInput(data, 8); has {
			fmt.Printf("input: %v, pair: %v, ", data, pair)
		}

		if pairIdx := findIndexOfPairWithSum(data, 8); pairIdx != nil {
			fmt.Println("indices: ", pairIdx)
		}
	}
}

//Simplest approach would be using the brute force approach where we use nested for loop,
// i and j loops, j starts at (i + 1) and check all through if any of the pairs sum up to 8,
//but then this approach would have a quadratic runtime complexity O(n^2)
//Instead think of another approach that uses a linear complexixty O(n)
//Since the input is sorted we go right away to our solution.
func hasPairWithSumWithSortedInput(sortedinput []int, sum int) (bool, []int) { //O(n)

	var low int //0
	var high int = len(sortedinput) - 1
	input := sortedinput
	for low < high {
		if input[low]+input[high] == sum {
			return true, []int{input[low], input[high]}
		}

		if input[low]+input[high] > sum {
			high = len(input) - 2
		}

		if input[low]+input[high] < sum {
			low++
		}
	}

	return false, nil
}

//What if we were working with unsorted data, ok so we are going to ask ourselves
// a question let's say while we are at the current value lets say data[0] = 1, then the question
// would be have we seen a 7 a complement of 1 which is the result of (8 -1) before, if we have then
//there we have our answer, if not then we are going to add the 1 in our hashset, but before doing so we
//check if we already have a one before and we do the same for the other elements
//{1,2,4,4}
func hasPairWithSumWithUnsortedInput(input []int, sum int) (bool, []int) {
	//our hashset in go is implemented using a map[int]int
	//the set will hold our element and its index/position in the input array

	set := make(map[int]int)

	for _, val := range input {
		comp := sum - val
		if is := set[comp]; is != 0 {
			return true, []int{comp, val}
		}
		set[val] = 1
	}

	return false, nil
}

//input is not sorted. This solution uses O(n)
func findIndexOfPairWithSum(input []int, sum int) []int {
	set := make(map[int]int)

	//in our map the key is the input element and the value is the element's index
	for key, val := range input {
		comp := sum - val
		if idx, ok := set[comp]; ok { //if we have seen the complement of the current value then we have our pairs- the index of the complement(idx), the current value's index (key)
			return []int{idx, key} //index of complement,index of current val
		}
		set[val] = key
	}
	return nil
}
