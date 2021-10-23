/*

Solution for LC 39. Combination Sum (https://leetcode.com/problems/combination-sum/)

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.


Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.


*/
package main

func combinationSum(candidates []int, target int) [][]int {
	uniqueSums := [][]int{}

	var dfs func(int, []int, int)
	dfs = func(i int, cur []int, targetSum int) {
		if i >= len(candidates) {
			return
		}

		if targetSum > target {
			return
		}

		if targetSum == target {
			uniqueSums = append(uniqueSums, append([]int{}, cur...))
			return
		}

		cur = append(cur, candidates[i])
		dfs(i, cur, targetSum+candidates[i])
		if len(cur) > 0 {
			cur = cur[:len(cur)-1]
		}
		dfs(i+1, cur, targetSum)
	}
	dfs(0, []int{}, 0)
	return uniqueSums

}
