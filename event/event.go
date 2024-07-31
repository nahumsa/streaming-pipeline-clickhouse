package event

import (
	"time"
)

type EntryMeta struct {
	Key   []string `json:"key"`
	Value []string `json:"value"`
}

type Event struct {
	Hostname               string            `validate:"required" json:"hostname"`
	SiteName               string            `validate:"required" json:"site_name"`
	EventName              string            `validate:"required" json:"event_name"`
	StartTime              time.Time         `validate:"required" json:"start_time"`
	Pathname               string            `validate:"required" json:"pathname"`
	NavigationFrom         map[string]string `validate:"required" json:"navigation_from"`
	EntryMeta              EntryMeta         `validate:"required" json:"entry_meta"`
	ReferrerSource         string            `validate:"required" json:"referrer_source"`
	ScreenSize             string            `validate:"required" json:"screen_size"`
	Device                 string            `validate:"required" json:"device"`
	OperatingSystem        string            `validate:"required" json:"operating_system"`
	OperatingSystemVersion string            `validate:"required" json:"operating_system_version"`
	Browser                string            `validate:"required" json:"browser"`
	BrowserVersion         string            `validate:"required" json:"browser_version"`
	Referrer               string            `json:"referrer"`
	SiteID                 string            `json:"site_id"`
	UtmTerm                *string           `json:"utm_term"`
	UtmMedium              *string           `json:"utm_medium"`
	UtmSource              *string           `json:"utm_source"`
	UtmCampaign            *string           `json:"utm_campaign"`
	UtmContent             *string           `json:"utm_content"`
}

type RequestEvent struct {
	Event Event `validate:"required" json:"event"`
}
