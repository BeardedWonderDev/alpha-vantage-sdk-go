package models

import (
	"fmt"
	"strings"
)

// SplitsParams defines parameters for the splits endpoint.
type SplitsParams struct {
	Symbol   string
	DataType string
}

// SplitsResponse models the splits API response.
type SplitsResponse struct {
	Symbol string        `json:"symbol"`
	Data   []SplitRecord `json:"data"`
}

// SplitRecord represents an individual split event.
type SplitRecord struct {
	EffectiveDate string  `json:"effective_date"`
	SplitFactor   float64 `json:"split_factor,string"`
}

// String renders a concise summary of the split history.
func (r SplitsResponse) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Splits for %s\n", r.Symbol))

	limit := len(r.Data)
	if limit > 10 {
		limit = 10
	}

	for i := 0; i < limit; i++ {
		s := r.Data[i]
		sb.WriteString(fmt.Sprintf("%s: %.4f\n", s.EffectiveDate, s.SplitFactor))
	}

	if len(r.Data) > limit {
		sb.WriteString(fmt.Sprintf("...and %d more\n", len(r.Data)-limit))
	}

	return sb.String()
}
