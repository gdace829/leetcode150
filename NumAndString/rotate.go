package NumAndSring

// 先整体反转再分段反转
func rotate(nums []int, k int) {
	start := k % len(nums)
	reverse(&nums, 0, len(nums)-1)
	reverse(&nums, 0, start-1)
	reverse(&nums, start, len(nums)-1)
}
func reverse(nums *[]int, i, j int) {
	nums1 := *nums
	for i < j {
		t := nums1[i]
		nums1[i] = nums1[j]
		nums1[j] = t
		i++
		j--
	}

}
