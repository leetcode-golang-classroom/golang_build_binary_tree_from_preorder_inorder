package sol

import (
	"reflect"
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	arr := *input
	result := make([]*TreeNode, len(arr))
	for idx, val := range arr {
		if val != "null" {
			num, _ := strconv.Atoi(val)
			result[idx] = &TreeNode{Val: num}
		}
	}
	tree = result[0]
	for idx, node := range result {
		if 2*idx+1 < len(result) {
			node.Left = result[2*idx+1]
		}
		if 2*idx+2 < len(result) {
			node.Right = result[2*idx+2]
		}
	}
	return tree
}
func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]",
			args: args{preorder: []int{3, 9, 20, 15, 7}, inorder: []int{9, 3, 15, 20, 7}},
			want: CreateBinaryTree(&[]string{"3", "9", "20", "null", "null", "15", "7"}),
		},
		{
			name: "preorder = [-1], inorder = [-1]",
			args: args{preorder: []int{-1}, inorder: []int{-1}},
			want: CreateBinaryTree(&[]string{"-1"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
