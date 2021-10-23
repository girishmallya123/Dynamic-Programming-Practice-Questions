/*
Solution to LC 70. Climbing Stairs (https://leetcode.com/problems/climbing-stairs/)
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

*/

package main

func climbStairs(n int) int {
	// creating a memoization map
	memo := map[int]int{}

	var climbStairsRec func(int, int, map[int]int) int
	climbStairsRec = func(i int, n int, memo map[int]int) int {
		if i == n {
			// if we reach the destination, return 1 back to the caller, indicating that we reached the top.
			return 1
		}
		if i > n {
			// if we have gone past the top of the stairs, return 0, since we cannot count the path we have taken.
			return 0
		}

		val, exists := memo[i]
		if exists {
			// if the climbStairs[i] is already computed, return it.
			return val
		}

		// add the number of ways by recursively climbing 1 step at a time and 2 steps at a time.
		memo[i] = climbStairsRec(i+1, n, memo) + climbStairsRec(i+2, n, memo)
		return memo[i]

	}

	return climbStairsRec(1, n, memo) + climbStairsRec(2, n, memo)

}
