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
