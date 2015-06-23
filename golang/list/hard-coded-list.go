/* link list example using container/list package */

package main

import (
	"fmt"
	"errors"
)

type Value struct {
	Name string
	Age  int
}

type Node struct {
	Value
	next, prev *Node
}

type List struct {
	head, tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Push(v Value) *List {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n    // First node
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n

	return l
}

func (l *List) Find(name string) *Node {
	found := false
	var ret *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.Value.Name == name {
			found = true
			ret = n
		}
	}

	return ret
}

func (l *List) Delete (name string) bool {
	success := false
	node2del := l.Find (name)
	if node2del != nil {
		fmt.Println ("Delete - Found: ", name)
		prev_node := node2del.prev
		next_node := node2del.next
		//remove this node
		prev_node.next = next_node
		next_node.prev = prev_node
		success = true
	}

	return success
}

func (l *List) Print_list () {
	for n := l.First(); n != nil; n = n.Next() {
		fmt.Printf ("%v\n", n.Value)
	}
}

var errEmpty = errors.New ("Error - List is empty")


func main () {
	l := new(List)  // Create doubly  link list
	if l.head != nil {
		fmt.Println ("head is not nil")
		}
	l.Push (Value{Name: "A", Age: 23})
	l.Push (Value{Name: "B", Age: 23})
	l.Push (Value{Name: "C", Age: 23})

	l.Print_list()

	temp := l.Delete ("B")
	if temp {
		fmt.Println ("node deleted successfully")
	}

	l.Print_list()
		
}
