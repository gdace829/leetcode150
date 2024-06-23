package NumAndSring

// 双指针往后遍历，需要考虑一次换一个还是换两个
func removeDuplicates2(nums []int) int {
	i := 0
	j := 1
	if len(nums) == 0 {
		return 0
	}
	if len(nums) > 1 && nums[i] == nums[j] {
		i = i + 1
	}
	for j < len(nums) {
		if nums[j] == nums[i] {
			j++
		} else {
			// 此时交换两个
			if j+1 < len(nums) && nums[j] == nums[j+1] {
				nums[i+1] = nums[j]
				nums[i+2] = nums[j+1]
				i = i + 2
			} else { //交换一个
				nums[i+1] = nums[j]
				i = i + 1
			}
		}
	}
	return i + 1
}
