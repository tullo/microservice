package stats

import "sync"

// Queue provides a queuing structure for Incoming.
type Queue struct {
	sync.RWMutex
	values []*Incoming
}

// NewQueue creates a new *Queue instance.
func NewQueue() *Queue {
	return &Queue{
		RWMutex: sync.RWMutex{},
		values:  make([]*Incoming, 0),
	}
}

// NewQueues creates a slice of *Queue instances.
func NewQueues(size int) []*Queue {
	result := make([]*Queue, size)
	for i := 0; i < size; i++ {
		result[i] = NewQueue()
	}

	return result
}

// Enqueue adds a new item to the queue.
func (q *Queue) Enqueue(item *Incoming) error {
	q.Lock()
	defer q.Unlock()
	q.values = append(q.values, item)

	return nil
}

// Dequeue returns current queue items and clears it.
func (q *Queue) Dequeue() (result []*Incoming) {
	q.Lock()
	defer q.Unlock()
	result, q.values = q.values[:len(q.values)], q.values[len(q.values):]

	return
}

// Length returns the current queue size.
func (q *Queue) Length() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.values)
}
