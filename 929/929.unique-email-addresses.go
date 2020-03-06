package UniqueEmailAddresses

/*
 * @lc app=leetcode id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	m := map[string]int{}
	for _, mail := range emails {
		by := []byte(mail)
		b, c := 0, 0
		for k, v := range by {
			if v == '@' {
				c = k
				if b == 0 {
					b = k
				}
			}
			if b == 0 && v == '+' {
				b = k
			}
		}
		str1, str2 := by[:b], by[c:]
		str := ""
		for _, v := range str1 {
			if v == '.' {
				continue
			}
			str += string(v)
		}
		m[str+string(str2)] = 1
	}
	return len(m)
}

// @lc code=end
