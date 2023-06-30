package main

import "fmt"

type Node struct {
	key  int
	next *Node
}

type Linkedlist struct {
	head   *Node
	length int
}

func (l *Linkedlist) appendList(n *Node) {
	currentNode := l.head
	n.next = currentNode
	l.head = n
	l.length++
}

func main() {
	ll := &Linkedlist{}
	node1 := &Node{key: 10}
	node2 := &Node{key: 20}
	ll.appendList(node1)
	ll.appendList(node2)
	fmt.Println(ll.length)
}
