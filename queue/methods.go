package queue

import (
	"strings"
	"strconv"
)
//import "fmt"

type datatype int

type Queue []datatype

func New(elements ...datatype) Queue {
	q := make(Queue, len(elements));
	for i, v := range elements {
		q[i] = v;
	}
	return q;
}

func (q Queue) isEmpty() bool {
	return len(q) == 0;
}

func (q Queue) String() string {
	str := "";
	for _, v := range q {
		str += strconv.Itoa(int(v)) + " ";
	}
	return strings.TrimSpace(str)
}

func (q Queue) Length() int {
	return len(q);
}

func (q *Queue) Push(data datatype) {
	*q = append(*q, data);
}

func (q *Queue) Pop() (datatype, bool) {
	if (*q).isEmpty() {
		return 0, false;
	}
	first := (*q)[0];
	*q = (*q)[1:len(*q)];
	return first, true;
}