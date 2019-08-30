/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for row := 1; row <= rowIndex; row++ {
		pre := 0 // 记录上一层同一位置左边的值
		for col := 0; col < row+1; col++ {
			ans[col], pre = pre+ans[col], ans[col]
		}
	}
	return ans
}
