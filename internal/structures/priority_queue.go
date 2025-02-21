package structures

import "container/heap" // да, это уже не pure go, но библиотека стандартная

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Frequency < pq[j].Frequency }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) Enqueue(node *Node) {
	heap.Push(pq, node)
}

func (pq *PriorityQueue) Dequeue() *Node {
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(pq).(*Node)
}
