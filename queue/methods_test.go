package queue


import "testing"
//import "fmt"

func TestString(t *testing.T) {
	q := New(1,2,3,4,5);
	if q.String() != "1 2 3 4 5" {
		t.Error("Error in String conversion of queue");
	}
}

func TestLength(t *testing.T) {
	q := New(1, 2, 3);
	if q.Length() != 3 {
		t.Error("Error calculating length of Queue");
	}
}

func TestAppend(t *testing.T) {
	q:= New(1, 2);
	q.Push(3);
	if q.Length() != 3 {
		t.Error("Error Appending Element to queue");
	}
}

func TestPop(t *testing.T) {
	q := New(1, 2, 3, 4);
	first, ok := q.Pop();
	if first != 1 || q.Length() != 3 || ok == false {
		t.Error("Error in Poping from queue.")
	}

	q = New();
	_, ok = q.Pop();
	if ok == true {
		t.Error("Error Poping empty queue returns ok is true")
	}
}