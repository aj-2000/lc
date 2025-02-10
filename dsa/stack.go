package dsa

import (
	"fmt"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty")
	}
	s.items = s.items[:len(s.items)-1]
	return nil
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
