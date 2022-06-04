package solution1

func isValid(s string) bool {
	stk := &stack{
		container: make([]byte, 0),
	}

	for i := range s {
		switch s[i] {
		case ')':
			if stk.Pop() == '(' {
				continue
			}
			return false

		case '}':
			if stk.Pop() == '{' {
				continue
			}
			return false
		case ']':
			if stk.Pop() == '[' {
				continue
			}
			return false
		default:
			stk.Push(s[i])
		}
	}

	return stk.Empty()
}

type stack struct {
	container []byte
}

func (s *stack) Push(b byte) {
	s.container = append(s.container, b)
}

func (s *stack) Pop() byte {
	if s.Empty() {
		return 0
	}
	e := s.container[len(s.container)-1]
	s.container = s.container[:len(s.container)-1]
	return e
}

func (s stack) Top() byte {
	if s.Empty() {
		return 0
	}
	return s.container[len(s.container)-1]
}

func (s stack) Empty() bool {
	return len(s.container) == 0
}
