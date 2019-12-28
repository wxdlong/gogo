package rbtree

const (
	left  = 1
	right = 0
)

type stacks struct {
	head *stack
	size int
}

type stack struct {
	data *Node
	tag  int
	next *stack
}

func newStack() *stacks {
	return &stacks{}
}

func (s *stacks) push(n *Node) {
	s.head = &stack{
		data: n,
		next: s.head,
	}
	s.size++
}

func (s *stacks) pushLeft(n *Node) {
	s.head = &stack{
		data: n,
		tag:  left,
		next: s.head,
	}
	s.size++
}

func (s *stacks) pushRight(n *Node) {
	s.head = &stack{
		data: n,
		tag:  right,
		next: s.head,
	}
	s.size++
}

func (s *stacks) popTag() (int, *Node) {
	st := s.head
	s.head = s.head.next
	s.size--
	return st.tag, st.data
}

func (s *stacks) pop() *Node {
	st := s.head
	s.head = s.head.next
	s.size--
	return st.data
}

func (s *stacks) isEmpty() bool {
	return s.size == 0
}
