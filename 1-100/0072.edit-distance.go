/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
package leetcode

import (
	"math"
)

// @lc code=start
func minDistance(word1 string, word2 string) int {
	word1, word2 = " "+word1, " "+word2
	m, n := len(word1), len(word2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		dp[i][0] = i
	}
	for j := 1; j < n; j++ {
		dp[0][j] = j
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else { //          insert   ||   delete  ||   replace
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m-1][n-1]
}

func min(a ...int) int {
	ret := math.MaxInt64
	for _, v := range a {
		if v < ret {
			ret = v
		}
	}
	return ret
}

// @lc code=end
