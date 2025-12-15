package client

import (
	"encoding/json"
	"fmt"
	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// GetIncomeStatement retrieves annual and quarterly income statements for the given symbol.
// The endpoint requires function=INCOME_STATEMENT and a stock symbol.
func (c *Client) GetIncomeStatement(params models.IncomeStatementParams) (*models.IncomeStatementResponse, error) {
	symbol := strings.TrimSpace(params.Symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("function", "INCOME_STATEMENT")
	queryParams.Add("symbol", symbol)
	queryParams.Add("apikey", c.apiKey)

	resp, err := http.Get(alphaVantageURL + "?" + queryParams.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := detectAPIMessage(data); err != nil {
		return nil, err
	}

	var statement models.IncomeStatementResponse
	if err := json.Unmarshal(data, &statement); err != nil {
		return nil, err
	}

	if statement.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve income statement for symbol %s", symbol)
	}

	return &statement, nil
}
