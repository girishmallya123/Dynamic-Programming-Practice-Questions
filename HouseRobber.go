/*

Solution to LC 198. House Robber (https://leetcode.com/problems/house-robber/)
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.


*/

package main

func rob(nums []int) int {

	n := len(nums)
	var res int

	// declare a memoization map
	memo := make(map[int]int)

	//create a function that will check the maximum money robbed from the ith house in the array
	var robFrom func(int, []int) int
	robFrom = func(i int, nums []int) int {
		//base case, we cannot rob anymore if we are past the last index in the arry
		if i >= n {
			return 0
		}

		// check if the value for robbing from the ith house exists in the memoized map
		val, exists := memo[i]
		if exists {
			return val
		}

		//check if it is more profitable to rob from the next house or continue robbing from the same house and proceed to the next non adjacent house.
		robFromNextHouse := robFrom(i+1, nums)
		robFromCurrentHouse := robFrom(i+2, nums) + nums[i]
		if robFromNextHouse > robFromCurrentHouse {
			res = robFromNextHouse
		} else {
			res = robFromCurrentHouse
		}
		memo[i] = res
		return res
	}
	return robFrom(0, nums)
}
