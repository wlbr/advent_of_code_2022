package main

type stack struct {
	content string
}

func (s *stack) push(c rune) {
	s.content += string(c)
}

func (s *stack) peek() rune {
	l := len(s.content)
	if l > 0 {
		return rune(s.content[l-1])
	}
	return ' '
}

func (s *stack) pop() rune {
	l := len(s.content)
	if l > 0 {
		last := s.content[l-1]
		s.content = s.content[:len(s.content)-1]
		return rune(last)
	}
	return ' '
}

func (s *stack) empty() bool {
	return len(s.content) == 0
}

func (s *stack) len() int {
	return len(s.content)
}

func (s *stack) reset() {
	s.content = ""
}
