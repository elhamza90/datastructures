package stack

import "strings"
//import "fmt"

type datatype string

type Stack []datatype

func New(elements ...datatype) Stack {
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
		str += string(v) + " ";
	}
	return strings.TrimSpace(str)
}

func (s Stack) Length() int {
	return len(s);
}

func (s *Stack) Push(data datatype) {
	*s = append(*s, data);
}

func (s *Stack) Pop() (datatype, bool) {
	if (*s).isEmpty() {
		return "", false;
	}
	last := (*s)[len(*s)-1];
	*s = (*s)[:len(*s)-1];
	return last, true;
}