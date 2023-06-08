// link: https://leetcode.com/problems/buy-two-chocolates/description/

func buyChoco(prices []int, money int) int {
	first := int(math.Min(float64(prices[0]), float64(prices[1])))
	second := int(math.Max(float64(prices[0]), float64(prices[1])))

	for i := 2; i < len(prices); i++ {
		if prices[i] < first {
			second = first
			first = prices[i]
		} else if prices[i] < second {
			second = prices[i]
		}
	}
	if first+second > money {
		return money
	}
	return money - (first + second)
}