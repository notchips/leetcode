/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */
func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = i * factorial[i-1]
	}
	vis := make([]bool, n)
	buf := make([]byte, n)
	for i := range buf {
		pos := (k - 1) / factorial[n-1-i]
		k -= pos * factorial[n-1-i]
		buf[i] = getNumberKNotVis(vis, pos)
	}
	return string(buf)
}

// 从1-n，选出第pos个未访问的数
func getNumberKNotVis(vis []bool, pos int) byte {
	cnt, i := 0, 0
	for ; i < len(vis); i++ {
		if !vis[i] {
			if cnt == pos {
				vis[i] = true
				break
			}
			cnt++
		}
	}
	return byte(i) + 1 + '0'
}
