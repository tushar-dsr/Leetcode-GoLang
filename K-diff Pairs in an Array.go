// link : https://leetcode.com/problems/k-diff-pairs-in-an-array/description/

func findPairs(nums []int, k int) int {

	mp := make(map[int]int)
	ans := 0

	for _, num := range nums {
		mp[num]++
	}

	for num, count := range mp {
		if k == 0 {
			if count > 1 {
				ans++
			}
		} else {
			if mp[num+k] > 0 {
				ans++
			}
		}
	}

	return ans
}
