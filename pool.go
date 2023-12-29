package tester

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/panjf2000/ants/v2"
)

// Stat is the statistics of the Pool.
type Stat struct {
	Running int
	Waiting int
	Free    int
	Cap     int
	Name    string
}

// Pool is a pool of goroutines that can be used to execute tasks.
type Pool struct {
	service string
	p       *ants.Pool
	wg      sync.WaitGroup
}

// NewPool creates a new pool with the specified size.
//
// Parameters:
// - size: an integer representing the size of the pool.
//
// Returns:
// - a pointer to a Pool object.
func NewPool(size int, service string) *Pool {
	p, err := ants.NewPool(size, ants.WithLogger(newLogger()))
	if err != nil {
		panic(err)
	}
	pool := &Pool{
		service: service,
		p:       p,
	}
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
}

// Finish waits until all goroutines have finished.
//
// No parameters.
// No return types.
func (p *Pool) Finish() {
	p.wg.Wait()
}

// Stat returns the statistics of the Pool.
//
// It returns a stat struct containing the number of running, waiting, free, and maximum capacity of the Pool,
// as well as the name of the service associated with the Pool.
func (p *Pool) Stat() Stat {
	return Stat{
		Running: p.p.Running(),
		Waiting: p.p.Waiting(),
		Free:    p.p.Free(),
		Cap:     p.p.Cap(),
		Name:    p.service,
	}
}

type logger struct {
	l *slog.Logger
}

func newLogger() logger {
	return logger{slog.Default()}
}

func (l logger) Printf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.l.Debug(msg)
}
