package main

import (
	"container/heap"
	"fmt"
)

// integerHeap of a type
type integerHeap []int

// Len integerHeap method - gets the length of the integerHeap
func (h integerHeap) Len() int {
	return len(h)
}

// Less integerHeap method - checks if element of i index is less j index
func (h integerHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap IntegerHeap method -swaps the element of i to j index
func (h integerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push integerHeap - pushes the item
func (h *integerHeap) Push(heapintf interface{}) {
	*h = append(*h, heapintf.(int))
}

// Pop IntegerHeap method -pops the item from the heap”
func (h *integerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous integerHeap = *h

	n = len(previous)
	x1 = previous[n-1]
	*h = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *integerHeap = &integerHeap{}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Println("minimum: %d\n", (*intHeap)[0])

	for intHeap.Len() > 0 {
		fmt.Println(heap.Pop(intHeap))
	}
}
