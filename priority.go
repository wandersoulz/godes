package godes

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type PriorityInterface interface {
	Equals(PriorityInterface) bool
}

type PriorityItem struct {
	Entity    PriorityInterface // The value of the item; arbitrary.
	Priority float64    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (p *PriorityItem) Equals(pi PriorityInterface) bool {
	return false
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PriorityItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(entity interface{}) {
	n := len(*pq)
	item := entity.(*PriorityItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	fmt.Println(n)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Find(item PriorityInterface) *PriorityItem {
	for i := 0; i < len(*pq); i++ {
		curr := (*pq)[i]
		if curr != nil {
			if item.Equals(curr.Entity) {
				return curr
			}
		}
	}
	return nil
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *PriorityItem, priority float64) {
	item.Priority = priority
	heap.Fix(pq, item.index)
}