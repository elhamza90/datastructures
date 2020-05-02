package stack


import "testing"
//import "fmt"

func TestString(t *testing.T) {
	s := New("1","2","3","4","5");
	if s.String() != "1 2 3 4 5" {
		t.Error("Error in String conversion of stack");
	}
}

func TestLength(t *testing.T) {
	s := New("1", "2", "3");
	if s.Length() != 3 {
		t.Error("Error calculating length of Stack");
	}
}

func TestAppend(t *testing.T) {
	s:= New("1", "2");
	s.Push("3");
	if s.Length() != 3 {
		t.Error("Error Appending Element to stack");
	}
}

func TestPop(t *testing.T) {
	s := New("1", "2", "3", "4");
	last, ok := s.Pop();
	if last != "4" || s.Length() != 3 || ok == false {
		t.Error("Error in Poping from stack.")
	}

	s = New();
	_, ok = s.Pop();
	if ok == true {
		t.Error("Error Poping empty stack returns ok is true")
	}
}