package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestETFProfileUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/etf_profile_QQQ.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var profile ETFProfile
	if err := json.Unmarshal(data, &profile); err != nil {
		t.Fatalf("failed to unmarshal ETF profile: %v", err)
	}

	if profile.NetAssets != 403000000000 {
		t.Fatalf("expected net assets 403000000000, got %d", profile.NetAssets)
	}

	if profile.NetExpenseRatio != 0.002 {
		t.Fatalf("expected expense ratio 0.002, got %f", profile.NetExpenseRatio)
	}

	if len(profile.Sectors) != 10 {
		t.Fatalf("expected 10 sectors, got %d", len(profile.Sectors))
	}

	if profile.Sectors[0].Sector != "INFORMATION TECHNOLOGY" || profile.Sectors[0].Weight != 0.534 {
		t.Fatalf("unexpected first sector: %+v", profile.Sectors[0])
	}

	if len(profile.Holdings) < 100 {
		t.Fatalf("expected at least 100 holdings, got %d", len(profile.Holdings))
	}

	if profile.Holdings[0].Symbol != "NVDA" || profile.Holdings[0].Weight != 0.0928 {
		t.Fatalf("unexpected first holding: %+v", profile.Holdings[0])
	}

	last := profile.Holdings[len(profile.Holdings)-1]
	if last.Symbol != "TTD" {
		t.Fatalf("expected last holding TTD, got %s", last.Symbol)
	}
}

func TestETFProfileString(t *testing.T) {
	profile := ETFProfile{
		NetAssets:         1000,
		NetExpenseRatio:   0.001,
		DividendYield:     0.02,
		InceptionDate:     "2000-01-01",
		Leveraged:         "NO",
		Sectors:           []ETFProfileSector{{Sector: "TECH", Weight: 0.5}},
		Holdings:          []ETFProfileHolding{{Symbol: "ABC", Description: "ABC INC", Weight: 0.1}},
		PortfolioTurnover: "n/a",
	}

	out := profile.String()
	wantSnippets := []string{
		"Net Assets: 1000",
		"Expense Ratio: 0.0010",
		"Sectors (weight)",
		"TECH: 0.500",
		"Holdings (weight)",
		"ABC (ABC INC): 0.1000",
	}

	for _, snippet := range wantSnippets {
		if !strings.Contains(out, snippet) {
			t.Fatalf("String() missing expected snippet %q\nOutput:\n%s", snippet, out)
		}
	}
}
