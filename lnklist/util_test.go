package lnklist

import "testing"

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
