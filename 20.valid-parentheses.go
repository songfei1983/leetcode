package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	pop := func(bs []byte) []byte { return bs[:len(bs)-1] }
	push := func(b byte, bs []byte) []byte { return append(bs, b) }
	last := func(bs []byte) byte { return bs[len(bs)-1] }
	bs := []byte(s)
	filo := []byte{}
	for _, v := range bs {
		if len(filo) > 0 && isPair(last(filo), v) {
			filo = pop(filo)
		} else {
			filo = push(v, filo)
		}
	}
	return len(filo) == 0
}

func isPair(a, b byte) bool {
	if a == '(' && b == ')' {
		return true
	}
	if a == '[' && b == ']' {
		return true
	}
	if a == '{' && b == '}' {
		return true
	}
	return false
}

// @lc code=end
