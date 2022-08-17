package main

import (
	"fmt"
)

type QueueElement struct {
	value    interface{}
	priority int
}

type PriorityQueue struct {
	maxSize int
	data    []QueueElement
	size    int
}

func (p *PriorityQueue) SetMaxSize(maxSize int) {
	p.maxSize = maxSize
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.size == 0
}

func (p *PriorityQueue) Display() {
	if p.IsEmpty() {
		panic("La cola esta vacia.")
	}
	arr := make([]interface{}, p.size)
	i := 0
	for i < p.size {
		arr[i] = p.data[i].value
		i++
	}
	fmt.Println("Elementos de la cola de prioridad:", arr)
}

func (p *PriorityQueue) Enqueue(el interface{}, priority int) {
	newElement := QueueElement{el, priority}
	if p.size == p.maxSize {
		panic("La cola ha alcanzado su límite de tamaño máximo.")
	}

	p.data = append(p.data, newElement)
	p.size++
}

func (p *PriorityQueue) Dequeue() QueueElement {
	if p.IsEmpty() {
		panic("La cola esta vacia.")
	}

	idx := 0
	for i := 0; i < p.size; i++ {
		if p.data[i].priority < p.data[idx].priority {
			idx = i
		}
	}
	dequeuedEl := p.data[idx]
	p.data = append(p.data[:idx], p.data[idx+1:]...)
	p.size--
	return dequeuedEl
}

func (p *PriorityQueue) Peek() QueueElement {
	if p.IsEmpty() {
		panic("La cola esta vacia.")
	}
	dequeuedEl := p.data[0]
	for _, el := range p.data {
		if el.priority < dequeuedEl.priority {
			dequeuedEl = el
		}
	}
	return dequeuedEl
}

func main() {
	pq := PriorityQueue{}
	pq.SetMaxSize(10)
	pq.Enqueue(4, 5)
	pq.Enqueue(3, 2)
	pq.Enqueue(8, 3)
	pq.Enqueue(6, 1)
	pq.Enqueue(7, 4)

	pq.Display()
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	fmt.Println(pq.Dequeue())
	pq.Display()
}
