package rbtree

import (
	"log"
	"testing"
)

func TestHello(t *testing.T) {
	q := newQueue()
	q.offer(&Node{data: 5})
	data := q.poll().data
	if data != 5 {
		log.Fatal("queue poll data should be 5 not ", data)
	}

	q.offer(&Node{data: 9})
	q.offer(&Node{data: 10})
	q.offer(&Node{data: 11})
	q.offer(&Node{data: 8})

	data = q.poll().data
	if data != 9 {
		log.Fatal("queue poll data should be 9 not ", data)
	}
}
