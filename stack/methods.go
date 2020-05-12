package stack

import (
	"strings"
	"strconv"
)
//import "fmt"


type Stack []int

func New(elements ...int) Stack {
	s := make(Stack, len(elements));
	for i, v := range elements {
		s[i] = v;
	}
	return s;
}

func (s Stack) isEmpty() bool {
	return len(s) == 0;
}

func (s Stack) String() string {
	str := "";
	for _, v := range s {
		str += strconv.Itoa(v) + " ";
	}
	return strings.TrimSpace(str)
}

func (s Stack) Length() int {
	return len(s);
}

func (s *Stack) Push(data int) {
	*s = append(*s, data);
}

func (s *Stack) Pop() (int, bool) {
	if (*s).isEmpty() {
		return 0, false;
	}
	last := (*s)[len(*s)-1];
	*s = (*s)[:len(*s)-1];
	return last, true;
}