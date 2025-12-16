package types

import "testing"

func TestUnmarshalLenient_ReplacesNAForNumericFields(t *testing.T) {
	type payload struct {
		A float64 `json:"a,string"`
		B float64 `json:"b"`
		C int64   `json:"c,string"`
		D uint64  `json:"d,string"`
	}

	var got payload
	data := []byte(`{"a":"n/a","b":"n/a","c":"N/A","d":"na"}`)
	if err := UnmarshalLenient(data, &got); err != nil {
		t.Fatalf("UnmarshalLenient returned error: %v", err)
	}

	if got.A != 0 || got.B != 0 || got.C != 0 || got.D != 0 {
		t.Fatalf("expected all numeric fields to be zeroed, got %+v", got)
	}
}

func TestUnmarshalLenient_ETFProfile_AllowsNAInNumericFields(t *testing.T) {
	data := []byte(`{
  "net_assets": "1000",
  "net_expense_ratio": "n/a",
  "portfolio_turnover": "n/a",
  "dividend_yield": "N/A",
  "inception_date": "2020-01-01",
  "leveraged": "NO",
  "sectors": [{"sector": "TECH", "weight": "n/a"}],
  "holdings": [{"symbol": "ABC", "description": "n/a", "weight": "n/a"}]
}`)

	var profile ETFProfile
	if err := UnmarshalLenient(data, &profile); err != nil {
		t.Fatalf("UnmarshalLenient returned error: %v", err)
	}

	if profile.NetAssets != 1000 {
		t.Fatalf("expected net assets 1000, got %d", profile.NetAssets)
	}
	if profile.NetExpenseRatio != 0 {
		t.Fatalf("expected net expense ratio 0, got %v", profile.NetExpenseRatio)
	}
	if profile.DividendYield != 0 {
		t.Fatalf("expected dividend yield 0, got %v", profile.DividendYield)
	}
	if profile.PortfolioTurnover != "n/a" {
		t.Fatalf("expected portfolio turnover to remain \"n/a\", got %q", profile.PortfolioTurnover)
	}
	if len(profile.Sectors) != 1 || profile.Sectors[0].Weight != 0 {
		t.Fatalf("expected sector weight to be 0, got %+v", profile.Sectors)
	}
	if len(profile.Holdings) != 1 || profile.Holdings[0].Weight != 0 {
		t.Fatalf("expected holding weight to be 0, got %+v", profile.Holdings)
	}
	if profile.Holdings[0].Description != "n/a" {
		t.Fatalf("expected holding description to remain \"n/a\", got %q", profile.Holdings[0].Description)
	}
}
