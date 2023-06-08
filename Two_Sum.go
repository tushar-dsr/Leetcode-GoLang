//https://leetcode.com/problems/two-sum/description/

//brute force apprach with O(n^2) complexcity

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	//no sum equalt to the target
	return []int{}
}

//O(n*log(n)) approaach

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, val := range nums {
		if v, found := mp[target-val]; found {
			return []int{v, i}
		}
		mp[val] = i
	}

	return []int{}
}
