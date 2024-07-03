package NumAndSring

import "slices"

// O(n)
func hIndex(citations []int) int {
	n := len(citations)
	nums := make([]int, len(citations)+1)
	for i := 0; i < n; i++ {
		if citations[i] >= n {
			nums[n]++
		} else {
			nums[citations[i]]++
		}

	}
	cur := n
	total := 0
	for cur >= 0 {
		total += nums[cur]
		if total >= cur {
			return cur
		}
		cur--
	}
	return 0
}

func hIndex1(citations []int) int {
	slices.Sort(citations)
	length := len(citations)
	cur := 0
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	for i := 0; i < length; i++ {
		if citations[i] <= length-i {
			cur = max(cur, citations[i])
		} else {
			cur = max(cur, length-i)
		}
	}
	return cur
}
