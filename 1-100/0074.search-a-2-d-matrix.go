/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// 二分确定行数
	top, down := 0, len(matrix)-1
	for top < down {
		mid := top + (down-top)/2 + 1
		if matrix[mid][0] > target {
			down = mid - 1
		} else {
			top = mid
		}
	}
	row := top

	// 在改行进行二分查找
	left, right := 0, len(matrix[0])-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
