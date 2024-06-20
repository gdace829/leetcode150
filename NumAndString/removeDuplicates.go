package NumAndSring

func removeDuplicates(nums []int) int {
	i := 0
	// 下一个与之不同的交换
	j := 1
	if len(nums)==0 {
		return 0
	}
	for j < len(nums) {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}
	return i+1
}
