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
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	var merics [][]string
	for _, v := range digits {
		if arr, ok := m[string(v)]; ok {
			merics = append(merics, arr)
		}
	}
	return dfs(merics, 0, nil)
}

func dfs(arr [][]string, depth int, result []string) []string {
	results := []string{}
	if result == nil {
		result = make([]string, len(arr))
	}
	for i := 0; i < len(arr[depth]); i++ {
		result[depth] = arr[depth][i]
		if depth < len(arr)-1 {
			results = append(results, dfs(arr, depth+1, result)...)
		} else {
			results = append(results, strings.Join(result, ""))
		}
	}
	return results
}

// @lc code=end
