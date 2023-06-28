//link: https://leetcode.com/problems/majority-element/description/?envType=study-plan-v2&envId=top-interview-150

// O(n) extra space by using map
func majorityElement(nums []int) int {
	mp := make(map[int]int)
	for _, val := range nums {
		mp[val]++
	}
	mi, ans := 0, 0
	for key, val := range mp {
		if val > mi {
			mi = val
			ans = key
		}
	}
	return ans
}

//O(n) time and O(1) space

func majorityElement(nums []int) int {
	var major int
	for i, cnt := 0, 0; i < len(nums); i++ {
		if cnt == 0 {
			cnt++
			major = nums[i]
		} else if major == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return major
}
