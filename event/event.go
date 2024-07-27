package event

import (
	"time"
)

type EntryMeta struct {
	Key   []string `json:"key"`
	Value []string `json:"value"`
}

type Event struct {
	Hostname               string            `json:"hostname"`
	SiteID                 string            `json:"site_id"`
	SiteName               string            `json:"site_name"`
	EventName              string            `json:"event_name"`
	StartTime              time.Time         `json:"start_time"`
	Pathname               string            `json:"pathname"`
	NavigationFrom         map[string]string `json:"navigation_from"`
	EntryMeta              EntryMeta         `json:"entry_meta"`
	UtmMedium              *string           `json:"utm_medium"`
	UtmSource              *string           `json:"utm_source"`
	UtmCampaign            *string           `json:"utm_campaign"`
	UtmContent             *string           `json:"utm_content"`
	UtmTerm                *string           `json:"utm_term"`
	Referrer               string            `json:"referrer"`
	ReferrerSource         string            `json:"referrer_source"`
	ScreenSize             string            `json:"screen_size"`
	Device                 string            `json:"device"`
	OperatingSystem        string            `json:"operating_system"`
	OperatingSystemVersion string            `json:"operating_system_version"`
	Browser                string            `json:"browser"`
	BrowserVersion         string            `json:"browser_version"`
}

type RequestEvent struct {
	Event Event `json:"event"`
}
