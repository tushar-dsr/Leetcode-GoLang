//link: https://leetcode.com/problems/number-of-islands/description/

//used dfs

var dx []int = []int{1, 0, -1, 0}
var dy []int = []int{0, 1, 0, -1}

func isValid(i int, j int, n int, m int) bool {
	if i >= 0 && i < n && j >= 0 && j < m {
		return true
	}
	return false
}

func dfs(grid [][]byte, x int, y int, n int, m int) {
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		cx := x + dx[i]
		cy := y + dy[i]
		if isValid(cx, cy, n, m) && grid[cx][cy] == '1' {
			dfs(grid, cx, cy, n, m)
		}
	}
}

func numIslands(grid [][]byte) int {
	cnt, n, m := 0, len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j, n, m)
				cnt++
			}
		}
	}
	return cnt
}