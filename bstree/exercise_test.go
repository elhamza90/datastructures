package bstree

import "testing"
import "fmt"

func TestNew(t *testing.T) {
	tr := New();
	if tr.root != nil {
		t.Error("Expecting empty tree with New");
	}
}

func TestAppend(t *testing.T) {
	tr := New();
	tr.Insert(44);
	if tr.root.data != 44 {
		t.Error("Error Inserting first value");
	}

	tr.Insert(100);
	if tr.root.right.data != 100 {
		t.Error("Error inserting greater elements in right side of root")
	}

	tr.Insert(10);
	if tr.root.left.data != 10 {
		t.Error("Error inserting greater elements in left side of root")
	}

	tr.Insert(99);
	if tr.root.right.left.data != 99 {
		t.Error("Error inserting third element greater than root smaller than second in right position")
	}
}

func TestString(t *testing.T) {
	tr := New();
	tr.Insert(8);
	tr.Insert(3);
	tr.Insert(10);
	tr.Insert(1);
	tr.Insert(6);
	tr.Insert(14);
	tr.Insert(13);
	fmt.Print(tr.String());
}

func TestFind(t *testing.T) {
	tr := New();
	tr.Insert(8);
	tr.Insert(3);
	tr.Insert(10);
	tr.Insert(1);
	tr.Insert(6);
	tr.Insert(14);
	tr.Insert(13);
	found, _ := tr.Find(14);
	if found == false {
		t.Error("Could not find element in existing in tree");
	}
	found, _ = tr.Find(1000);
	if found == true {
		t.Error("Found Element not existing in tree");
	}
}

func TestFindInOrder(t *testing.T) {
	tr := New();
	tr.Insert(8);
	tr.Insert(3);
	tr.Insert(10);
	tr.Insert(1);
	tr.Insert(6);
	tr.Insert(4);
	tr.Insert(5);
	tr.Insert(14);
	tr.Insert(13);
	tr.Insert(12);
	//fmt.Println(tr.String());
	res := tr.root.right.findInOrder();
	if res.data != 12 {
		t.Error("Error Exercise FindInOrder BST");
	}
	res = tr.root.right.right.findInOrder();
	if res != nil {
		t.Error("Error Exercise FindInOrder BST");
	}
}

func TestDelete(t *testing.T) {
	tr := New();
	tr.Insert(8);
	tr.Insert(3);
	tr.Insert(10);
	tr.Insert(1);
	tr.Insert(6);
	tr.Insert(4);
	tr.Insert(5);
	tr.Insert(14);
	tr.Insert(13);
	done := tr.Delete(13);
	if done != true || tr.root.right.right.left != nil {
		t.Error("Error Exercise Delete BST Leaf Node");
	}
	tr.Insert(17);
	tr.Insert(12);
	tr.Insert(13);
	tr.Insert(15);
	tr.Insert(30);
	tr.Insert(23);
	done = tr.Delete(14);
	if done != true || tr.root.right.right.data != 15 {
		t.Error("Error Exercise Delete BST Node")
	}
}