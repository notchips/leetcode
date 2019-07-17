/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	tempArr := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if i == 0 && (str[i] == '-' || str[i] == '+') {
			tempArr = append(tempArr, str[i])
		} else if str[i] < '0' || str[i] > '9' {
			break
		} else {
			tempArr = append(tempArr, str[i])
		}
	}
	if len(tempArr) == 0 || (len(tempArr) == 1 && (tempArr[0] == '-' || tempArr[0] == '+')) {
		return 0
	}
	ret, err := strconv.ParseInt(string(tempArr), 10, 32)
	if err != nil {
		if tempArr[0] == '-' {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	return int(ret)
}
