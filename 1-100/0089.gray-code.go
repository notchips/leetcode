/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */
func grayCode(n int) []int {
	ans := make([]int, 1, 1<<uint(n))
	p := 1
	for i := 0; i < n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ans[j]+p)
		}
		p *= 2
	}
	return ans
}
