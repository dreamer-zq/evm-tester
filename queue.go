package tester

import (
	"sync"

	"container/list"
)

// Queue is a wrapper of github.com/eapache/queue/v2.Queue
type Queue[T any] struct {
	q *list.List
	p *Pool
	l sync.Mutex
}

// NewQueue creates a new Queue instance.
//
// This function does not take any parameters.
// It returns a pointer to a Queue[T] object.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: list.New(),
		p: NewPool(500, "validator"),
	}
}

// Add adds the given value to the queue.
//
// It takes a single parameter:
// - v: the value to be added to the queue.
func (q *Queue[T]) Add(v T) {
	q.l.Lock()
	defer q.l.Unlock()

	q.q.PushBack(v)
}

// AddBatch adds a batch of elements to the queue.
//
// It takes a slice of elements, v, and adds each element to the queue.
// The function does not return anything.
func (q *Queue[T]) AddBatch(v []T) {
	q.l.Lock()
	defer q.l.Unlock()

	for _, item := range v {
		q.q.PushBack(item)
	}
}

// Iterate iterates over the elements of the queue and applies a function to each element.
//
// f: The function to apply to each element of the queue. It takes an element as a parameter and returns a boolean value.
//
//	If the function returns true, the element will be removed from the queue. If the function returns false,
//	the element will not be removed from the queue.
func (q *Queue[T]) Iterate(f func(T) bool) {
	q.l.Lock()
	defer q.l.Unlock()

	var next *list.Element
	for e := q.q.Front(); e != nil; e = next {
		next = e.Next()
		if f(e.Value.(T)) {
			q.q.Remove(e)
		}
	}
}

// IterateParallel iterates over the elements in the queue in parallel and applies a function to each element.
//
// The function `f` is applied to each element in the queue. If the function returns `true`, the element is removed from the queue.
// The function does not guarantee the order in which the elements are processed.
func (q *Queue[T]) IterateParallel(f func(T) bool) {
	var next *list.Element
	for e := q.q.Front(); e != nil; e = next {
		next = e.Next()
		q.p.Submit(func() {
			if f(e.Value.(T)) {
				q.remove(e)
			}
		})
	}
	q.p.Finish()
}

func (q *Queue[T]) Length() int {
	return q.q.Len()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue[T]) remove(e *list.Element) {
	q.l.Lock()
	defer q.l.Unlock()

	q.q.Remove(e)
}