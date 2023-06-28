//Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/?envType=study-plan-v2&envId=top-interview-150

func maxProfit(prices []int) int {
	curr := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		if curr > prices[i] {
			curr = prices[i]
		}
		tempAns := prices[i] - curr
		if tempAns > ans {
			ans = tempAns
		}
	}
	return ans
}