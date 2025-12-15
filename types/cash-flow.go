package types

import (
	"fmt"
	"strings"
)

// CashFlowResponse models the CASH_FLOW API response.
type CashFlowResponse struct {
	Symbol           string           `json:"symbol"`
	AnnualReports    []CashFlowReport `json:"annualReports"`
	QuarterlyReports []CashFlowReport `json:"quarterlyReports"`
}

// CashFlowReport represents a single annual or quarterly cash flow entry.
//
// Alpha Vantage returns most numeric values as strings and may return "None" for missing fields.
// Fields are therefore represented as strings so callers can apply their own parsing rules.
type CashFlowReport struct {
	FiscalDateEnding string `json:"fiscalDateEnding"`
	ReportedCurrency string `json:"reportedCurrency"`

	OperatingCashflow                    string `json:"operatingCashflow"`
	PaymentsForOperatingActivities       string `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities      string `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities         string `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets              string `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization string `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                  string `json:"capitalExpenditures"`
	ChangeInReceivables                  string `json:"changeInReceivables"`
	ChangeInInventory                    string `json:"changeInInventory"`
	ProfitLoss                           string `json:"profitLoss"`

	CashflowFromInvestment string `json:"cashflowFromInvestment"`
	CashflowFromFinancing  string `json:"cashflowFromFinancing"`

	ProceedsFromRepaymentsOfShortTermDebt                     string `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            string `json:"dividendPayout"`
	DividendPayoutCommonStock                                 string `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            string `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`

	ChangeInCashAndCashEquivalents string `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate           string `json:"changeInExchangeRate"`
	NetIncome                      string `json:"netIncome"`
}

// String renders a concise summary of the cash flow response.
func (r CashFlowResponse) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Cash Flow for %s\n", r.Symbol))

	if len(r.AnnualReports) > 0 {
		a := r.AnnualReports[0]
		sb.WriteString(fmt.Sprintf("Latest Annual (%s %s): Operating Cashflow %s | CapEx %s | Net Income %s\n", a.FiscalDateEnding, a.ReportedCurrency, a.OperatingCashflow, a.CapitalExpenditures, a.NetIncome))
	}

	if len(r.QuarterlyReports) > 0 {
		q := r.QuarterlyReports[0]
		sb.WriteString(fmt.Sprintf("Latest Quarterly (%s %s): Operating Cashflow %s | CapEx %s | Net Income %s\n", q.FiscalDateEnding, q.ReportedCurrency, q.OperatingCashflow, q.CapitalExpenditures, q.NetIncome))
	}

	sb.WriteString(fmt.Sprintf("Annual Reports: %d | Quarterly Reports: %d\n", len(r.AnnualReports), len(r.QuarterlyReports)))

	return sb.String()
}
