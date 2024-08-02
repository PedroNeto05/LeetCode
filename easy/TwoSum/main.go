package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {

	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {

			if nums[i]+nums[j] == target && i != j {
				result := append(result, i, j)
				return result
			}

		}
	}
	return result
}
