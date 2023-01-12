package selectionsort

// selectionSort searches for the smallest element and swaps it with the current index,
// and it repeats with each iteration increasing the index.It does this inplace. No need to allocate a new array(memory)
func selectionSort(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}
