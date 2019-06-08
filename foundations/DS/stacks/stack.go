package stacks

//Stack - Stack Data Structure
type Stack struct {
	stack []int
}

//New - Initializes the stack
func New() *Stack {
	return &Stack{stack: []int{}}
}

//Push - Inserting an item into the stack
func (s *Stack) Push(item int) {
	s.stack = append(s.stack, item)
}

//Pop - Deleting an item from the stack
func (s *Stack) Pop() int {
	//Remove the last item from the slice
	if len(s.stack) == 0 {
		return -1
	}
	var lastItem = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return lastItem

}
