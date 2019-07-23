/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	cnt := 1
	for i, j := 0, 0; i < (n+1)/2; i, j = i+1, j+1 {
		for k := j; k <= n-1-j; k++ {
			ans[i][k] = cnt
			cnt++
		}

		for k := i + 1; k <= n-2-i; k++ {
			ans[k][n-1-j] = cnt
			cnt++
		}

		if i != n-1-i {
			for k := n - 1 - j; k >= j; k-- {
				ans[n-1-i][k] = cnt
				cnt++
			}
		}

		if j != n-1-j {
			for k := n - 2 - i; k >= i+1; k-- {
				ans[k][j] = cnt
				cnt++
			}
		}
	}

	return ans
}
