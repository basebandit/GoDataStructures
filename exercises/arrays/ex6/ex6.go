package main

import "fmt"

func main() {
	sample := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	count := 10
	fmt.Println(findMissingNum(sample, count))
}

//find the missing number in an array. Things to consider is the array sorted. If not the approach might differ.
//we use this nifty mathematical formulae to find the expectedTotalSum n(n+1)/2
///then we find the current sum and subtract it from the expected sum
//Our solution does not matter if the array is sorted or whether it contains duplicates.
func findMissingNum(arr []int, count int) int { //O(n)
	var currentSum int

	expectedSum := count * (count + 1) / 2

	for i := range arr {
		currentSum += arr[i]
	}
	return expectedSum - currentSum
}
