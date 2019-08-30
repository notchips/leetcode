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
	i, j := 0, n-1 //从右上角开始遍历，每次可排除一行或一列
	for i < m-1 && j > 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	if i == m-1 { // 遍历到最后一行，二分搜索
		left, right := 0, j
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
	} else { // j == 0， 遍历到第一列，二分搜索
		left, right := i, m-1
		for left <= right {
			mid := left + (right-left)/2
			if matrix[mid][0] == target {
				return true
			} else if matrix[mid][0] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
