package lnklist

import "testing"

func TestAppend(t *testing.T) {

	// Create the Linked List
	my_linked_list := LinkedList {}

	// Test append first element
	my_linked_list.Append("Data of First Node")
	if my_linked_list.length != 1 {
		t.Error("Error Appending first element. Length after appending is", my_linked_list.length)
	}

	// Test append second element
	my_linked_list.Append("Data of First Node");
	if my_linked_list.length != 2 {
		t.Error("Error Appending second element. Length after appending is", my_linked_list.length)
	}
}

func TestPrint(t *testing.T) {

	my_linked_list := LinkedList {
		head : &Node {
			data: "1",
			next: &Node {
				data: "2",
				next: nil,
			},
		},
		length: 2,
	}

	my_linked_list.Print()
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
		length: 2,
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
		length: 3,
	};

	last := my_linked_list.Pop();
	if last.data != "3" && my_linked_list.length == 2 {
		t.Error("Error Poping last element")
	}
}
