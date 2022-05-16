# golang_build_binary_tree_from_preOrder_inOrder

Given two integer arrays `preorder` and `inorder` where `preorder` is the preorder traversal of a binary tree and `inorder` is the inorder traversal of the same tree, construct and return *the binary tree*.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/02/19/tree.jpg](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

```

**Example 2:**

```
Input: preorder = [-1], inorder = [-1]
Output: [-1]

```

**Constraints:**

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` and `inorder` consist of **unique** values.
- Each value of `inorder` also appears in `preorder`.
- `preorder` is **guaranteed** to be the preorder traversal of the tree.
- `inorder` is **guaranteed** to be the inorder traversal of the tree.

## 解析

題目以 Pre-Order Traversal 與 In-Order Traversal 給出兩個整數陣列

求實作一個演算法建立出這棵二元樹

Pre-Order Traversal 特性會有以下走訪順序如下：

二元樹root , root.Left 子樹, root.Right 子樹

In-Order Traversal 特性會有以下走訪順序如下：

root.Left 子樹, 二元樹root, root.Right 子樹

觀察下圖

![](https://i.imgur.com/KYVbaXL.png)

可以發現

每個 pre-order的點先遇到的都是子樹的跟結點

透過 pre-order的點去找到在 in-order的位置剛好可以把左右子數分開

所以可以透過遞迴關係來建立二元樹

## 程式碼

## 困難點

1. 看出 Pre-Order Traversal 與 In-Order Traversal 的特性
2. 看出遞迴關係

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity 