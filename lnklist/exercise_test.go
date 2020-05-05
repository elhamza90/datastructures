package lnklist

import "testing"


func TestNew(t *testing.T) {
	l := New("1", "2", "3");
	if l.head.data != "1" || l.head.next.data !="2" || l.head.next.next.data != "3" {
		t.Error("Error creating new Linked list with initial values");
	}
}

func TestString(t *testing.T) {

	my_linked_list := LinkedList {
		head : &node {
			data: "1",
			next: &node {
				data: "2",
				next: &node {
					data: "3",
					next: &node {
						data: "4",
					},
				},
			},
		},
	}

	res := my_linked_list.String();
	if res != "1 -> 2 -> 3 -> 4" {
		t.Error("Error Formating Linked List to String");
	}
}

func TestLength(t *testing.T) {
	my_linked_list := LinkedList {
		head: &node {
			data: "1",
			next: &node {
				data: "2",
				next: &node {
					data: "3",
					next: &node {
						data: "4",
					},
				},
			},
		},
	};

	if my_linked_list.Length() != 4 {
		t.Error("Error calculating Length of list!")
	}
}

func TestPush(t *testing.T) {

	// Create the Linked List
	my_linked_list := LinkedList {}

	// Test append first element
	my_linked_list.Push("Data of First node")
	if my_linked_list.Length() != 1 {
		t.Error("Error Pushing first element. Length after appending is", my_linked_list.Length())
	}

	// Test append second element
	my_linked_list.Push("Data of First node");
	if my_linked_list.Length() != 2 {
		t.Error("Error Pushing second element. Length after appending is", my_linked_list.Length())
	}
}
func TestContains(t *testing.T) {

	my_linked_list := LinkedList {
		head : &node {
			data: "1",
			next: &node {
				data: "2",
				next: nil,
			},
		},
	}
	if my_linked_list.Contains("2") != true {
		t.Error("Could not find an element that exists.")
	}

	if my_linked_list.Contains("22") != false {
		t.Error("Found an element that does not exists.")
	}
	
}

func TestPop(t *testing.T) {
	my_linked_list := LinkedList {
		head : &node {
			data: "1",
			next: &node {
				data: "2",
				next: &node {
					data: "3",
					next: nil,
				},
			},
		},
	};

	last := my_linked_list.Pop();
	if last.data != "3" && my_linked_list.Length() == 2 {
		t.Error("Error Poping last element")
	}
}


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

func TestAppend(t *testing.T) {
	l := New("1", "2", "3", "4", "5");
	l2 := New("6", "7", "8");
	l.Append(l2);
	if l.Length() != 8 {
		t.Error("Error Exercise Append List");
	}
}

func TestReverse(t *testing.T) {
	l := New("1", "2", "3", "4");
	l.Reverse();
	if l.head.data != "4" || l.head.next.data != "3" || l.head.next.next.data != "2" || l.head.next.next.next.data != "1" {
		t.Error("Error Exercise Reversing List");
	}
}

func TestRemoveNth(t *testing.T) {
	l := New("1", "2", "3", "4");
	l.RemoveNth(1);
	if l.head.next.data != "3" {
		t.Error("Error Exercise Remove Nth");
	}
}