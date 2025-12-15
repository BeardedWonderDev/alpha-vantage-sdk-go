package fundamentaldata

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

// CompanyOverview retrieves the Alpha Vantage company overview for the given symbol.
func (c *FundamentalDataService) CompanyOverview(symbol string) (*types.CompanyOverviewResponse, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("OVERVIEW", queryParams)
	if err != nil {
		return nil, err
	}

	var overview types.CompanyOverviewResponse
	if err := json.Unmarshal(data, &overview); err != nil {
		return nil, err
	}

	if overview.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve company overview for symbol %s", symbol)
	}

	return &overview, nil
}

// IncomeStatement retrieves annual and quarterly income statements for the given symbol.
func (c *FundamentalDataService) IncomeStatement(symbol string) (*types.IncomeStatementResponse, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("INCOME_STATEMENT", queryParams)
	if err != nil {
		return nil, err
	}

	var statement types.IncomeStatementResponse
	if err := json.Unmarshal(data, &statement); err != nil {
		return nil, err
	}

	if statement.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve income statement for symbol %s", symbol)
	}

	return &statement, nil
}

// BalanceSheet retrieves annual and quarterly balance sheets for the given symbol.
// The endpoint requires function=BALANCE_SHEET and a stock symbol.
func (c *FundamentalDataService) BalanceSheet(symbol string) (*types.BalanceSheetResponse, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("BALANCE_SHEET", queryParams)
	if err != nil {
		return nil, err
	}

	var sheet types.BalanceSheetResponse
	if err := json.Unmarshal(data, &sheet); err != nil {
		return nil, err
	}

	if sheet.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve balance sheet for symbol %s", symbol)
	}

	return &sheet, nil
}

// CashFlow retrieves annual and quarterly cash flow statements for the given symbol.
// The endpoint requires function=CASH_FLOW and a stock symbol.
func (c *FundamentalDataService) CashFlow(symbol string) (*types.CashFlowResponse, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("CASH_FLOW", queryParams)
	if err != nil {
		return nil, err
	}

	var cashFlow types.CashFlowResponse
	if err := json.Unmarshal(data, &cashFlow); err != nil {
		return nil, err
	}

	if cashFlow.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve cash flow for symbol %s", symbol)
	}

	return &cashFlow, nil
}
