package bstree

import (
	"strconv"
	"strings"
)


type node struct {
	data int
	left *node
	right *node
};

type bst struct {
	root *node
};


// New returns a pointer to an empty tree
func New() *bst {
	return &bst {};
}

// to2D prints the tree sideways in 2D
func to2D(root *node, space int) string {
	if root == nil {
		return "";
	}
	s := ""
	space += 10;

	s += to2D(root.right, space);

	for i:=0; i < space; i++ {
		s += " ";
	}
	s += strconv.Itoa(root.data) + "\n";

	s += to2D(root.left, space);
	return s;
}

// String is wrapper over to2D
func (tr bst) String() string {
	if tr.root == nil {
		return "";
	}
	return to2D(tr.root, 0);
}

// Insert creates a node and inserts it on the tree.
// It follows iteratively the algorithm:
//  - if value < current node, go to left side
//  - else go to right side
func (tr *bst) Insert(val int) {
	nodeToInsert := node {
			data: val,
	};
	if tr.root == nil {
		tr.root = &nodeToInsert
	} else {
		next := tr.root;
		for {
			if val < next.data {
				if next.left == nil {
					next.left = &nodeToInsert;
					break;
				} else {
					next = next.left;
				}
			} else {
				if next.right == nil {
					next.right = &nodeToInsert;
					break;
				} else {
					next = next.right;
				}
			}
		}
	}
}

// Find searches for a value in the tree and returns
// 	- true and search path to value if value exists
// 	- false and search path if value doesn't exist
func (tr bst) Find(val int) (bool, string) {
	if tr.root == nil {
		return false, "";
	}
	found := false;
	path := "";
	cur := tr.root;
	for cur != nil {
		path += strconv.Itoa(cur.data) + " ";
		if val == cur.data {
			found = true;
			break;
		} else if val < cur.data {
			cur = cur.left;
		} else {
			cur = cur.right;
		}
	}
	return found, strings.TrimSpace(path);
}

// findInOrder finds in-orde successor to the node in subtree.
func (n *node) findInOrder() (s *node) {
	s = nil;
	var cur *node = n;
	var potentialSuccessor *node;
	var successorDiff int = 0;
	for cur != nil {
		if potentialSuccessor == nil {
			if cur.right != nil {
				potentialSuccessor = cur.right;
				successorDiff = potentialSuccessor.data - n.data;
				cur = cur.right;
			}
		} else {
			if cur.data - n.data < successorDiff {
				potentialSuccessor = cur;
				successorDiff = cur.data - n.data;
			}
			cur = cur.left;
		}
	}
	return potentialSuccessor;
}

// Delete deletes a node given its key
func (tr *bst) Delete(k int) (done bool) {
	done = true;
	if tr.root != nil {
		var prev, cur *node = nil, tr.root;
		var isLeftChild bool = false;
		for {
			if cur.data == k {
				isLeftChild = k < prev.data;
				if cur.left == nil && cur.right == nil {
					// Case 1: Leaf
					if isLeftChild {
						prev.left = nil;
					} else {
						prev.right = nil;
					}
				} else {
					// Case 2: subtree

				}
				done = true;
				break;
			} else {
				if k < cur.data {
					if cur.left == nil {
						break;
					} else {
						prev = cur;
						cur = cur.left;
					}
				} else {
					if cur.right == nil {
						break;
					} else {
						prev = cur;
						cur = cur.right;
					}
				}
			}
		}
	}
	return done;
}