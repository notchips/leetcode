/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
func minWindow(s string, t string) string {
	var tCharCnt [256]int // 记录t中各字符的个数
	for i := 0; i < len(t); i++ {
		tCharCnt[t[i]]++
	}

	sCharIndex := make([]int, 0, 10) // 记录s关键索引（即s[i]包含在t中的索引i）
	for i := 0; i < len(s); i++ {
		// 如果字符s[i]在t中
		if tCharCnt[s[i]] > 0 {
			sCharIndex = append(sCharIndex, i)
		}
	}

	var minLeft, minRight int       // 记录最小子串索引值
	minLen, totalCnt := len(s)+1, 0 // minLen记录最小子串长度
	var sCharCnt [256]int           // 记录子串的每个字符的个数
	// 遍历s的关键索引
	for left, right := 0, 0; right < len(sCharIndex); right++ {
		leftChar, rightChar := s[sCharIndex[left]], s[sCharIndex[right]]
		sCharCnt[rightChar]++

		// 记录有效字符数
		if sCharCnt[rightChar] <= tCharCnt[rightChar] {
			totalCnt++
		}

		// 去除左边多余字符
		for sCharCnt[leftChar] > tCharCnt[leftChar] {
			sCharCnt[leftChar]--
			left++
			leftChar = s[sCharIndex[left]]
		}

		// 当s[sCharIndex[left], sCharIndex[right]] 包含t时
		if totalCnt == len(t) {
			curLen := sCharIndex[right] - sCharIndex[left] + 1
			if curLen < minLen {
				minLen = curLen
				minLeft, minRight = sCharIndex[left], sCharIndex[right]
			}
		}
	}
	if totalCnt < len(t) {
		return ""
	}
	return s[minLeft : minRight+1]
}
