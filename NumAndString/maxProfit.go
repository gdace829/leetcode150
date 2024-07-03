package NumAndSring

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return res
}
