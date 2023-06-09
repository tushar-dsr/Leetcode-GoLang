//link: https://leetcode.com/problems/flood-fill/description/

var dx []int = []int{0, 1, 0, -1}
var dy []int = []int{1, 0, -1, 0}

func dfs(image [][]int, sr int, sc int, curr int) {
	image[sr][sc] = -1
	for i := 0; i < 4; i++ {
		cr := sr + dx[i]
		cc := sc + dy[i]
		if cr >= 0 && cr < len(image) && cc >= 0 && cc < len(image[0]) && image[cr][cc] == curr && image[cr][cc] != -1 {
			dfs(image, cr, cc, curr)
		}
	}
}
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	dfs(image, sr, sc, image[sr][sc])
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if image[i][j] == -1 {
				image[i][j] = color
			}
		}
	}
	return image
}