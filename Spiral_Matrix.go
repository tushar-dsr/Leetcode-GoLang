//link: https://leetcode.com/problems/spiral-matrix/description/

func spiralOrder(mat [][]int) []int {
	ans := []int{}
	if len(mat) == 0 || len(mat[0]) == 0 {
		return ans
	}

	n, m := len(mat), len(mat[0])
	l, r, t, b := 0, m-1, 0, n-1

	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			ans = append(ans, mat[t][i])
		}
		t++

		for i := t; i <= b; i++ {
			ans = append(ans, mat[i][r])
		}
		r--

		if t <= b {
			for i := r; i >= l; i-- {
				ans = append(ans, mat[b][i])
			}
			b--
		}

		if l <= r {
			for i := b; i >= t; i-- {
				ans = append(ans, mat[i][l])
			}
			l++
		}
	}
	return ans
}