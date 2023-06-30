package main

import "fmt"

// Hash Tables or Hash Map

// Data in hashmap is stored based on Key, Value Pair
// and these Key are stored in array and each array has an bucket
// this bucket contains an linked list which stores the data incase of the hash collision
// Hash map is having O(1) for insert, search, delete
// worst case scenario for hashmap is when the hash collision happens
// and the data is pilled up in the linked list under the bucket
// when the last element in the linked list is searched then performance hit happens

const arraySize = 10

type HashTable struct {
	array [arraySize]*Buckets
}

type Buckets struct {
	length     int
	linkedList *LinkedList
}

type LinkedList struct {
	head *Node
}

type Node struct {
	key  string
	next *Node
}

func (h *HashTable) insert(key string) {
	index := hash(key)
	h.array[index].insertBucket(key)
	fmt.Println(key)
}

func (b *Buckets) insertBucket(key string) {
	if !b.searchBucket(key) {
		// newNode := &LinkedList{key: key} // initilize the linkedlist
		// newNode.next = b.head            // add the head of the linkedlist to the next of the newNode
		// b.head = newNode                 // make the newNode as the head of the linkedlist
		b.linkedList.insertLinkedList(key)
		b.length++
	} else {
		fmt.Println("Key exists: ", key)
	}
}

func (l *LinkedList) insertLinkedList(key string) {
	newNode := &Node{key: key}
	newNode.next = l.head
	l.head = newNode
}

func returnIndex(key string) int {
	return hash(key)
}

func (h *HashTable) delete(key string) {
	index := returnIndex(key)
	h.array[index].deleteKey(key)
}

func (b *Buckets) deleteKey(key string) {
	if b.searchBucket(key) {
		if b.linkedList.head.key == key {
			b.linkedList.head = b.linkedList.head.next // if the head is the target deletion then make head.next as the head
			b.length--
			return
		}

		currentNode := b.linkedList.head
		for currentNode.next != nil {
			if currentNode.next.key == key {
				currentNode.next = currentNode.next.next // if the nest of currenNode is the target deletion then make the next of the next of the list the current.next
			}
			currentNode = currentNode.next
		}
		b.length--
	} else {
		fmt.Println("Key Doesn't exists: ", key)
	}
}

func (h *HashTable) search(key string) bool {
	return h.array[returnIndex(key)].searchBucket(key)
}

func (b *Buckets) searchBucket(key string) bool {
	return b.linkedList.searchLinkedList(key)
}

func (l *LinkedList) searchLinkedList(key string) bool {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (h *HashTable) update(key string, value string) {
	h.array[returnIndex(key)].updateBucket(key, value)
}

func (b *Buckets) updateBucket(key string, value string) {
	if b.searchBucket(key) {
		currentNode := b.linkedList.head
		for currentNode != nil {
			if currentNode.key == key {
				currentNode.key = value
			}
			currentNode = currentNode.next
		}
	}
}

func hash(key string) int {
	sum := 0
	for _, k := range key {
		sum += int(k)
	}
	return sum % arraySize
}

func initialize() *HashTable {
	ht := &HashTable{}
	for arr := range ht.array {
		ht.array[arr] = &Buckets{}
		ht.array[arr].linkedList = &LinkedList{}
	}
	return ht
}

func main() {
	ht := initialize()
	listOfKeys := []string{
		"Sumanth",
		"Yuvan",
		"Krishna",
		"Basetty",
		"Sai",
		"Swathi",
		"Pratyusha",
		"Koushika",
		"Srawan",
		"Anurag",
		"Hari",
		"Rithesh",
		"Anoop",
		"Mahesh",
		"Swathi",
		"Sumalatha",
		"Sudhakar",
		"Roger Federer",
		"Rafeal Nadal",
		"Pikachuuuuu",
	}

	for _, key := range listOfKeys {
		ht.insert(key)
	}
	fmt.Println(ht.search("Pikachuuuuu"))
	ht.delete("Pikachuuuuu")
	fmt.Println(ht.search("Pikachuuuuu"))
	ht.update("Swathi", "Krishna Swathi")

	fmt.Println("length of each bucket")
	for index := range ht.array {
		fmt.Println(index, ht.array[index].length)
	}
}
