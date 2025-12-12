package models

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// AnalyticsFixedWindowParams defines parameters for the fixed window analytics endpoint.
type AnalyticsFixedWindowParams struct {
	Symbols      string   // comma-separated symbols
	Range        []string // one or two RANGE parameters
	Interval     string
	Calculations string // comma-separated list
	Ohlc         string // optional
	DataType     string // optional
}

// AnalyticsFixedMetaData captures meta_data for fixed window analytics.
type AnalyticsFixedMetaData struct {
	Symbols  string `json:"symbols"`
	MinDT    string `json:"min_dt"`
	MaxDT    string `json:"max_dt"`
	Ohlc     string `json:"ohlc"`
	Interval string `json:"interval"`
}

// AnalyticsFixedWindowResponse holds the response; payload kept generic for flexibility.
type AnalyticsFixedWindowResponse struct {
	MetaData AnalyticsFixedMetaData                `json:"meta_data"`
	Payload  map[string]map[string]json.RawMessage `json:"payload"`
}

// String renders a concise summary.
func (r AnalyticsFixedWindowResponse) String() string {
	var sb strings.Builder
	sb.WriteString("Analytics Fixed Window\n")
	sb.WriteString(fmt.Sprintf("Symbols: %s | Interval: %s\n", r.MetaData.Symbols, r.MetaData.Interval))
	sb.WriteString(fmt.Sprintf("Range: %s → %s | OHLC: %s\n", r.MetaData.MinDT, r.MetaData.MaxDT, r.MetaData.Ohlc))

	if len(r.Payload) > 0 {
		groupKeys := make([]string, 0, len(r.Payload))
		for k := range r.Payload {
			groupKeys = append(groupKeys, k)
		}
		sort.Strings(groupKeys)
		sb.WriteString(fmt.Sprintf("Groups: %s\n", strings.Join(groupKeys, ", ")))

		first := r.Payload[groupKeys[0]]
		calcs := make([]string, 0, len(first))
		for k := range first {
			calcs = append(calcs, k)
		}
		sort.Strings(calcs)
		if len(calcs) > 0 {
			sb.WriteString(fmt.Sprintf("Calculations in %s: %s\n", groupKeys[0], strings.Join(calcs, ", ")))
		}
	}

	return sb.String()
}
