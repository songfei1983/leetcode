package lettercombinationsofaphonenumber

import "strings"

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	var merics [][]string
	for _, v := range digits {
		if arr, ok := m[string(v)]; ok {
			merics = append(merics, arr)
		}
	}
	results = make([]string, 0)
	result = make([]string, len(merics))
	dfs(merics, 0)
	return results
}

var results []string
var result []string

func dfs(arr [][]string, depth int) {
	for i := 0; i < len(arr[depth]); i++ {
		result[depth] = arr[depth][i]
		if depth < len(arr)-1 {
			dfs(arr, depth+1)
		} else {
			results = append(results, strings.Join(result, ""))
		}
	}
}

// @lc code=end
