package examples

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	inorder := ""
	preorder := ""

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		val := strconv.Itoa(root.Val)
		preorder += val + ","
		dfs(root.Left)
		inorder += val + ","
		dfs(root.Right)
	}

	dfs(root)
	return preorder + "$" + inorder
}

func (c *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, "$")
	in := s[1]
	pre := s[0]

	ina := strings.Split(in, ",")
	prea := strings.Split(pre, ",")
	l := len(ina)
	i, p := make([]int, l-1), make([]int, l-1)
	for k := 0; k < l-1; k++ {
		i[k], _ = strconv.Atoi(ina[k])
		p[k], _ = strconv.Atoi(prea[k])
	}

	var ct func(i, p []int) *TreeNode
	ct = func(i, p []int) *TreeNode {
		if len(i) == 0 {
			return nil
		}

		ii := 0
		for idx, ele := range i {
			if ele == p[0] {
				ii = idx
				break
			}
		}

		leftl := ii
		left := ct(i[0:ii], p[1:1+leftl])
		right := ct(i[ii+1:], p[1+leftl:])

		root := TreeNode{
			Val:   p[0],
			Left:  left,
			Right: right,
		}

		return &root
	}

	return ct(i, p)
}

// Helper function to build the tree [4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2].
func buildTree() *TreeNode {
	root := &TreeNode{Val: 4}

	root.Left = &TreeNode{Val: -7}

	root.Right = &TreeNode{Val: -3}
	root.Right.Left = &TreeNode{Val: -9}
	root.Right.Right = &TreeNode{Val: -3}

	root.Right.Left.Left = &TreeNode{Val: 9}
	root.Right.Left.Right = &TreeNode{Val: -7}

	root.Right.Right.Left = &TreeNode{Val: -4}
	root.Right.Right.Right = &TreeNode{Val: 6}

	root.Right.Left.Left.Left = &TreeNode{Val: -6}
	root.Right.Left.Left.Right = &TreeNode{Val: -6}

	root.Right.Left.Right.Left = &TreeNode{Val: 0}
	root.Right.Right.Left.Right = &TreeNode{Val: 6}

	root.Right.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right.Right = &TreeNode{Val: 9}

	root.Right.Right.Right.Right.Left = &TreeNode{Val: -1}
	root.Right.Right.Right.Right.Right = &TreeNode{Val: -4}

	root.Right.Right.Right.Right.Right.Right = &TreeNode{Val: -2}

	return root
}

// Helper function to print the tree (Pre-order traversal).
func printPreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printPreOrder(root.Left)
	printPreOrder(root.Right)
}

// Helper function to print the tree (In-order traversal).
func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInOrder(root.Right)
}

// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
func LC297() {
	fmt.Println("\nRunning Serialize and Deserialize Binary Tree (LC297) testcase...")
	// Build the tree [1,2,3,null,null,4,5].
	tree := buildTree()

	// Serialize the tree.
	codec := Constructor()
	serializedData := codec.serialize(tree)
	fmt.Println("Serialized Tree:", serializedData)

	// Deserialize the serialized tree.
	deserializedTree := codec.deserialize(serializedData)
	fmt.Print("Deserialized Tree (Pre-order): ")
	printPreOrder(deserializedTree)
	fmt.Print("\nDeserialized Tree (In-order): ")
	printInOrder(deserializedTree)
	fmt.Println()
}
