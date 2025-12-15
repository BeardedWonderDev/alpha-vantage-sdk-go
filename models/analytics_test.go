package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestAnalyticsSlidingWindowUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/analytics_sliding_window.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp AnalyticsSlidingWindowResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal analytics response: %v", err)
	}

	if resp.MetaData.Symbols != "AAPL,IBM" {
		t.Fatalf("expected symbols AAPL,IBM got %s", resp.MetaData.Symbols)
	}

	if resp.MetaData.WindowSize != 20 {
		t.Fatalf("expected window size 20 got %d", resp.MetaData.WindowSize)
	}

	if len(resp.Payload) == 0 {
		t.Fatalf("expected payload entries")
	}

	retGroup, ok := resp.Payload["RETURNS_CALCULATIONS"]
	if !ok {
		t.Fatalf("expected RETURNS_CALCULATIONS group")
	}

	if _, ok := retGroup["MEAN"]; !ok {
		t.Fatalf("expected MEAN calculation")
	}
}

func TestAnalyticsSlidingWindowString(t *testing.T) {
	resp := AnalyticsSlidingWindowResponse{
		MetaData: AnalyticsMetaData{
			Symbols:    "AAPL,IBM",
			WindowSize: 20,
			MinDT:      "2025-10-13",
			MaxDT:      "2025-12-11",
			Ohlc:       "Close",
			Interval:   "DAILY",
		},
		Payload: map[string]map[string]json.RawMessage{
			"RETURNS_CALCULATIONS": {
				"MEAN":   []byte("{}"),
				"STDDEV": []byte("{}"),
			},
		},
	}

	out := resp.String()
	for _, snippet := range []string{"AAPL,IBM", "Window: 20", "RETURNS_CALCULATIONS", "MEAN"} {
		if !strings.Contains(out, snippet) {
			t.Fatalf("String() missing %q\nOutput:\n%s", snippet, out)
		}
	}
}
