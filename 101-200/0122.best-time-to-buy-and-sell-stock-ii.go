/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	maxProfit := make([]int, n) // maxProfit[i] 表示前i天的最大利润
	for today := 1; today < n; today++ {
		maxProfit[today] = maxProfit[today-1] // 不选择today售出时，最大利润不变
		// 遍历之前的价格，枚举如果before买进，today卖出产生的最大利润
		for before := today - 1; before >= 0 && prices[before] < prices[today]; before-- {
			if before == 0 {
				maxProfit[today] = max(maxProfit[today], prices[today]-prices[before])
			} else {
				maxProfit[today] = max(maxProfit[today], prices[today]-prices[before]+maxProfit[before-1])
			}
		}
	}
	return maxProfit[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
