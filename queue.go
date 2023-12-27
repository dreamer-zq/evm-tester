package tester

import (
	"sync"

	v2 "github.com/eapache/queue/v2"
)

// Queue is a wrapper of github.com/eapache/queue/v2.Queue
type Queue[T any] struct {
	q *v2.Queue[T]
	p *Pool
	l sync.Mutex
}

// NewQueue creates a new Queue instance.
//
// This function does not take any parameters.
// It returns a pointer to a Queue[T] object.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: v2.New[T](),
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

	q.q.Add(v)
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

	length := q.Length()
	for i := 0; i < length; i++ {
		element := q.q.Peek()
		if f(element) {
			_ = q.q.Remove()
		}
	}
}

// IterateParallel iterates over the elements in the queue in parallel and applies a function to each element.
//
// The function `f` is applied to each element in the queue. If the function returns `true`, the element is removed from the queue.
// The function does not guarantee the order in which the elements are processed.
func (q *Queue[T]) IterateParallel(f func(T) bool) {
	q.l.Lock()
	defer q.l.Unlock()

	length := q.Length()
	for i := 0; i < length; i++ {
		element := q.q.Peek()
		q.p.Submit(func() {
			if f(element) {
				_ = q.q.Remove()
			}
		})
	}
	q.p.Finish()
}

func (q *Queue[T]) Length() int {
	return q.q.Length()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.q.Length() == 0
}
