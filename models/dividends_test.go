package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestDividendsUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/dividends_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp DividendsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal dividends: %v", err)
	}

	if resp.Symbol != "IBM" {
		t.Fatalf("expected symbol IBM, got %s", resp.Symbol)
	}

	if len(resp.Data) != 6 {
		t.Fatalf("expected 6 dividend records, got %d", len(resp.Data))
	}

	first := resp.Data[0]
	if first.ExDividendDate != "2025-11-10" || first.Amount != 1.68 {
		t.Fatalf("unexpected first record: %+v", first)
	}
}

func TestDividendsString(t *testing.T) {
	resp := DividendsResponse{
		Symbol: "IBM",
		Data: []DividendRecord{
			{ExDividendDate: "2025-11-10", PaymentDate: "2025-12-10", Amount: 1.68},
			{ExDividendDate: "2025-08-08", PaymentDate: "2025-09-10", Amount: 1.68},
		},
	}

	out := resp.String()
	want := []string{"Dividends for IBM", "2025-11-10", "1.68", "2025-12-10"}
	for _, w := range want {
		if !strings.Contains(out, w) {
			t.Fatalf("String() missing %q\nOutput:\n%s", w, out)
		}
	}
}
