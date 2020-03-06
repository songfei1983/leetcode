package DataStreamAsDisjointIntervals

import "sort"

/*
 * @lc app=leetcode id=352 lang=golang
 *
 * [352] Data Stream as Disjoint Intervals
 */

// @lc code=start
type SummaryRanges struct {
	x []int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{[]int{}}
}

func (this *SummaryRanges) AddNum(val int) {
	for _, v := range this.x {
		if val == v {
			return
		}
	}
	this.x = append(this.x, val)
	sort.Ints(this.x)
}

func (this *SummaryRanges) GetIntervals() [][]int {
	if len(this.x) == 0 {
		return nil
	}
	if len(this.x) == 1 {
		v := this.x[0]
		return [][]int{[]int{v, v}}
	}
	res := [][]int{}
	start, mem, end := this.x[0], this.x[0], this.x[0]
	for i := 1; i < len(this.x); i++ {
		if this.x[i]-1 == mem {
			mem = this.x[i]
		} else {
			end = this.x[i-1]
			res = append(res, []int{start, end})
			start, mem, end = this.x[i], this.x[i], this.x[i]
		}
		if i == len(this.x)-1 {
			end = mem
			res = append(res, []int{start, end})
		}
	}
	return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
// @lc code=end

//[null,null,[[1,1]],null,[[1,1],[3,3]],null,[[1,1],[3,3],[7,7]],null,[[1,1],[3,3],[7,7]],null,[[1,1],[3,3],[6,7]]]
// [null,null,[[1,1]],null,[[1,1],[3,3]],null,[[1,1],[3,3],[7,7]],null,[[1,3],[7,7]],null,[[1,3],[6,7]]]
