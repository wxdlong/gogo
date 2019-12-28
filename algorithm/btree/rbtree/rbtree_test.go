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
	fmt.Println("\n前序遍历: 递归")
	rb.root.preTraverse()
	fmt.Println("\n前序遍历: 非递归")
	rb.preTraverse2()

	fmt.Println("\n前序遍历: 非递归3")
	rb.preTraverse3()

	fmt.Println("\n中序遍历: 递归")

	rb.root.inTraverse()
	fmt.Println("\n中序遍历: 非递归")

	rb.inTraverse()
	fmt.Println("\n后序遍历")

	rb.root.postTraverse()

	fmt.Println("\n后序遍历2")
	rb.postTraverse()

	fmt.Println("\n后序遍历3")
	rb.postTraverse2()

	fmt.Println("\n后序遍历4")
	rb.postTraverse3()

	fmt.Println("\n层次遍历")
	rb.levelTraverse()

}
