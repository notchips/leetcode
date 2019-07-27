/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
/*
    0 1 2 3 4 5 6 7 8 9
    _ e x e c u t i o n
0 _ 0 1 2 3 4 5 6 7 8 9
1 i 1 1 2 3 4 5 6 6 7 8
2 n 2 2 2 3 4 5 6 7 7 7
3 t 3 3 3 3 4 5 5 6 7 8
4 e 4 3 4 3 4 5 6 6 7 8
5 n 5 4 4 4 4 5 6 7 7 7
6 t 6 5 5 5 5 5 5 6 7 8
7 i 7 6 6 6 6 6 6 5 6 7
8 o 8 7 7 7 7 7 7 6 5 6
9 n 9 8 8 8 8 8 8 7 6 5
*/
func minDistance(word1 string, word2 string) int {
	word1, word2 = " "+word1, " "+word2
	m, n := len(word1), len(word2)
	// dp[i][j] 表示 word1[0,i] 转换到 word2[0,j] 所需的最少操作数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// word1 转换到 空字符串 最少操作数（删除），等于len(word1)
	for i := 0; i < m; i++ {
		dp[i][0] = i
	}

	// 空字符串 转换到 word2 最少操作数（插入），等于len(word12)
	for j := 0; j < n; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] { // 不需要操作
				dp[i][j] = dp[i-1][j-1]
			} else {     // 取 替换/删除/插入 中的最小值
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[m-1][n-1]
}

func min(nums ...int) int {
	t := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[t] {
			t = i
		}
	}
	return nums[t]
}
