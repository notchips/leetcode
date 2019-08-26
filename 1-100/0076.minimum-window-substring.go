/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
func minWindow(s string, t string) string {
	hashT := make([]int, 128)
	for i := 0; i < len(t); i++ {
		hashT[t[i]]++
	}
	indexes := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		if hashT[s[i]] > 0 {
			indexes = append(indexes, i)
		}
	}
	hashS := make([]int, 128)
	var minLeft, minRight int
	minLen := math.MaxInt64
	cnt := 0
	for curLeft, curRight := 0, 0; curRight < len(indexes); curRight++ {
		leftChar, rightChar := s[indexes[curLeft]], s[indexes[curRight]]
		if hashS[rightChar] < hashT[rightChar] {
			cnt++
		}
		hashS[rightChar]++
		// 去除左边多余字符
		for hashS[leftChar] > hashT[leftChar] {
			hashS[leftChar]--
			curLeft++
			leftChar = s[indexes[curLeft]]
		}
		if cnt == len(t) {
			curLen := indexes[curRight] - indexes[curLeft] + 1
			if curLen < minLen {
				minLen, minLeft, minRight = curLen, curLeft, curRight
			}
		}
	}
	if cnt < len(t) {
		return ""
	}
	return s[indexes[minLeft] : indexes[minRight]+1]
}
