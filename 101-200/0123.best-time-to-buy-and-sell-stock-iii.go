/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// maxProfitEndWith[i] 表示从0到i天内只交易一次的最大利润
	maxProfitEndWith := make([]int, n)
	for minPrice, today := prices[0], 1; today < n; today++ {
		maxProfitEndWith[today] = max(maxProfitEndWith[today-1], prices[today]-minPrice)
		minPrice = min(minPrice, prices[today])
	}
	// maxProfitStartWith[i] 表示从i到n-1天内只交易一次的最大利润
	maxProfitStartWith := make([]int, n)
	for maxPrice, today := prices[n-1], n-2; today >= 0; today-- {
		maxProfitStartWith[today] = max(maxProfitStartWith[today+1], maxPrice-prices[today])
		maxPrice = max(maxPrice, prices[today])
	}
	maxProfit := maxProfitStartWith[0]
	for mid := 0; mid < n-1; mid++ {
		maxProfit = max(maxProfit, maxProfitEndWith[mid]+maxProfitStartWith[mid+1])
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
