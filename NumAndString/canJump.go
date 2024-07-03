package NumAndSring

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	curMax := 0
	cur := 0
	for cur <= curMax {
		fmt.Println(cur, curMax)
		if cur+nums[cur] > curMax {
			curMax = cur + nums[cur]
			if curMax >= len(nums)-1 {
				return true
			}
		}
		if cur <= curMax {
			cur++
		}
	}
	return curMax >= len(nums)-1
}
