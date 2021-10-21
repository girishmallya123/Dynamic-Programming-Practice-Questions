/*

This is the solution to: Leetcode 62. Unique Paths (https://leetcode.com/problems/unique-paths/) using Dynamic Programming with Memoization.

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

*/

package main

import (
	"strconv"
)

func uniquePathsRec(m int, n int, memo map[string]int) int {

	key := strconv.Itoa(m) + "," + strconv.Itoa(n)

	val, exist := memo[key]
	if exist {
		return val
	}

	if m == 1 || n == 1 {
		return 1
	}
	memo[key] = uniquePathsRec(m-1, n, memo) + uniquePathsRec(m, n-1, memo)
	return memo[key]

}

func uniquePaths(m int, n int) int {

	memo := map[string]int{}

	return uniquePathsRec(m, n, memo)
}
