package NumAndSring

// 逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	last := len(nums1) - 1
	m--
	n--
	for last != -1 {
		if m >= 0 {
			nums1[last] = nums1[m]
		} else {
			nums1[last] = nums2[n]
		}
		if n>=0 && nums1[last] <= nums2[n] {
			nums1[last] = nums2[n]
			n--
		} else {
			m--
		}
		last--
	}
}
