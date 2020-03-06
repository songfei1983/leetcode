package MostCommonWord

/*
 * @lc app=leetcode id=819 lang=golang
 *
 * [819] Most Common Word
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	wc := map[string]int{}
	ch := []byte{}
	isBanned := func(w string) bool {
		for _, b := range banned {
			if b == w {
				return true
			}
		}
		return false
	}
	findWord := func(wc map[string]int) string {
		var w string
		var c int
		for k, v := range wc {
			if v > c {
				c = v
				w = k
			}
		}
		return w
	}
	for k, a := range paragraph {
		c := byte(a)
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			if c >= 65 && c <= 90 {
				c += 32
			}
			ch = append(ch, c)
		} else {
			w := string(ch)
			if len(w) > 0 && !isBanned(w) {
				wc[string(ch)]++
			}
			ch = []byte{}
		}
		if len(paragraph)-1 == k {
			w := string(ch)
			if !isBanned(w) {
				wc[string(ch)]++
			}
			ch = []byte{}
		}
	}
	return findWord(wc)
}

// @lc code=end
