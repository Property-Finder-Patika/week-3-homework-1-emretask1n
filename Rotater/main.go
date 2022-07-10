package main

import "fmt"

//a function that rotates given slice by given number of steps to the right
func rotater(nums []int, k int) []int {

	// If k is negative or len of numbers is equal to 0, function does not print anything
	if k < 0 || len(nums) == 0 {
		return nums
	}

	//else
	//--initial slice
	fmt.Printf(" len: %d cap: %d slice: %v\n", len(nums), cap(nums), nums)

	//rotation
	r := len(nums) - k%len(nums)
	nums = append(nums[r:], nums[:r]...)

	//after rotation
	fmt.Printf(" len: %d cap: %d new slice: %v\n", len(nums), cap(nums), nums)

	return nums
}

func main() {
	numbers := []int{1, 2, 3, 4}
	numbers = rotater(numbers, 2)
}
