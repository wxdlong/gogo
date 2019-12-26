package rbtree

import (
	"fmt"
	"testing"
)

func TestRBTree_Insert(t *testing.T) {
	rb := NewTree(3)
	rb.Insert(4)
	rb.Insert(14)
	rb.Insert(15)
	rb.Insert(8)
	rb.Insert(11)
	rb.Insert(1)
	rb.Insert(2)
	rb.Insert(5)
	rb.Insert(22)
	fmt.Println("\n前序遍历")
	rb.root.preTraverse()
	fmt.Println("\n中序遍历")

	rb.root.inTraverse()
	fmt.Println("\n后序遍历")

	rb.root.postTraverse()
}
