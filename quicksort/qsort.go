// Copyright  Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package quicksort

import (
	"math/rand"
)

// qsortWithRandomPivot worst case is O(n2) but on average it is O(nlogn)
func qsortWithRandomPivot(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left := 0
	right := len(a) - 1
	pivot := rand.Int() % len(a)
	// move the pivot to the edge of the array
	a[pivot], a[right] = a[right], a[pivot]

	// iterate
	for i, _ := range a {
		// check if the element is less than the pivot(moved to the right now, remember that)
		if a[i] < a[right] {
			// if smaller move the element to the left of the pivot
			a[left], a[i] = a[i], a[left]
			// increment our left pointer (index)
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	qsortWithRandomPivot(a[:left])
	qsortWithRandomPivot(a[left+1:])

	return a
}

// qsort worst case is O(n2) but on average it is O(nlogn)
func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left := 0
	pivot := len(a) - 1

	for i, _ := range a {
		// check if the element is less than the pivot
		if a[i] < a[pivot] {
			// if smaller swap the left element with the current element
			a[left], a[i] = a[i], a[left]
			// increment our left pointer (index)
			left++
		}
	}

	a[left], a[pivot] = a[pivot], a[left]
	qsortWithRandomPivot(a[:left])
	qsortWithRandomPivot(a[left+1:])

	return a
}
