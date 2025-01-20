package main

import (
	"fmt"
	"go-algorithms/bst"
)

func main() {
	binaryTree := bst.CreateNode(56, 0)

	arr := []int{4, 7, 89, 54, 34, 3}

	for _, v := range arr {
		binaryTree.Insert(v)
	}

	binaryTree.PrintTree(binaryTree, "", false)
	result := binaryTree.SortBST(binaryTree)
	fmt.Println(result)
}
