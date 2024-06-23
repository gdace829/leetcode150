package NumAndSring

func majorityElement(nums []int) int {
	cur := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {

		if nums[i] != cur {
			sum--
		} else {
			sum++
		}
		if sum == 0 {
			cur = nums[i+1]
		}
	}
	return cur
}
