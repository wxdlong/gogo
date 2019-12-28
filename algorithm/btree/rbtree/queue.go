package rbtree

type queues struct {
	head *queue
	tail *queue
	size int
}

type queue struct {
	data *Node
	next *queue
}

func newQueue() *queues {
	return &queues{}
}

func (q *queues) offer(data *Node) {
	qs := &queue{data: data}
	if q.tail == nil {
		q.tail = qs
		q.head = q.tail
	} else {
		q.tail.next = &queue{data: data}
		q.tail = q.tail.next
	}
	q.size++
}

func (q *queues) poll() (n *Node) {

	n = q.head.data
	q.head = q.head.next
	q.size--
	if q.size == 0 {
		q.tail = nil
	}
	return
}

func (q *queues) isEmpty() bool {
	return q.size == 0
}
