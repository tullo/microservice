package stats

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/atomic"
)

// Flusher is a context-driven background data flush job.
type Flusher struct {
	context.Context
	finish func()

	enabled *atomic.Bool

	db *sqlx.DB
}

// NewFlusher creates a *Flusher.
func NewFlusher(ctx context.Context, db *sqlx.DB) (*Flusher, error) {
	f := &Flusher{
		db:      db,
		enabled: atomic.NewBool(true),
	}
	f.Context, f.finish = context.WithCancel(context.Background())
	go f.run(ctx)
	return f, nil
}

func (f *Flusher) run(ctx context.Context) {
	log.Println("Started background job")

	defer f.finish()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			f.flush()
			continue
		case <-ctx.Done():
			log.Println("Got cancel")
			f.enabled.Store(false)
			f.flush()
		}
		break
	}

	log.Println("Exiting Run")
}

func (f *Flusher) flush() {
	log.Println("Background flush")
}
