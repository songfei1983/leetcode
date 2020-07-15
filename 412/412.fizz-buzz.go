package fizzbuzz

import "strconv"

/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	ws := make([]string, n)
	for i := 1; i <= n; i++ {
		w := strconv.Itoa(i)
		if i%3 == 0 && i%5 == 0 {
			w = "FizzBuzz"
		} else if i%3 == 0 {
			w = "Fizz"
		} else if i%5 == 0 {
			w = "Buzz"
		}
		ws[i-1] = w
	}
	return ws
}

// @lc code=end
