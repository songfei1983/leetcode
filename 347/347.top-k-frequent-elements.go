package topkfrequentelements

import "sort"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	res := make([]int, k)
	pl := make(PairList, len(nums))
	for _, v := range nums {
		m[v]++
	}
	i := 0
	for k, v := range m {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	if len(pl) < k {
		k = len(pl)
	}
	for k, p := range pl[:k] {
		res[k] = p.Key
	}
	return res
}

// Pair is struct
type Pair struct {
	Key   int
	Value int
}

// PairList is struct
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// @lc code=end
