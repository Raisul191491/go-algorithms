package bst

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Depth int
}

func CreateNode(value, Depth int) *BinaryTreeNode {
	return &BinaryTreeNode{Value: value, Depth: Depth}
}

func (n *BinaryTreeNode) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = CreateNode(value, n.Depth+1)
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = CreateNode(value, n.Depth+1)
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *BinaryTreeNode) PrintTree(node *BinaryTreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Printf("├L── ")
	} else {
		fmt.Printf("└R── ")
	}
	fmt.Printf("%d\n", node.Value)

	node.PrintTree(node.Left, prefix+(map[bool]string{true: "│   ", false: "    "}[isLeft]), true)
	node.PrintTree(node.Right, prefix+(map[bool]string{true: "│   ", false: "    "}[isLeft]), false)
}

// InOrderTraversal performs in-order traversal of the BST and appends values to a slice
func (n *BinaryTreeNode) InOrderTraversal(node *BinaryTreeNode, result *[]int) {
	if node == nil {
		return
	}

	// Traverse left subtree
	node.InOrderTraversal(node.Left, result)

	// Visit root node
	*result = append(*result, node.Value)

	// Traverse right subtree
	node.InOrderTraversal(node.Right, result)
}

func (n *BinaryTreeNode) SortBST(root *BinaryTreeNode) []int {
	var result []int
	root.InOrderTraversal(root, &result)
	return result
}
