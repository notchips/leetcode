/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)

	// 以(i, j)为起点，每次循环遍历完一圈
	for i, j := 0, 0; i < (m+1)/2 && j < (n+1)/2; i, j = i+1, j+1 { 

		// 从左往右，遍历上行
		for k := j; k < n-j; k++ {
			ans = append(ans, matrix[i][k])
		}

		// 从上往下，遍历右列
		for k := i + 1; k < m-1-i; k++ {
			ans = append(ans, matrix[k][n-1-j])
		}

		// 从右往左，遍历下行
		if i != m-1-i {
			for k := n - 1 - j; k >= j; k-- {
				ans = append(ans, matrix[m-1-i][k])
			}
		}
		
		// 从下往上，遍历左列
		if j != n-1-j {
			for k := m - 2 - i; k >= i+1; k-- {
				ans = append(ans, matrix[k][j])
			}
		}
	}

	return ans
}
