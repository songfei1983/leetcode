package AddBinary

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	x, y, res := []byte(a), []byte(b), []byte{}
	if len(x) < len(y) {
		y, x = x, y
	}
	ansLen := len(x)
	c := byte(0)
	for i := ansLen - 1; i >= 0; i-- {
		a := x[i] - '0'
		r, b := byte(0), byte(0)
		index := i - (len(x) - len(y))
		if index < len(y) && index >= 0 {
			b = y[index] - '0'
		}
		r, c = add(a, b, c)
		res = append([]byte{r + '0'}, res...)
	}
	if c > 0 {
		res = append([]byte{c + '0'}, res...)
	}
	return string(res)
}

func add(a, b, c byte) (r byte, carrier byte) {
	r = (a + b + c)
	return r % 2, r / 2
}

// @lc code=end
