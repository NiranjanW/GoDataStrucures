package main

import "fmt"

type MaxHeap struct {
	array []int
}

//Inserts  adds
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)

}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}

}

//Extracts the largest key and removes from Heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l = len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array is of len 0")
		return -1
	}

	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:]
	h.maxHeapifyDown(0)

	return extracted

}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	//loop while index has at least one child
	for l <= lastIndex {
		//when left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] { //when left child is larger
			childToCompare = l
		} else {
			childToCompare = r //when right child is larger
		}

	}
	//compare array value and swap
	if h.array[parent(index)] < h.array[childToCompare] {
		h.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
	} else {
		return
	}

}

//index -1/2
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 5}
	fmt.Println(buildHeap)
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

}
