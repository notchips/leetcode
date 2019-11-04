/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */
package leetcode

// @lc code=start
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	setFirstRow, setFirstCol := false, false
	// 判断第一列是否需要置为0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			setFirstCol = true
			break
		}
	}
	// 判断第一行是否需要置为0
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			setFirstRow = true
			break
		}
	}

	// 利用第一行和第一列，记录对应行或列是否需要置为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j], matrix[i][0] = 0, 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if setFirstRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if setFirstCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end
