/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1 //从右上角开始遍历，每次可排除一行或一列
	for row < m-1 && col > 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}
	if row == m-1 { // 遍历到最后一行，二分搜索
		left, right := 0, col
		for left <= right {
			mid := left + (right-left)/2
			if matrix[m-1][mid] == target {
				return true
			} else if matrix[m-1][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	} else { // col == 0， 遍历到第一列，二分搜索
		top, down := row, m-1
		for top <= down {
			mid := top + (down-top)/2
			if matrix[mid][0] == target {
				return true
			} else if matrix[mid][0] < target {
				top = mid + 1
			} else {
				down = mid - 1
			}
		}
	}
	return false
}
