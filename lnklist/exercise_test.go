package lnklist

import "testing"

func TestCount(t *testing.T) {
	l := New("1", "2", "3", "4", "5", "2", "4", "2");
	if l.Count("2") != 3 {
		t.Error("Error Counting times element appears in list");
	}
}

func TestGetNth(t *testing.T) {
	l := New("1", "2", "3", "4", "5", "2", "4", "2");
	ret, found := l.GetNth(4);
	if ret != "5" || found != true {
		t.Error("Error Exercise GetNth")
	}
}

func TestInsertNth(t *testing.T) {
	l := New("1", "2", "3", "4", "5", "2", "4", "2");
	oldLength := l.Length();
	done := l.InsertNth("11", 3); // Insert "11" in the third index

	if done != true || l.head.next.next.next.data != "11" || l.Length() != oldLength+1 {
		t.Error("Error InsertNth ")
	}	
}