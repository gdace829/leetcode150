package NumAndSring

func productExceptSelf(nums []int) []int {
	formNums := make([]int, len(nums))
	formNums[0] = nums[0]
	laterNums := make([]int, len(nums))
	laterNums[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		formNums[i] = nums[i] * formNums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		laterNums[i] = nums[i] * laterNums[i+1]
	}
	ans := make([]int, len(nums))
	ans[0] = laterNums[1]
	ans[len(nums)-1] = formNums[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		ans[i] = formNums[i-1] * laterNums[i+1]
	}
	return ans
}
