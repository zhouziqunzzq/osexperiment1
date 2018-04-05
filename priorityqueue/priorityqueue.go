package priorityqueue

import (
	"container/heap"
	. "github.com/zhouziqunzzq/osexperiment1/PCB"
	"fmt"
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PCB

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use less than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PCB)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) PriPush(x interface{}) {
	heap.Push(pq, x)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) PriPop() interface{} {
	return heap.Pop(pq)
}

func (pq PriorityQueue) IsEmpty() bool {
	return len(pq) == 0
}

func (pq PriorityQueue) First() interface{} {
	return pq[0]
}

func (pq PriorityQueue) Last() interface{} {
	return pq[len(pq)-1]
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *PCB) {
	//item.value = value
	//item.priority = priority
	heap.Fix(pq, item.Index)
}

func (pq PriorityQueue) PrintQueue() {
	fmt.Printf("PID\tName\tPrority\n")
	for i := 0; i < len(pq); i++ {
		pq[i].PrintPCB()
	}
}
