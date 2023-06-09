//link: https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	hash := make([]int, 128)
	ans := 0
	for i, j := 0, 0; i < len(s); i++ {
		hash[s[i]]++
		for j < i && hash[s[i]] > 1 {
			hash[s[j]]--
			j++
		}
		if i-j+1 > ans {
			ans = i - j + 1
		}
	}
	return ans
}