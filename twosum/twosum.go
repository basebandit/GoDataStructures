// Copyright  Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package twosum

import "sort"

// o(n2)
func twoSum(n []int, sum int) []int {
	for i := 0; i < len(n); i++ {
		for j := i + 1; i < len(n); j++ {
			if n[i]+n[j] == sum {
				return []int{n[i], n[j]}
			}
		}
	}

	return []int{}
}

// O(nlogn) because of sorting
func twoSumOptimized(n []int, sum int) []int {
	sort.Ints(n)

	left := 0
	right := len(n) - 1

	for left < right {
		if n[left]+n[right] == sum {
			return []int{n[left], n[right]}
		}
		if n[right]+n[left] > sum {
			right--
		} else if n[right]+n[left] < sum {
			left++
		}

	}
	return []int{}
}
