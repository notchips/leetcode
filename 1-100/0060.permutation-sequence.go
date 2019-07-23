/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */
func getPermutation(n int, k int) string {
	ans := make([]byte, n)
	vis := make([]bool, n+1)
	f := make([]int, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] * i
	}
	for i := 0; i < n-1; i++ {
		ans[i] = getNum(vis, n, (k-1)/f[n-1-i]+1)
		k = (k-1)%f[n-1-i] + 1
	}
	ans[n-1] = getNum(vis, n, 1)
	return string(ans)
}

// 从1到n中，返回第m个未被选的数
func getNum(vis []bool, n, m int) byte {
	i, j := 1, 0
	for ; i <= n; i++ {
		if !vis[i] {
			j++
			if j == m {
				break
			}
		}
	}
	vis[i] = true
	return byte(i + '0')
}
