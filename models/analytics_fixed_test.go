package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestAnalyticsFixedWindowUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/analytics_fixed_window.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp AnalyticsFixedWindowResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal analytics fixed response: %v", err)
	}

	if resp.MetaData.Symbols != "IBM,AAPL,MSFT" {
		t.Fatalf("symbols mismatch: %s", resp.MetaData.Symbols)
	}

	if resp.MetaData.MinDT != "2023-07-03" || resp.MetaData.MaxDT != "2023-08-31" {
		t.Fatalf("unexpected range: %s-%s", resp.MetaData.MinDT, resp.MetaData.MaxDT)
	}

	group, ok := resp.Payload["RETURNS_CALCULATIONS"]
	if !ok {
		t.Fatalf("expected RETURNS_CALCULATIONS group")
	}

	if _, ok := group["MEAN"]; !ok {
		t.Fatalf("expected MEAN calc present")
	}
}

func TestAnalyticsFixedWindowString(t *testing.T) {
	resp := AnalyticsFixedWindowResponse{
		MetaData: AnalyticsFixedMetaData{
			Symbols:  "IBM,AAPL,MSFT",
			MinDT:    "2023-07-03",
			MaxDT:    "2023-08-31",
			Ohlc:     "Close",
			Interval: "DAILY",
		},
		Payload: map[string]map[string]json.RawMessage{
			"RETURNS_CALCULATIONS": {
				"MEAN":   []byte("{}"),
				"STDDEV": []byte("{}"),
			},
		},
	}

	out := resp.String()
	for _, s := range []string{"IBM,AAPL,MSFT", "RETURNS_CALCULATIONS", "MEAN"} {
		if !strings.Contains(out, s) {
			t.Fatalf("String() missing %q\nOutput:\n%s", s, out)
		}
	}
}
