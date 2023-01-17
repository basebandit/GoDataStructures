package zerosum

// Given an integer array, check if it contains a subarray having zero-sum.
//
// For example,
//
// Input:  { 3, 4, -7, 3, 1, 3, 1, -4, -2, -2 }
//
// Output: Subarray with zero-sum exists
//
// The subarrays with a sum of 0 are:
//
// { 3, 4, -7 }
// { 4, -7, 3 }
// { -7, 3, 1, 3 }
// { 3, 1, -4 }
// { 3, 1, 3, 1, -4, -2, -2 }
// { 3, 4, -7, 3, 1, 3, 1, -4, -2, -2 }
func zeroSumSubarray(n []int) bool {
	set := map[int]struct{}{}
	var sum int
	for _, i := range n {
		//sum of elements so far
		sum += i

		// if the sum is seen before, we have found a sublist with zero sum
		if _, ok := set[sum]; ok {
			return true
		}
		set[sum] = struct{}{}
	}
	return false
}
