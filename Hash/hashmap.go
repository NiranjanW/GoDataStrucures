package main

import (
	"fmt"
)

//ArraySize is the size of the hash table Array
const Arraysize = 7

//HashTable will hold an array
type HashTable struct {
	array [Arraysize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

//Int will create abucket into hashtable
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

//Insert

//Insert will take a key and add it to the hastable array ( each element in array is a bucket)

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search will take a key and return true if key in HashTable

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

//Delete will take a key and delete it from the HT

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(k string) {
	if !b.search((k)) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("already exits")
	}

}

//search will take a key and return true if the bucket has the key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next.next
		}
	}

}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % Arraysize
}

func main() {
	testHashTable := Init()
	list := []string{
		"RANDY",
		"MANDY",
		"KENNY",
		"ERIC",
	}
	for _, v := range list {
		testHashTable.Insert(v)
	}
	fmt.Println(testHashTable.array)
	// fmt.Println(hash("RANDY"))

	// testBucket := &bucket{}
	// testBucket.insert("RANDY")
	// testBucket.delete(("RANDY"))
	// fmt.Println(testBucket.search("RANDY"))
}
