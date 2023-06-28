//link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

func removeDuplicates(nums []int) int {
	l := 0
	for r := 0; r < len(nums); {
		for r < len(nums) && nums[l] == nums[r] {
			r++
		}
		if r < len(nums) {
			nums[l+1] = nums[r]
			l++
			r++
		}
	}
	return l + 1
}