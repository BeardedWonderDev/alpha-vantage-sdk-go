package models

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestIncomeStatementUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/income_statement_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp IncomeStatementResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal income statement response: %v", err)
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
	if annual.TotalRevenue != "62753000000" {
		t.Fatalf("expected annual totalRevenue 62753000000, got %s", annual.TotalRevenue)
	}
	if annual.NetInterestIncome != "-965000000" {
		t.Fatalf("expected annual netInterestIncome -965000000, got %s", annual.NetInterestIncome)
	}
	if annual.InvestmentIncomeNet != "None" {
		t.Fatalf("expected annual investmentIncomeNet None, got %s", annual.InvestmentIncomeNet)
	}

	quarterly := resp.QuarterlyReports[0]
	if quarterly.FiscalDateEnding != "2025-09-30" {
		t.Fatalf("expected quarterly fiscalDateEnding 2025-09-30, got %s", quarterly.FiscalDateEnding)
	}
	if quarterly.TotalRevenue != "16331000000" {
		t.Fatalf("expected quarterly totalRevenue 16331000000, got %s", quarterly.TotalRevenue)
	}
}

func TestIncomeStatementString(t *testing.T) {
	data, err := os.ReadFile("testdata/income_statement_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var resp IncomeStatementResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Fatalf("failed to unmarshal income statement response: %v", err)
	}

	rendered := resp.String()

	required := []string{
		"Income Statement for IBM",
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
