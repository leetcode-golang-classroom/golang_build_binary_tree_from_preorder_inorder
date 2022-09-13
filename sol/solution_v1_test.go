package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTestV1(b *testing.B) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	for idx := 0; idx < b.N; idx++ {
		buildTreeV1(preorder, inorder)
	}
}
func Test_buildTreeV1(t *testing.T) {
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
			if got := buildTreeV1(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTreeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
