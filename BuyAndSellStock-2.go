//link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/submissions/976783844/?envType=study-plan-v2&envId=top-interview-150

func maxProfit(prices []int) int {
	profit, curr := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-curr > 0 {
			profit += prices[i] - curr
		}
		curr = prices[i]
	}
	return profit
}