/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
package leetcode

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)
	rowStart, colStart := 0, 0
	rowEnd, colEnd := m-1, n-1
	for rowStart <= rowEnd && colStart <= colEnd {
		for col := colStart; col <= colEnd; col++ {
			ans = append(ans, matrix[rowStart][col])
		}
		for row := rowStart + 1; row < rowEnd; row++ {
			ans = append(ans, matrix[row][colEnd])
		}
		if rowStart < rowEnd {
			for col := colEnd; col >= colStart; col-- {
				ans = append(ans, matrix[rowEnd][col])
			}
		}
		if colStart < colEnd {
			for row := rowEnd - 1; row > rowStart; row-- {
				ans = append(ans, matrix[row][colStart])
			}
		}
		rowStart++
		colStart++
		rowEnd--
		colEnd--
	}
	return ans
}

// @lc code=end
