package FindWordsThatCanBeFormedByCharacters

/*
 * @lc app=leetcode id=1160 lang=golang
 *
 * [1160] Find Words That Can Be Formed by Characters
 */

// @lc code=start
func countCharacters(words []string, chars string) int {
	if len(words) < 1 || len(words) > 1000 || len(chars) < 1 || len(chars) > 100 {
		return 0
	}
	count := 0
	for _, w := range words {
		if len(w) < 1 || len(w) > 1000 {
			return 0
		}
		if containsWord(chars, w) {
			count += len(w)
		}
	}
	return count
}

func containsWord(a, b string) bool {
	r := 0
	ra, rb := []byte(a), []byte(b)
	for _, x := range rb {
		for ky, y := range ra {
			if x == y {
				ra[ky] = 0
				r++
				break
			}
		}
	}
	return r == len(b)
}

// @lc code=end
