package models

import (
	"encoding/json"
	"math"
	"os"
	"strings"
	"testing"
)

func TestSplitsUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/splits_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp SplitsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal splits response: %v", err)
	}

	if resp.Symbol != "IBM" {
		t.Fatalf("expected symbol IBM, got %s", resp.Symbol)
	}

	if len(resp.Data) != 2 {
		t.Fatalf("expected 2 split records, got %d", len(resp.Data))
	}

	first := resp.Data[0]
	if first.EffectiveDate != "2021-11-04" {
		t.Fatalf("expected first effective date 2021-11-04, got %s", first.EffectiveDate)
	}
	if math.Abs(first.SplitFactor-1.0460) > 1e-9 {
		t.Fatalf("expected first split factor 1.0460, got %f", first.SplitFactor)
	}

	second := resp.Data[1]
	if second.EffectiveDate != "1999-05-27" {
		t.Fatalf("expected second effective date 1999-05-27, got %s", second.EffectiveDate)
	}
	if math.Abs(second.SplitFactor-2.0) > 1e-9 {
		t.Fatalf("expected second split factor 2.0, got %f", second.SplitFactor)
	}
}

func TestSplitsString(t *testing.T) {
	resp := SplitsResponse{
		Symbol: "IBM",
		Data: []SplitRecord{
			{EffectiveDate: "2021-11-04", SplitFactor: 1.0460},
			{EffectiveDate: "1999-05-27", SplitFactor: 2.0000},
		},
	}

	rendered := resp.String()

	required := []string{
		"Splits for IBM",
		"2021-11-04: 1.0460",
		"1999-05-27: 2.0000",
	}

	for _, snippet := range required {
		if !strings.Contains(rendered, snippet) {
			t.Fatalf("String() missing expected snippet %q\nRendered:\n%s", snippet, rendered)
		}
	}
}
