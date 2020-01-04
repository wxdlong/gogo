package beibao

//假定有N=4件商品，有一个背包可以装9Kg，怎样挑选物品使价值最大？
//      重量(kg)  价值(元)
//    A  3         66
//    B  2         40
//    C  5         95
//    D  4         40

type node struct {
	name            string
	weight          int
	value           int
	left, right     *node
	parent, brother *node
	selected        bool
}

type tree struct {
	root *node
}

type stacks struct {
	head *stack
	size int
}

type stack struct {
	*node
	next *node
}

//
//func NewTree(ns []*node) *tree {
//	t := &tree{}
//
//	t.root = &node{}
//	root := t.root
//	for i := 0; i < len(ns); i++ {
//		n: = ns[i]
//		root.left = ns[i]
//
//		root.right = ns[i]
//	}
//	return t
//}
