package stats_test

import (
	"testing"

	"github.com/tullo/microservice/server/stats"
)

func TestQueue(t *testing.T) {
	assert := func(ok bool, format string, params ...interface{}) {
		if !ok {
			t.Fatalf(format, params...)
		}
	}

	q := stats.NewQueue()

	assert(q.Length() == 0, "Unexpected queue length: %d != 0", q.Length())
	assert(nil == q.Enqueue(new(stats.Incoming)), "Expected no error on q.Enqueue")
	assert(nil == q.Enqueue(new(stats.Incoming)), "Expected no error on q.Enqueue")
	assert(nil == q.Enqueue(new(stats.Incoming)), "Expected no error on q.Enqueue")

	assert(q.Length() == 3, "Unexpected queue length: %d != 3", q.Length())

	items := q.Dequeue()
	assert(len(items) == 3, "Unexpected items length: %d != 3", len(items))
	assert(q.Length() == 0, "Unexpected queue length: %d != 0", q.Length())

	queues := stats.NewQueues(16)
	assert(len(queues) == 16, "Unexpected queue count: %d != 16", len(queues))

	for i := range queues {
		assert(queues[i] != nil, "Unexpected queue value: expected not nil, index %d", queues[i])
		assert(queues[i].Length() == 0, "Unexpected queue length: %d != 0", queues[i].Length())
	}
}
