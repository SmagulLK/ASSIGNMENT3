package main

import "fmt"

func main() {
	hashtable := Init()
	list := []string{
		"Meirkhan",
		"Diyar",
		"Alkey",
		"Adil",
		"Timur",
		"Asanali",
		"Asan",
	}
	for _, v := range list {
		hashtable.Insert(v)
	}

	//var name string
	//fmt.Println()
	//fmt.Scanln(&name)

	fmt.Println(hashtable.Search("Asan"))
	hashtable.Delete("Asan")
	fmt.Println(hashtable.Search("Asan"))

}

const Size = 7

type Hashtable struct {
	array [Size]*bucket
}
type bucket struct {
	head *bucketnode
}
type bucketnode struct {
	key  string
	next *bucketnode
}

//HASHTABLE
//func (h *Hashtable) REIndex(key string) int {
//	for _, v := range h.array {
//		if key == h.array[v] {
//			return v
//		}
//	}
//	return 0
//}

func (h *Hashtable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}
func (h *Hashtable) Search(key string) bool {
	index := hash(key)
	if h.array[index].search(key) == true {
		fmt.Println(index)
	}

	return h.array[index].search(key)

}

func (h *Hashtable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}
func hash(key string) int {
	sum := 0
	for _, i := range key {
		sum = sum + int(i)
	}
	return sum % Size
}

//bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketnode{key: k}
		newNode.next = b.head
		b.head = newNode
	}
}
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

//func (b *bucket) delete(k string) {
//
//	if b.head.key == k {
//		b.head = b.head.next
//		return
//	}
//	previousNode := b.head
//	for previousNode.next != nil {
//		if previousNode.key == k {
//			previousNode.next = previousNode.next.next
//			break
//		}
//		previousNode = previousNode.next
//
//	}
//
//}
func (b *bucket) delete(k string) bool {

	tempNode := b.head
	previousNode := b.head

	if tempNode != nil && tempNode.key == k {
		b.head = tempNode.next
		return true
	}

	for tempNode != nil && tempNode.key != k {
		previousNode = tempNode
		tempNode = tempNode.next
	}

	if tempNode == nil {
		return false
	}
	previousNode.next = tempNode.next

	return false
}

func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
