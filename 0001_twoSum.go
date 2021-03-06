package main

/*
---------------------------------------------------------EASY---------------------------------------------------------

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

// First try	Runtime: 73 ms, Memory Usage: 3.9 MB
func twoSum(nums []int, target int) []int {
	newNums := []int{}
	flag := false
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				newNums = append(newNums, i, j)
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	return newNums
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
}
