package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestCashFlowUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/cash_flow_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp CashFlowResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal cash flow response: %v", err)
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
	if annual.OperatingCashflow != "13445000000" {
		t.Fatalf("expected annual operatingCashflow 13445000000, got %s", annual.OperatingCashflow)
	}
	if annual.PaymentsForOperatingActivities != "None" {
		t.Fatalf("expected annual paymentsForOperatingActivities None, got %s", annual.PaymentsForOperatingActivities)
	}
	if annual.CapitalExpenditures != "1685000000" {
		t.Fatalf("expected annual capitalExpenditures 1685000000, got %s", annual.CapitalExpenditures)
	}
	if annual.ChangeInInventory != "-166000000" {
		t.Fatalf("expected annual changeInInventory -166000000, got %s", annual.ChangeInInventory)
	}
	if annual.NetIncome != "6023000000" {
		t.Fatalf("expected annual netIncome 6023000000, got %s", annual.NetIncome)
	}

	quarterly := resp.QuarterlyReports[0]
	if quarterly.FiscalDateEnding != "2025-09-30" {
		t.Fatalf("expected quarterly fiscalDateEnding 2025-09-30, got %s", quarterly.FiscalDateEnding)
	}
	if quarterly.OperatingCashflow != "3082000000" {
		t.Fatalf("expected quarterly operatingCashflow 3082000000, got %s", quarterly.OperatingCashflow)
	}
	if quarterly.DepreciationDepletionAndAmortization != "-2442000000" {
		t.Fatalf("expected quarterly depreciationDepletionAndAmortization -2442000000, got %s", quarterly.DepreciationDepletionAndAmortization)
	}
	if quarterly.NetIncome != "-3249000000" {
		t.Fatalf("expected quarterly netIncome -3249000000, got %s", quarterly.NetIncome)
	}
}

func TestCashFlowString(t *testing.T) {
	data, err := os.ReadFile("testdata/cash_flow_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp CashFlowResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal cash flow response: %v", err)
	}

	rendered := resp.String()

	required := []string{
		"Cash Flow for IBM",
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
