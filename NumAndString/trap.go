package NumAndSring

func trap(height []int) int {
	var min func(int, int) int
	min = func(i1, i2 int) int {
		if i1 > i2 {
			return i2
		}
		return i1
	}
	stack := []int{}
	res := 0
	for i := 0; i < len(height); i++ {
		// 初始化0
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		//比zhan顶大
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			if len(stack) >= 2 {
				res += (i - (stack[len(stack)-2]) - 1) * (min(height[stack[len(stack)-2]], height[i]) - height[stack[len(stack)-1]])

			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res

}
