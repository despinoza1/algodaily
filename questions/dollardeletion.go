package questions

// Stack is a stack of chars
type Stack []rune

// Push pushes a new char onto stack
func (s Stack) Push(v rune) Stack {
	return append(s, v)
}

// Pop pops the top of the stack
func (s Stack) Pop() (Stack, rune) {
	l := len(s)

	if l == 0 {
		return s, 0
	}

	res := s[l-1]
	s = s[:l-1]

	return s, res
}

func getProccessedStr(str string) string {
	stk := make(Stack, 0)

	for _, char := range str {
		if char == '$' {
			stk, _ = stk.Pop()
		} else {
			stk = stk.Push(char)
		}
	}

	res := ""
	var tmp rune
	for i := 0; i < len(stk); i++ {
		stk, tmp = stk.Pop()
		res += string(tmp)
	}

	return res
}

// DollarSignDeletion returns if after remove every character
// preceding a $ if all elements in array are equivalent
func DollarSignDeletion(strs []string) bool {
	tmp := getProccessedStr(strs[0])

	for _, str := range strs {
		if tmp != getProccessedStr(str) {
			return false
		}
	}

	return true
}
