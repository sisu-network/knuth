package atomic

import "sync"

type Queue interface {
	Enqueue(el any)
	Dequeue() any
	Peek() any
	Len()
}

type queue struct {
	lock     *sync.RWMutex
	elements []any
}

func NewQueue() *queue {
	return &queue{
		lock:     &sync.RWMutex{},
		elements: make([]any, 0),
	}
}

func (q *queue) Enqueue(el any) {
	q.lock.Lock()
	q.elements = append(q.elements, el)
	q.lock.Unlock()
}

func (q *queue) Peek() any {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if len(q.elements) == 0 {
		return nil
	}

	return q.elements[0]
}

func (q *queue) Dequeue() any {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.elements) == 0 {
		return nil
	}

	el := q.elements[0]
	q.elements = q.elements[1:]

	return el
}

func (q *queue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return len(q.elements)
}
