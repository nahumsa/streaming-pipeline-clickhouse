package repositories

import (
	"context"

	"github.com/nahumsa/streaming-pipeline-clickhouse/event"
)

type EventRepository interface {
	InsertEvent(ctx context.Context, event event.Event) error
}
