package types

import (
	"fmt"
	"strings"
)

// IncomeStatementResponse models the INCOME_STATEMENT API response.
type IncomeStatementResponse struct {
	Symbol           string                  `json:"symbol"`
	AnnualReports    []IncomeStatementReport `json:"annualReports"`
	QuarterlyReports []IncomeStatementReport `json:"quarterlyReports"`
}

// IncomeStatementReport represents a single annual or quarterly income statement entry.
//
// Alpha Vantage returns most numeric values as strings and may return "None" for missing fields.
// Fields are therefore represented as strings so callers can apply their own parsing rules.
type IncomeStatementReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       string `json:"grossProfit"`
	TotalRevenue                      string `json:"totalRevenue"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold        string `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   string `json:"operatingIncome"`
	SellingGeneralAndAdministrative   string `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment"`
	OperatingExpenses                 string `json:"operatingExpenses"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	InterestIncome                    string `json:"interestIncome"`
	InterestExpense                   string `json:"interestExpense"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	Depreciation                      string `json:"depreciation"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	InterestAndDebtExpense            string `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax"`
	EBIT                              string `json:"ebit"`
	EBITDA                            string `json:"ebitda"`
	NetIncome                         string `json:"netIncome"`
}

// String renders a concise summary of the income statement response.
func (r IncomeStatementResponse) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Income Statement for %s\n", r.Symbol))

	if len(r.AnnualReports) > 0 {
		a := r.AnnualReports[0]
		sb.WriteString(fmt.Sprintf("Latest Annual (%s %s): Revenue %s | Net Income %s\n", a.FiscalDateEnding, a.ReportedCurrency, a.TotalRevenue, a.NetIncome))
	}

	if len(r.QuarterlyReports) > 0 {
		q := r.QuarterlyReports[0]
		sb.WriteString(fmt.Sprintf("Latest Quarterly (%s %s): Revenue %s | Net Income %s\n", q.FiscalDateEnding, q.ReportedCurrency, q.TotalRevenue, q.NetIncome))
	}

	sb.WriteString(fmt.Sprintf("Annual Reports: %d | Quarterly Reports: %d\n", len(r.AnnualReports), len(r.QuarterlyReports)))

	return sb.String()
}
