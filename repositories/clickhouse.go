package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/nahumsa/streaming-pipeline-clickhouse/config"
	"github.com/nahumsa/streaming-pipeline-clickhouse/event"
)

type clickhouseRepository struct {
	conn clickhouse.Conn
}

func NewClickhouseRepository(conn clickhouse.Conn) EventRepository {
	return &clickhouseRepository{
		conn: conn,
	}
}

func (r *clickhouseRepository) InsertEvent(ctx context.Context, event event.Event) error {
	batch, err := r.conn.PrepareBatch(ctx, "INSERT INTO events")
	if err != nil {
		return err
	}

	err = batch.Append(
		event.Hostname,
		event.SiteID,
		event.SiteName,
		event.EventName,
		event.StartTime,
		event.Pathname,
		fmt.Sprint(event.NavigationFrom),
		event.EntryMeta.Key,
		event.EntryMeta.Value,
		event.UtmMedium,
		event.UtmSource,
		event.UtmCampaign,
		event.UtmContent,
		event.UtmTerm,
		event.Referrer,
		event.ReferrerSource,
		event.ScreenSize,
		event.Device,
		event.OperatingSystem,
		event.OperatingSystemVersion,
		event.Browser,
		event.BrowserVersion,
	)
	if err != nil {
		return err
	}

	return batch.Send()
}

func SetupClickhouse(environment config.Environment) (clickhouse.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{environment.ClickhouseHost},
		Auth: clickhouse.Auth{
			Database: environment.ClickhouseDB,
			Username: environment.ClickhouseUsername,
			Password: environment.ClickhousePass,
		},
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Fatal(err)
		}
	}

	return conn, nil
}
