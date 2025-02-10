package dsa

import (
	"fmt"
	"strconv"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() (string, error) {
	if s.IsEmpty() {
		return "", fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func (s *Stack) Clear() {
	s.items = []string{}
}

func (s *Stack) ForEach(f func(string)) {
	for _, item := range s.items {
		f(item)
	}
}

func NewStack() *Stack {
	return &Stack{}
}

func NewStackFromStringSlice(items []string) *Stack {
	return &Stack{items: items}
}

func NewStackFromIntSlice(items []int) *Stack {
	stack := &Stack{}
	for _, item := range items {
		stack.Push(strconv.Itoa(item))
	}
	return stack
}
