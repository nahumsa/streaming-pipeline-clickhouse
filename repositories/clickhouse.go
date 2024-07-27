package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
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

func SetupClickhouse(conString string, createTable bool) (clickhouse.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{conString},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
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

	if createTable {
		err = conn.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS events (
            hostname               String,
            site_id                String,
            site_name              String,
            event_name             String,
            start_time             DateTime,
            pathname               String,
            navigation_from        String,
            entry_meta_key         Array(String),
            entry_meta_value       Array(String),
            utm_medium             Nullable(String),
            utm_source             Nullable(String),
            utm_campaign           Nullable(String),
            utm_content            Nullable(String),
            utm_term               Nullable(String),
            referrer               String,
            referrer_source        String,
            screen_size            String,
            device                 String,
            operating_system       String,
            operating_system_version String,
            browser                String,
            browser_version        String
        ) ENGINE = MergeTree() ORDER BY start_time
    `)
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}
