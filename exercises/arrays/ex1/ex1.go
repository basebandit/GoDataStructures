package main

import (
	"fmt"
	"time"
)

func main() {
	//boxes := []int{1, 2, 3, 4, 5}
	dates := []string{"2009", "2005", "2006","2020", "2010", "2011", "2017", "2019"}

	fmt.Println(getRecentDate(dates))
	fmt.Println(getRecentDate2(dates))
	//logAllPairsOfArray(boxes)
	//printAllNosThenAllPairSums(boxes)
}

func logAllPairsOfArray(arr []int) {
	start := time.Now()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[i], arr[j])
		}
	}

	fmt.Println("time taken", time.Since(start))
}

func printAllNosThenAllPairSums(nos []int) {
	fmt.Println("These are the numbers: ")
	for _, no := range nos {
		fmt.Println(no)
	}

	fmt.Println("and these are their sums: ")
	for i := 0; i < len(nos); i++ {
		for j := 0; j < len(nos); j++ {
			sum := nos[i] + nos[j]
			fmt.Println(sum)
		}
	}
}

//Takes less time O(n) Fair run time
func getRecentDate(dates []string) string { //o(n)
	start := time.Now()
	recent := dates[0]
	for i := 0; i < len(dates); i++ {
		if dates[i] > recent {
			recent = dates[i]
		}

	}
	fmt.Println("running time1: ", time.Since(start))
	return recent
}

//Takes more time O(n^2) Horrible runtime
func getRecentDate2(dates []string) string { //O(n^2)
	start := time.Now()
	var recent string
	for i := 0; i < len(dates); i++ {
		recent = dates[i]
		for j := 0; j < len(dates); j++ {
			if dates[j] > recent {
				recent = dates[j]
			}
		}
	}
	fmt.Println("running time2: ", time.Since(start))
	return recent
}
