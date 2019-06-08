package stacks

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type DataType generic.Type

//Stack - Stack Data Structure
type Stack struct {
	stack []DataType
	lock  sync.RWMutex
}

//New - Initializes the stack
func New() *Stack {
	return &Stack{stack: []DataType{}}
}

//Push - Inserting an item into the stack
func (s *Stack) Push(item int) {
	s.lock.Lock()
	s.stack = append(s.stack, item)
	s.lock.Unlock()
}

//Pop - Deleting an item from the stack
func (s *Stack) Pop() DataType {
	//Remove the last item from the slice
	if len(s.stack) == 0 {
		return -1
	}
	s.lock.Lock()
	var lastItem = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.lock.Unlock()
	return lastItem

}
