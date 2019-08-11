package datastructure

// Stack is a stack of chars
type Stack []interface{}

// Push pushes a new char onto stack
func (s Stack) Push(v interface{}) Stack {
	return append(s, v)
}

// Pop pops the top of the stack
func (s Stack) Pop() (Stack, interface{}) {
	l := len(s)

	if l == 0 {
		return s, 0
	}

	res := s[l-1]
	s = s[:l-1]

	return s, res
}
