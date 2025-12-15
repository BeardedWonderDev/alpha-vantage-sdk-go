package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestCompanyOverviewUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/company_overview_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var overview CompanyOverview
	if err := json.Unmarshal(data, &overview); err != nil {
		t.Fatalf("failed to unmarshal company overview: %v", err)
	}

	if overview.Symbol != "IBM" {
		t.Fatalf("expected symbol IBM, got %s", overview.Symbol)
	}

	if overview.Name != "International Business Machines" {
		t.Fatalf("unexpected company name: %s", overview.Name)
	}

	if overview.MarketCapitalization != 292263690000 {
		t.Fatalf("expected market cap 292263690000, got %d", overview.MarketCapitalization)
	}

	if overview.EPS != 8.38 {
		t.Fatalf("expected EPS 8.38, got %f", overview.EPS)
	}

	if overview.Week52High != 324.9 || overview.Week52Low != 209.2 {
		t.Fatalf("unexpected 52 week range: high %f low %f", overview.Week52High, overview.Week52Low)
	}
}

func TestCompanyOverviewString(t *testing.T) {
	overview := CompanyOverview{
		Symbol:               "IBM",
		Name:                 "International Business Machines",
		Exchange:             "NYSE",
		Sector:               "TECHNOLOGY",
		Industry:             "INFORMATION TECHNOLOGY SERVICES",
		MarketCapitalization: 292263690000,
		Currency:             "USD",
		EPS:                  8.38,
		PERatio:              37.31,
		PEGRatio:             2.14,
		DividendPerShare:     6.7,
		DividendYield:        0.0216,
		LatestQuarter:        "2025-09-30",
		Week52High:           324.9,
		Week52Low:            209.2,
		MovingAverage50Day:   297.6,
		MovingAverage200Day:  267.92,
	}

	rendered := overview.String()

	required := []string{
		"International Business Machines",
		"IBM",
		"Market Cap",
		"EPS: 8.38",
		"P/E: 37.31",
		"52W High/Low: 324.90 / 209.20",
	}

	for _, snippet := range required {
		if !strings.Contains(rendered, snippet) {
			t.Fatalf("String() missing expected snippet %q\nRendered:\n%s", snippet, rendered)
		}
	}
}
