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

// GetCashFlow retrieves annual and quarterly cash flow statements for the given symbol.
// The endpoint requires function=CASH_FLOW and a stock symbol.
func (c *Client) GetCashFlow(params models.CashFlowParams) (*models.CashFlowResponse, error) {
	symbol := strings.TrimSpace(params.Symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("function", "CASH_FLOW")
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

	var cashFlow models.CashFlowResponse
	if err := json.Unmarshal(data, &cashFlow); err != nil {
		return nil, err
	}

	if cashFlow.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve cash flow for symbol %s", symbol)
	}

	return &cashFlow, nil
}
