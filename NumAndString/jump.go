package NumAndSring

// O(n) 打破边界就跳一次
func jump(nums []int) int {
	end := 0
	step := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > max {
			max = nums[i] + i
		}
		if i == end {
			end = max
			step++
		}
	}
	return step
}
