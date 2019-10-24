package main

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	sum := 0
	bs := []byte(s)
	for i := len(bs) - 1; i >= 0; i-- {
		l, r := 0, 0
		if i > 0 {
			l = convert(bs[i-1])
		}
		r = convert(bs[i])
		if l < r {
			r -= l
			i--
		}
		sum += r
	}
	return sum
}

func convert(s byte) int {
	i := 0
	switch s {
	case 'I':
		i = 1
	case 'V':
		i = 5
	case 'X':
		i = 10
	case 'L':
		i = 50
	case 'C':
		i = 100
	case 'D':
		i = 500
	case 'M':
		i = 1000
	}
	return i
}

// @lc code=end
