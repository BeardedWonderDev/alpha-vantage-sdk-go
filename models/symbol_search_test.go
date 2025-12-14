package models

import (
	"encoding/json"
	"math"
	"os"
	"testing"
)

func TestSymbolSearchUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/symbol_search_sample.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp SymbolSearchResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal symbol search response: %v", err)
	}

	if len(resp.BestMatches) != 10 {
		t.Fatalf("expected 10 matches, got %d", len(resp.BestMatches))
	}

	first := resp.BestMatches[0]
	if first.Symbol != "BA" {
		t.Fatalf("expected first symbol BA, got %s", first.Symbol)
	}
	if first.Name != "Boeing Company" {
		t.Fatalf("expected first name Boeing Company, got %s", first.Name)
	}
	if first.Type != "Equity" {
		t.Fatalf("expected first type Equity, got %s", first.Type)
	}
	if first.Region != "United States" {
		t.Fatalf("expected first region United States, got %s", first.Region)
	}
	if first.MarketOpen != "09:30" || first.MarketClose != "16:00" {
		t.Fatalf("unexpected market hours: open %s close %s", first.MarketOpen, first.MarketClose)
	}
	if first.TimeZone != "UTC-04" {
		t.Fatalf("expected first timezone UTC-04, got %s", first.TimeZone)
	}
	if first.Currency != "USD" {
		t.Fatalf("expected first currency USD, got %s", first.Currency)
	}
	if math.Abs(first.MatchScore-1.0) > 1e-9 {
		t.Fatalf("expected first matchScore 1.0, got %f", first.MatchScore)
	}

	second := resp.BestMatches[1]
	if second.Symbol != "BA.LON" {
		t.Fatalf("expected second symbol BA.LON, got %s", second.Symbol)
	}
	if second.Currency != "GBX" {
		t.Fatalf("expected second currency GBX, got %s", second.Currency)
	}
	if math.Abs(second.MatchScore-0.6667) > 1e-6 {
		t.Fatalf("expected second matchScore 0.6667, got %f", second.MatchScore)
	}

	seventh := resp.BestMatches[6]
	if seventh.Symbol != "BAAPL" {
		t.Fatalf("expected seventh symbol BAAPL, got %s", seventh.Symbol)
	}
	if seventh.Name != "null" {
		t.Fatalf("expected seventh name \"null\", got %q", seventh.Name)
	}
}
