// Copyright  Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package binarysearch

// binarySearch finds the target integer in the array and returns its index,otherwise it returns -1
func binarySearch(arr []int, target int) int {
	var mid int
	low := 0
	high := len(arr) - 1
	for low <= high { //while we have not yet narrowed down to one element
		mid = (low + high) / 2
		if arr[mid] == target { // check if the middle element is our target
			return mid
		}
		if arr[mid] < target { // if our middle element is lower than our target, we adjust our mid by incrementing by one
			// effectively discarding the lower half of our array (search space)
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
