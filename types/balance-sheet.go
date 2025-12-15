package types

import (
	"fmt"
	"strings"
)

// BalanceSheetResponse models the BALANCE_SHEET API response.
type BalanceSheetResponse struct {
	Symbol           string               `json:"symbol"`
	AnnualReports    []BalanceSheetReport `json:"annualReports"`
	QuarterlyReports []BalanceSheetReport `json:"quarterlyReports"`
}

// BalanceSheetReport represents a single annual or quarterly balance sheet entry.
//
// Alpha Vantage returns most numeric values as strings and may return "None" for missing fields.
// Fields are therefore represented as strings so callers can apply their own parsing rules.
type BalanceSheetReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            string `json:"totalAssets"`
	TotalCurrentAssets                     string `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  string `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            string `json:"cashAndShortTermInvestments"`
	Inventory                              string `json:"inventory"`
	CurrentNetReceivables                  string `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  string `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 string `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE string `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       string `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      string `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               string `json:"goodwill"`
	Investments                            string `json:"investments"`
	LongTermInvestments                    string `json:"longTermInvestments"`
	ShortTermInvestments                   string `json:"shortTermInvestments"`
	OtherCurrentAssets                     string `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  string `json:"otherNonCurrentAssets"`
	TotalLiabilities                       string `json:"totalLiabilities"`
	TotalCurrentLiabilities                string `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 string `json:"currentAccountsPayable"`
	DeferredRevenue                        string `json:"deferredRevenue"`
	CurrentDebt                            string `json:"currentDebt"`
	ShortTermDebt                          string `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             string `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                string `json:"capitalLeaseObligations"`
	LongTermDebt                           string `json:"longTermDebt"`
	CurrentLongTermDebt                    string `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 string `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 string `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                string `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             string `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 string `json:"totalShareholderEquity"`
	TreasuryStock                          string `json:"treasuryStock"`
	RetainedEarnings                       string `json:"retainedEarnings"`
	CommonStock                            string `json:"commonStock"`
	CommonStockSharesOutstanding           string `json:"commonStockSharesOutstanding"`
}

// String renders a concise summary of the balance sheet response.
func (r BalanceSheetResponse) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Balance Sheet for %s\n", r.Symbol))

	if len(r.AnnualReports) > 0 {
		a := r.AnnualReports[0]
		sb.WriteString(fmt.Sprintf("Latest Annual (%s %s): Total Assets %s | Total Liabilities %s | Equity %s\n", a.FiscalDateEnding, a.ReportedCurrency, a.TotalAssets, a.TotalLiabilities, a.TotalShareholderEquity))
	}

	if len(r.QuarterlyReports) > 0 {
		q := r.QuarterlyReports[0]
		sb.WriteString(fmt.Sprintf("Latest Quarterly (%s %s): Total Assets %s | Total Liabilities %s | Equity %s\n", q.FiscalDateEnding, q.ReportedCurrency, q.TotalAssets, q.TotalLiabilities, q.TotalShareholderEquity))
	}

	sb.WriteString(fmt.Sprintf("Annual Reports: %d | Quarterly Reports: %d\n", len(r.AnnualReports), len(r.QuarterlyReports)))

	return sb.String()
}
