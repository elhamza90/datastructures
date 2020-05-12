package queue

import (
	"strings"
	"strconv"
)
//import "fmt"


type Queue []int

func New(elements ...int) Queue {
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
		str += strconv.Itoa(v) + " ";
	}
	return strings.TrimSpace(str)
}

func (q Queue) Length() int {
	return len(q);
}

func (q *Queue) Push(data int) {
	*q = append(*q, data);
}

func (q *Queue) Pop() (int, bool) {
	if (*q).isEmpty() {
		return 0, false;
	}
	first := (*q)[0];
	*q = (*q)[1:len(*q)];
	return first, true;
}