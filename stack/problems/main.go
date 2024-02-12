package main

import "fmt"

type Stack struct {
	stack []rune
	i     int
}

func (s *Stack) remove(val rune) bool {
	if len(s.stack) == 0 {
		return false
	}
	if val == '}' {
		if s.stack[len(s.stack)-1] == '{' {
			return true
		}
	}
	return false
}

func main() {
	input := "{{{{}}}"
	var s Stack
	s.i = -1
	for _, val := range input {
		if val == '{' {
			s.stack = append(s.stack, val)
			s.i++
			continue
		} else {
			if !s.remove(val) {
				fmt.Println("not satisfied")
				break
			}
			s.stack = s.stack[:s.i]
			s.i--
		}
		

	}
if len(s.stack)!=0{
	fmt.Println("err")
}
}
