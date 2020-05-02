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

func (s *Stack) Append(data datatype) {
	*s = append(*s, data);
}

func (s *Stack) Pop() datatype {
	last := (*s)[len(*s)-1];
	*s = (*s)[:len(*s)-1];
	return last;
}