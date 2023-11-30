package tester

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
	"golang.org/x/exp/slog"
)

// Pool is a pool of goroutines that can be used to execute tasks.
type Pool struct {
	service string
	p  *ants.Pool
	t  *time.Ticker
	wg sync.WaitGroup
}

// NewPool creates a new pool with the specified size.
//
// Parameters:
// - size: an integer representing the size of the pool.
//
// Returns:
// - a pointer to a Pool object.
func NewPool(size int,service string) *Pool {
	p, err := ants.NewPool(size, ants.WithLogger(newLogger()))
	if err != nil {
		panic(err)
	}
	pool := &Pool{
		service: service,
		p: p,
		t: time.NewTicker(5 * time.Second),
	}
	go pool.start()
	return pool
}

// Submit submits a task to the pool.
//
// The task parameter is a function that will be executed by the pool.
// It does not take any parameters and does not return any values.
func (p *Pool) Submit(task func()) {
	p.wg.Add(1)
	p.p.Submit(func() {
		defer p.wg.Done()
		task()
	})
}

// Close stops the goroutines in the Pool and waits for them to finish.
//
// No parameters.
// No return types.
func (p *Pool) Close() {
	p.wg.Wait()
	p.p.Release()
	p.t.Stop()
}

// Finish waits until all goroutines have finished.
//
// No parameters.
// No return types.
func (p *Pool) Finish() {
	p.wg.Wait()
}

func (p *Pool) start() {
	for range p.t.C {
		slog.Info("goroutine counter","service",p.service, "running", p.p.Running(), "waiting", p.p.Waiting())
	}
}

type logger struct {
	l *slog.Logger
}

func newLogger() logger {
	return logger{slog.Default()}
}

func (l logger) Printf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.l.Debug(msg)
}
