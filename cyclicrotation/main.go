package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums[3:])
	cyclicRotation(nums, 2)
}

// cyclicRotation rotates array n in k cycles.
func cyclicRotation(n []int, k int) {
	nHeader := (*reflect.SliceHeader)(unsafe.Pointer(&n))
	fmt.Printf("%+v\n", nHeader)
	length := len(n)
	k = length - (k % length)
	temp := n[:k]
	tempHeader := (*reflect.SliceHeader)(unsafe.Pointer(&temp))
	fmt.Printf("%+v\n", tempHeader)
	n = append(n[:k], n[k:]...)
	nHeader = (*reflect.SliceHeader)(unsafe.Pointer(&n))
	fmt.Printf("%+v\n", nHeader)
	n = append(n[k:], temp...)
	nHeader = (*reflect.SliceHeader)(unsafe.Pointer(&n))
	fmt.Printf("%+v\n", nHeader)
	fmt.Println(n)
}
