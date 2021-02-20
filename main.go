package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// Get Len of IntHeap
func (iheap IntHeap) Len() int {return len(iheap)}

// Check if element of i index is less than j index
func (iheap IntHeap) Less(i, j int) bool {return iheap[i] < iheap[j]}

// Swaps the element of i to j index
func (iheap IntHeap) Swap(i, j int)  {iheap[i], iheap[j] = iheap[j], iheap[i]}

// Pushes the item
func (iheap *IntHeap) Push(heapintf interface{})  {*iheap = append(*iheap, heapintf.(int))}

// Pops the item from the heap
func (iheap *IntHeap) Pop() interface{} {
	var n, x1 int
	var previous = *iheap

	n = len(previous)
	x1 = previous[n - 1]
	*iheap = previous[0 : n - 1]

	return x1
}

func main() {
	var intHeap = &IntHeap{1,2,3,4,5,7,8,9,10,11,12}

	heap.Init(intHeap)
	heap.Push(intHeap, 6)

	fmt.Printf("Minimum Heap : %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d\n", heap.Pop(intHeap))
	}
}