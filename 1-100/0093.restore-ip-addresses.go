/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */
func restoreIpAddresses(s string) []string {
	addresses := make([]string, 0, 32)
	for cnt1 := 1; cnt1 <= 3; cnt1++ {
		for cnt2 := 1; cnt2 <= 3; cnt2++ {
			for cnt3 := 1; cnt3 <= 3; cnt3++ {
				cnt4 := len(s) - cnt1 - cnt2 - cnt3
				if 1 <= cnt4 && cnt4 <= 3 {
					IPNums := []string{s[:cnt1], s[cnt1 : cnt1+cnt2],
						s[cnt1+cnt2 : cnt1+cnt2+cnt3], s[cnt1+cnt2+cnt3:]}
					if isValidIPNums(IPNums) {
						address := strings.Join(IPNums, ".")
						addresses = append(addresses, address)
					}
				}
			}
		}
	}
	return addresses
}

func isValidIPNums(IPNums []string) bool {
	if len(IPNums) != 4 {
		return false
	}
	for _, IPNum := range IPNums {
		if !isValidIPNum(IPNum) {
			return false
		}
	}
	return true
}

func isValidIPNum(IPNum string) bool {
	if (len(IPNum) == 1 && "0" <= IPNum && IPNum <= "9") ||
		(len(IPNum) == 2 && "10" <= IPNum && IPNum <= "99") ||
		(len(IPNum) == 3 && "100" <= IPNum && IPNum <= "255") {
		return true
	}
	return false
}
