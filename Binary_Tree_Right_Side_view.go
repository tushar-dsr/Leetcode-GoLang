// link: https://leetcode.com/problems/binary-tree-right-side-view/description/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var ans []int
	qu := [](*TreeNode){root}
	if root == nil {
		return []int{}
	}
	for len(qu) != 0 {
		ans = append(ans, qu[0].Val)
		sz := len(qu)
		for sz != 0 {
			top := qu[0]
			if top.Right != nil {
				qu = append(qu, top.Right)
			}
			if top.Left != nil {
				qu = append(qu, top.Left)
			}
			qu = qu[1:]
			sz--
		}
	}
	return ans
}