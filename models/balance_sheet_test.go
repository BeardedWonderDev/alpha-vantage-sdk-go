package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestBalanceSheetUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/balance_sheet_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp BalanceSheetResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal balance sheet response: %v", err)
	}

	if resp.Symbol != "IBM" {
		t.Fatalf("expected symbol IBM, got %s", resp.Symbol)
	}

	if len(resp.AnnualReports) != 2 {
		t.Fatalf("expected 2 annual reports, got %d", len(resp.AnnualReports))
	}
	if len(resp.QuarterlyReports) != 2 {
		t.Fatalf("expected 2 quarterly reports, got %d", len(resp.QuarterlyReports))
	}

	annual := resp.AnnualReports[0]
	if annual.FiscalDateEnding != "2024-12-31" {
		t.Fatalf("expected annual fiscalDateEnding 2024-12-31, got %s", annual.FiscalDateEnding)
	}
	if annual.ReportedCurrency != "USD" {
		t.Fatalf("expected annual reportedCurrency USD, got %s", annual.ReportedCurrency)
	}
	if annual.TotalAssets != "137175000000" {
		t.Fatalf("expected annual totalAssets 137175000000, got %s", annual.TotalAssets)
	}
	if annual.CashAndCashEquivalentsAtCarryingValue != "13947000000" {
		t.Fatalf("expected annual cashAndCashEquivalentsAtCarryingValue 13947000000, got %s", annual.CashAndCashEquivalentsAtCarryingValue)
	}
	if annual.AccumulatedDepreciationAmortizationPPE != "None" {
		t.Fatalf("expected annual accumulatedDepreciationAmortizationPPE None, got %s", annual.AccumulatedDepreciationAmortizationPPE)
	}
	if annual.TotalLiabilities != "109782000000" {
		t.Fatalf("expected annual totalLiabilities 109782000000, got %s", annual.TotalLiabilities)
	}
	if annual.TotalShareholderEquity != "27307000000" {
		t.Fatalf("expected annual totalShareholderEquity 27307000000, got %s", annual.TotalShareholderEquity)
	}

	quarterly := resp.QuarterlyReports[0]
	if quarterly.FiscalDateEnding != "2025-09-30" {
		t.Fatalf("expected quarterly fiscalDateEnding 2025-09-30, got %s", quarterly.FiscalDateEnding)
	}
	if quarterly.TotalAssets != "146312000000" {
		t.Fatalf("expected quarterly totalAssets 146312000000, got %s", quarterly.TotalAssets)
	}
	if quarterly.ShortTermInvestments != "3286000000" {
		t.Fatalf("expected quarterly shortTermInvestments 3286000000, got %s", quarterly.ShortTermInvestments)
	}
}

func TestBalanceSheetString(t *testing.T) {
	data, err := os.ReadFile("testdata/balance_sheet_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp BalanceSheetResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal balance sheet response: %v", err)
	}

	rendered := resp.String()

	required := []string{
		"Balance Sheet for IBM",
		"Latest Annual (2024-12-31 USD)",
		"Latest Quarterly (2025-09-30 USD)",
		"Annual Reports: 2",
		"Quarterly Reports: 2",
	}

	for _, snippet := range required {
		if !strings.Contains(rendered, snippet) {
			t.Fatalf("String() missing expected snippet %q\nRendered:\n%s", snippet, rendered)
		}
	}
}
