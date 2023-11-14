package main

import (
	"fmt"
	bt "main/backTracking"
	bf "main/bruteForce"
	"time"
)

func main() {
	defer fmt.Println("Finished")
	var subset []int
	numbers := []int{
		10, 2, 7, 3, 4, 5,
	}
	sum := 11

	initTime := time.Now()
	for i := 0; i < 1; i++ {
		subset = BackTracking(numbers, sum)
		// if i%10 == 0 {
		// 	fmt.Println("Brute Force:", i)
		// }
	}

	finalTime := time.Now()
	//fmt.Println("Time elapsed:", float64(finalTime.Nanosecond()-initTime.Nanosecond())/1000, "us or ", float64(finalTime.Nanosecond()-initTime.Nanosecond())/1000000, "ms")
	fmt.Println("Time elapsed:", finalTime.Sub(initTime))
	fmt.Println("Numbers:", numbers)
	fmt.Println("Sum:", sum)
	fmt.Println("Subset:", subset)
}

func BruteForce(numbers []int, sum int) []int {
	// Get subsets with sum with brute force
	subset := bf.GetSumSets(numbers, sum)
	return subset
}

func BackTracking(numbers []int, sum int) []int {
	// Get subsets with sum with backtracking
	subset := bt.GetSumSets(numbers, sum, []int{}, 0)
	return subset
}
