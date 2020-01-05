package beibao

//假定有N=4件商品，有一个背包可以装9Kg，怎样挑选物品使价值最大？
//      重量(kg)  价值(元)
//    A  3         66
//    B  2         40
//    C  5         95
//    D  4         40

type node struct {
	cv, rv int //当前背包物品价值，剩余物品价值
	rw, id int //剩余容量， 物品ID
	xs     []bool
}

type good struct {
	name          string
	weight, value int
}

type tree struct {
	root *node
}

type queues struct {
	head *queue
	tail *queue
	size int
}

type queue struct {
	value *node
	next  *node
}
