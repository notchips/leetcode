/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */
// 对于长度为n的s，可以平均分为两部分:s0:[0, mid), s1:[mid, n)；
// 则numDecodings(s) = numDecodings(s0) * numDecodings(s1)
// 也可以分为三部分：s2:[0, mid-1), s3:[mid-1, mid+1), s4:[mid+1, n)
// 当中间部分满足 "10" <= s3 <= "26"，也就是说s3可以编码时，
// 则numDecodings(s) 还要加上 numDecodings(s2) * numDecodings(s4)
// 递归边界为n等于0和1
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make(map[string]int) // 保存中间结果，防止重复计算
	return numDecodingsHelper(s, dp)
}

func numDecodingsHelper(s string, dp map[string]int) int {
	if _, ok := dp[s]; ok {
		return dp[s]
	}
	switch len(s) {
	case 0:
		// 当长度等于2和3，在切分成三部分时，会出现头或尾部分长度为0
		// 为了防止漏算，故令空字符串dp=1
		dp[s] = 1 
	case 1:
		if s == "0" {
			dp[s] = 0
		} else {
			dp[s] = 1
		}
	default:
	  	mid := len(s) / 2
		dp[s] = numDecodingsHelper(s[:mid], dp) * numDecodingsHelper(s[mid:], dp)
		if "10" <= s[mid-1:mid+1] && s[mid-1:mid+1] <= "26" {
			dp[s] += numDecodingsHelper(s[:mid-1], dp) * numDecodingsHelper(s[mid+1:], dp)
		}
	}
	return dp[s]
}
