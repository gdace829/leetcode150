package NumAndSring

// 交换到后面开始算，On复杂度
func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	if len(nums) == 0 {
		return 0
	}
	for i != j {
		if nums[i] == val {
			t := nums[i]
			nums[i] = nums[j]
			nums[j] = t
			j--
		} else {
			i++
		}
	}
	if nums[i] == val {
		return i
	} else {
		return i + 1
	}

}
