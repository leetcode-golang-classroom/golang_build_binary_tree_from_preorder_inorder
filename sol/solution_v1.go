package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTreeV1(preorder []int, inorder []int) *TreeNode {
	inorderIdxMap := make(map[int]int)
	for idx, val := range inorder {
		inorderIdxMap[val] = idx
	}
	preorderIdx := 0
	var buildTreeRecursive func(preorder []int, left int, right int) *TreeNode
	buildTreeRecursive = func(preorder []int, left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		rootValue := preorder[preorderIdx]
		root := &TreeNode{Val: rootValue}
		preorderIdx++
		root.Left = buildTreeRecursive(preorder, left, inorderIdxMap[rootValue]-1)
		root.Right = buildTreeRecursive(preorder, inorderIdxMap[rootValue]+1, right)
		return root
	}
	return buildTreeRecursive(preorder, 0, len(preorder)-1)
}
