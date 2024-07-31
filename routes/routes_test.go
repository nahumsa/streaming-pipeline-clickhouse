package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nahumsa/streaming-pipeline-clickhouse/event"
	"github.com/stretchr/testify/assert"
	// add Testify package
)

func GetJsonTestRequestResponse(app *fiber.App, method string, url string, reqBody any) (code int, respBody map[string]any, err error) {
	var req *http.Request

	if reqBody == nil {
		req = httptest.NewRequest(method, url, nil)
	} else {
		bodyJson, _ := json.Marshal(reqBody)
		req = httptest.NewRequest(method, url, bytes.NewReader(bodyJson))

	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, 10)
	code = resp.StatusCode
	// If error we're done
	if err != nil {
		return
	}
	// If no body content, we're done
	if resp.ContentLength == 0 {
		return
	}
	bodyData := make([]byte, resp.ContentLength)
	_, _ = resp.Body.Read(bodyData)
	err = json.Unmarshal(bodyData, &respBody)
	return
}

type MockEventRepository struct{}

func (r *MockEventRepository) InsertEvent(ctx context.Context, event event.Event) error {
	return nil
}

func TestRouters(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		payload      event.RequestEvent
	}{
		{
			description:  "post sendEvent valid request",
			route:        "/api/v1/sendEvent",
			expectedCode: 201,
			payload: event.RequestEvent{
				Event: event.Event{
					Hostname:               "https://www.casaevideo.com.br",
					SiteID:                 "",
					SiteName:               "casaevideo",
					EventName:              "pageview",
					StartTime:              time.Date(2023, 7, 27, 2, 52, 8, 0, time.UTC),
					Pathname:               "/",
					NavigationFrom:         map[string]string{},
					EntryMeta:              event.EntryMeta{Key: []string{"Teste AB", "pageId"}, Value: []string{"false", "pages-testeperformance-95d70f581f49"}},
					UtmMedium:              nil,
					UtmSource:              nil,
					UtmCampaign:            nil,
					UtmContent:             nil,
					UtmTerm:                nil,
					Referrer:               "",
					ReferrerSource:         "direct",
					ScreenSize:             "2560x1080",
					Device:                 "desktop",
					OperatingSystem:        "Linux",
					OperatingSystemVersion: "Unknown",
					Browser:                "Chrome",
					BrowserVersion:         "126",
				},
			},
		},
		// TODO: add edge cases
	}

	app := fiber.New()
	SetupRoutes(app, &MockEventRepository{})

	for _, test := range tests {
		code, _, err := GetJsonTestRequestResponse(app, "POST", test.route, test.payload)
		assert.Nil(t, err)

		assert.Equalf(t, test.expectedCode, code, test.description)
	}
}
