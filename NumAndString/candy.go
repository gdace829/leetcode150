package NumAndSring

func candy(ratings []int) int {
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	candys := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candys[i] = 1
		if i > 0 {
			if ratings[i] > ratings[i-1] {
				candys[i] = candys[i-1] + 1
			}
		}

	}
	for i := len(ratings) - 2; i >= 0; i-- {

		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i+1]+1, candys[i])
		}

	}
	var sum int
	for i := 0; i < len(ratings); i++ {
		sum += candys[i]

	}
	return sum
}
