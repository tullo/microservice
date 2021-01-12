package stats

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/atomic"
)

// Flusher is a context-driven background data flush job.
type Flusher struct {
	context.Context
	db         *sqlx.DB
	enabled    *atomic.Bool
	finish     func()
	queueIndex *atomic.Uint32 // queueIndex is a key for []queues.
	queueMask  uint32         // queueMask is a masking value for queueIndex -> key.
	queues     []*Queue       // queues hold a set of writable queues.
}

// NewFlusher creates a *Flusher.
func NewFlusher(ctx context.Context, db *sqlx.DB) (*Flusher, error) {
	queueSize := 1 << 4 // 2⁴ = 16
	ctx, cancel := context.WithCancel(ctx)
	f := &Flusher{
		Context:    ctx,
		finish:     cancel,
		db:         db,
		enabled:    atomic.NewBool(true),
		queueIndex: atomic.NewUint32(0),
		queueMask:  uint32(queueSize - 1),
		queues:     NewQueues(queueSize),
	}

	go f.run()

	return f, nil
}

// Push spreads queue writes evenly across all queues.
func (f *Flusher) Push(item *Incoming) error {
	if f.enabled.Load() {
		index := f.queueIndex.Inc() & f.queueMask

		return f.queues[index].Enqueue(item)
	}

	return errFlusherDisabled
}

func (f *Flusher) run() {
	log.Println("Started background job")

	defer f.finish()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			f.flush()

			continue
		case <-f.Context.Done():
			log.Println("Got cancel")
			f.enabled.Store(false)
			f.flush()
		}

		break
	}

	log.Println("Exiting Run")
}

func (f *Flusher) flush() {
	var err error

	fields := strings.Join(IncomingFields, ",")
	named := ":" + strings.Join(IncomingFields, ",:")
	query := fmt.Sprintf("insert into %s (%s) values (%s)", IncomingTable, fields, named)

	var batchInsertSize int
	for _, queue := range f.queues {
		rows := queue.Dequeue()
		for len(rows) > 0 {
			batchInsertSize = 1000
			if len(rows) < batchInsertSize {
				batchInsertSize = len(rows)
			}
			if _, err = f.db.NamedExec(query, rows[:batchInsertSize]); err != nil {
				log.Println("Error when flushing data:", err)
			}
			rows = rows[batchInsertSize:]
		}
	}
}

/*
func (f *Flusher) flush() {
	log.Println("Background flush")
}
*/
