package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// AnalyticsSlidingWindowParams defines parameters for the sliding window analytics endpoint.
type AnalyticsSlidingWindowParams struct {
	Symbols      string   // comma-separated symbols
	Range        []string // one or two RANGE parameters; order preserved
	Interval     string   // e.g., 1min, DAILY, WEEKLY
	WindowSize   int
	Calculations string // comma separated list
	Ohlc         string // optional
}

// AnalyticsMetaData represents the metadata section of the response.
type AnalyticsMetaData struct {
	Symbols    string `json:"symbols"`
	WindowSize int    `json:"window_size"`
	MinDT      string `json:"min_dt"`
	MaxDT      string `json:"max_dt"`
	Ohlc       string `json:"ohlc"`
	Interval   string `json:"interval"`
}

// AnalyticsSlidingWindowResponse represents the full response.
// Payload is left as nested raw messages so callers can unmarshal specific metrics as needed.
type AnalyticsSlidingWindowResponse struct {
	MetaData AnalyticsMetaData                     `json:"meta_data"`
	Payload  map[string]map[string]json.RawMessage `json:"payload"`
}

// String prints a concise summary of the response contents.
func (r AnalyticsSlidingWindowResponse) String() string {
	var sb strings.Builder
	sb.WriteString("Analytics Sliding Window\n")
	sb.WriteString(fmt.Sprintf("Symbols: %s | Interval: %s | Window: %d\n", r.MetaData.Symbols, r.MetaData.Interval, r.MetaData.WindowSize))
	sb.WriteString(fmt.Sprintf("Range: %s → %s | OHLC: %s\n", r.MetaData.MinDT, r.MetaData.MaxDT, r.MetaData.Ohlc))

	if len(r.Payload) > 0 {
		topKeys := make([]string, 0, len(r.Payload))
		for k := range r.Payload {
			topKeys = append(topKeys, k)
		}
		sort.Strings(topKeys)
		sb.WriteString(fmt.Sprintf("Groups: %s\n", strings.Join(topKeys, ", ")))

		// Also list calculation names within the first group for quick visibility
		firstGroup := r.Payload[topKeys[0]]
		calcs := make([]string, 0, len(firstGroup))
		for k := range firstGroup {
			calcs = append(calcs, k)
		}
		sort.Strings(calcs)
		if len(calcs) > 0 {
			sb.WriteString(fmt.Sprintf("Calculations in %s: %s\n", topKeys[0], strings.Join(calcs, ", ")))
		}
	}

	return sb.String()
}

