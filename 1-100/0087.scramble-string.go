/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */
type pair struct {
	s1, s2 string
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	dp := make(map[pair]bool) // 保存结果，避免重复计算
	return isScrambleHelper(s1, s2, dp)
}

func isScrambleHelper(s1, s2 string, dp map[pair]bool) bool {
	p := pair{s1, s2}
	// 如果s1,s2已经有了结果，则直接返回
	if r, ok := dp[p]; ok {
		return r
	}

	if s1 == s2 {
		return true
	}

	n := len(s1)
	// 检测s1和s2是否包含相同的字符
	charCnt := make(map[byte]int)
	for i := 0; i < n; i++ {
		charCnt[s1[i]]++
		charCnt[s2[i]]--
	}
	for _, cnt := range charCnt {
		if cnt != 0 {
			return false
		}
	}

	// 切分s1和s2，判断对应的子串是否满足，其中s2有两种切分（对应左右互换）
	for i := 1; i < n; i++ {
		dp[p] = isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) ||
			isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i])
		if dp[p] {
			return true
		}
	}
	return false
}
