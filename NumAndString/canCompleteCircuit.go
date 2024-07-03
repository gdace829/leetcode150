package NumAndSring

import "fmt"

// 如果x到达不了y+1，那么x-y之间的点也不可能到达
// y+1开始
func canCompleteCircuit(gas []int, cost []int) int {
	cur := 0
	curYou := 0
	start := 0
	for start < len(cost) {
		fmt.Println(start)
		curYou += gas[cur]
		if curYou-cost[cur] >= 0 {
			curYou -= cost[cur]
			cur++
			cur = cur % len(gas)
		} else {
			curYou = 0
			if cur < start {
				return -1
			}
			start = cur + 1
			cur = start
			continue
		}
		if cur == start {
			return start
		}

	}
	return -1
}
