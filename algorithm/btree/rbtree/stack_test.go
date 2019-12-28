package rbtree

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := newStack()
	s.push(&Node{data: 4})
	s.push(&Node{data: 5})
	s.push(&Node{data: 1})
	s.push(&Node{data: 13})
	s.push(&Node{data: 3})
	s.push(&Node{data: 9})
	s.push(&Node{data: 8})

	for !s.isEmpty() {
		fmt.Print(s.pop().data, " ")
	}
}
