package lnklist

import "testing"

func TestPrint(t *testing.T) {

	my_linked_list := LinkedList {
		head : &Node {
			data: "1",
			next: &Node {
				data: "2",
				next: nil,
			},
		},
	}

	my_linked_list.Print()
}

func TestLength(t *testing.T) {
	my_linked_list := LinkedList {
		head: &Node {
			data: "1",
			next: &Node {
				data: "2",
				next: &Node {
					data: "3",
					next: &Node {
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

func TestAppend(t *testing.T) {

	// Create the Linked List
	my_linked_list := LinkedList {}

	// Test append first element
	my_linked_list.Append("Data of First Node")
	if my_linked_list.Length() != 1 {
		t.Error("Error Appending first element. Length after appending is", my_linked_list.Length())
	}

	// Test append second element
	my_linked_list.Append("Data of First Node");
	if my_linked_list.Length() != 2 {
		t.Error("Error Appending second element. Length after appending is", my_linked_list.Length())
	}
}
func TestContains(t *testing.T) {

	my_linked_list := LinkedList {
		head : &Node {
			data: "1",
			next: &Node {
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
		head : &Node {
			data: "1",
			next: &Node {
				data: "2",
				next: &Node {
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
