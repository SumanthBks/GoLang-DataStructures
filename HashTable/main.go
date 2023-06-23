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

// Arraysize
const Arraysize = 10

// hashtable
type HashTable struct {
	array [Arraysize]*buckets
}

// Buckets
type buckets struct {
	head *linkedlist
}

// linkedlist
type linkedlist struct {
	key  string
	next *linkedlist
}

// insert into array and bucket
func (h *HashTable) insert(key string) {
	index := hash(key)
	h.array[index].bucketInsert(key)
}

// insert into the bucket and into the linked list
func (b *buckets) bucketInsert(k string) {
	// TODO: check the below linked list structure and understand
	newNode := &linkedlist{key: k}
	newNode.next = b.head
	b.head = newNode
}

func hash(key string) int {
	hashSum := 0
	for _, k := range key {
		hashSum += int(k)
	}
	return hashSum % Arraysize
}

func initialiseBucket() *HashTable {
	ht := &HashTable{}
	for i := range ht.array {
		ht.array[i] = &buckets{}
	}
	return ht
}

func main() {
	ht := initialiseBucket()
	fmt.Println(ht)
	ht.insert("SUMANTH")
}
