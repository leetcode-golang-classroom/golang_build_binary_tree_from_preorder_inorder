package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	tree := &TreeNode{Val: preorder[0]}
	mid := FindIdx(&inorder, preorder[0])
	tree.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	tree.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return tree
}

func FindIdx(arr *[]int, target int) int {
	for idx, val := range *arr {
		if val == target {
			return idx
		}
	}
	return -1
}
