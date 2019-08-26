/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
// i 记录 0 的个数， 同时写入0，会覆盖之前的1或2
// j 记录 0 和 1 的个数，同时写入1，会覆盖之前的2
// k 记录 0，1，2 的个数，同时写入2
// 算法等同于：先写入k个2（全置为2），然后在开头覆写j个1，最后在开头覆写i个0
func sortColors(nums []int) {
	i, j, k := 0, 0, 0
	for _, num := range nums {
		switch num {
		case 0:
			nums[k] = 2
			nums[j] = 1
			nums[i] = 0
			i, j, k = i+1, j+1, k+1
		case 1:
			nums[k] = 2
			nums[j] = 1
			j, k = j+1, k+1
		case 2:
			nums[k] = 2
			k++
		}
	}
}
