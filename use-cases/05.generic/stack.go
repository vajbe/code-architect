package main

import "fmt"

type Stack[T any] struct {
	elems    []T
	capacity int
}

func (s *Stack[T]) Push(elem T) bool {
	if len(s.elems) == s.capacity {
		return false
	}
	s.elems = append(s.elems, elem)
	return true
}

func (s *Stack[T]) Pop() *T {

	if len(s.elems) == 0 {
		return nil
	}
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return &elem
}

func (s *Stack[T]) Peek() *T {
	if len(s.elems) == 0 {
		return nil
	}
	return &s.elems[len(s.elems)-1]
}

func StackExample() {
	stack := &Stack[int]{capacity: 10}

	stack.Push(10)
	// stack.Pop()
	fmt.Println(*stack.Peek())
}
