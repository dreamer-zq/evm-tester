package tester

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue_Iterate(t *testing.T){
	q := NewQueue[int]()
	for i := 0; i < 10; i++ {
		q.Add(i)
	}
	q.Iterate(func(v int) bool {
		if v%2 == 0 {
			return true
		}
		return false
	})
	require.Equal(t, 5, q.Length())
}