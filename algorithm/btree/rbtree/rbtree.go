package rbtree

const (
	RED   = true
	BLACK = false
)

type RBTree struct {
	root *Node
}

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  bool
	data   int
}

//NewTree 返回一棵新红黑树。
//新树根节点颜色为黑色
func NewTree(data int) *RBTree {
	root := New(data)
	root.color = BLACK
	return &RBTree{root: root}
}

//New return a Node with init color RED.
//New insert Node color should be RED.
func New(data int) *Node {
	return &Node{data: data, color: RED}
}

//Insert data into tree.
func (t *RBTree) Insert(data int) {
	n := t.root
	d := New(data)
	for n != nil {
		if data > n.data {
			//如果值>当前节点,则插入右子树。
			//1. 右子树不空时，以右子节点循环继续。
			//2. 右子树为空时，直接插入
			if n.right == nil {
				n.right = d
				d.parent = n
				break
			} else {
				n = n.right
			}
		} else {
			//如果值<=当前节点,则插入右子树。
			//1. 左子树不空时，以左子节点循环继续。
			//2. 左子树为空时，直接插入
			if n.left != nil {
				n = n.left
			} else {
				n.left = d
				d.parent = n
				break
			}
		}
	}

	t.fixUp(d) //新插入节点后可能会破坏平衡
}

//fixUp 重新平衡红黑树。
//1. 插入节点的父节点和其叔叔节点（祖父节点的另一个子节点）均为红色的；
//   此时祖父节点肯定存在，且一定是黑色。但不确定其是左子树，还是右子树。
//     a. 将父叔节点变色->黑色，将祖父节点变色->红色。
//     b. 将祖父节点设置为当前节点，此时情况变成2。其叔叔节点必定有且为黑
//2. 插入节点的父节点是红色，叔叔节点是黑色，且插入节点是其父节点的右子节点；
//     a. 以父节点为支点左旋，因为需要将当前节点上移一级。
//     b. 交换父子
//3. 插入节点的父节点是红色，叔叔节点是黑色，且插入节点是其父节点的左子节点。
//     a. 将父节点变黑，祖父变红。
//     b. 以祖父节点为支点，右旋。
func (t *RBTree) fixUp(n *Node) {
	var p, g, u *Node   //父，祖，叔节点

	//父节点为红色，祖父节点必定存在
	for p = n.parent; p != nil && p.color == RED; p = n.parent {
		g = p.parent
		//如果父节点是其左子树
		if p == g.left {
			u = g.right //获取叔父节点

			//case1: 叔父节点也为红, 祖叔父皆变黑色
			if u != nil && u.color == RED {
				p.color = BLACK
				u.color = BLACK
				g.color = RED
				n = g
				continue
			}

			//case2:
			if n == p.right {
				t.leftRotate(p)
				n, p = p, n //调换父子节点
			}

			//case3:
			p.color = BLACK
			g.color = RED
			t.rightRotate(g)
		} else { //若父节点是祖父节点的右子节点，怀上面完全相反
			u = g.left

			//case1: 叔叔节点也是红色
			if u != nil && u.color == RED {
				p.color = BLACK
				u.color = BLACK
				g.color = RED
				n = g
				continue
			}

			//case2: 叔叔节点是黑色
			if n == p.left {
				t.rightRotate(p)
				n, p = p, n
			}

			p.color = BLACK
			g.color = RED
			t.leftRotate(g)
		}
	}

	t.root.color = BLACK    //将根节点设置红色。
}

func (n *Node) Delete() {

}

func (n *Node) Search(data int) {

}

//左旋, 右子树过深，需要调整高度，将Y提高一级，且相对原位置左移，故称左旋。
// 父X降一级为Y左子树，且由于原来B值肯定是小于Y，且大于X,这时放在X右子树没问题。
// 想像成有一个钩子将Y拎到X的位置，X自然偏左下降。可视以X为中心作逆时针旋转。

//           <-------- left-rotate
//           ---------> right-rotate

//      |                     |
//      Y                     X
//     /  \                  /  \
//    X    c                a    Y
//   / \                        / \
//  a	 b                     b    c
// 1. 将Y左子节点(b)赋给X的右节点，同时将其父节点设置为X
// 2. 将X的父节点赋值给Y的父节点，并更新其子节点。
// 3. 将X设置为Y的左子树，并将其父设置为Y.


func (t RBTree) leftRotate(x *Node) {
	y := x.right
	if y.left != nil {
		y.left.parent = x
	}
	x.right = y.left


	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}

	y.left = x
	x.parent = y
}

//右旋 左子树过深，需要调整高度，将降一级，且相对原位置右移，故称右旋。
//子X升一级，且由于原来B值肯定是小于Y，且大于X,这时放在Y左子树没问题。
// 想像成有一个钩子将X拎Y到的位置,Y自然偏右下落。可视以X为中心作顺时针旋转。
//           <-------- left-rotate
//           ---------> right-rotate

//      |                     |
//      Y                     X
//     /  \                  /  \
//    X    c                a    Y
//   / \                        / \
//  a	 b                     b    c
// 1. 将x右子节点(b)赋给Y的左节点，同时将其父节点设置为Y
// 2. 将y的父节点赋值给x的父节点，并更新其子节点。
// 3. 将Y设置为X的右子树，并将其父设置为X.
func (t *RBTree) rightRotate(x *Node) {
	y := x.parent
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}

	x.right = y
	y.parent = x

}

//change color.
func (n *Node) discolor(c bool) {
	n.color = c
}

//
//func (t *RBTree) String() string{
//	n:=t.root
//	s:=""
//	for i=20; i>0; i++ {
//		s+=" "
//	}
//	for i=20; i>0; i++ {
//		s+=" "
//	}
//
//
//}