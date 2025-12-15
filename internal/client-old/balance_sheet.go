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

// GetBalanceSheet retrieves annual and quarterly balance sheets for the given symbol.
// The endpoint requires function=BALANCE_SHEET and a stock symbol.
func (c *Client) GetBalanceSheet(params models.BalanceSheetParams) (*models.BalanceSheetResponse, error) {
	symbol := strings.TrimSpace(params.Symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("function", "BALANCE_SHEET")
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

	var sheet models.BalanceSheetResponse
	if err := json.Unmarshal(data, &sheet); err != nil {
		return nil, err
	}

	if sheet.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve balance sheet for symbol %s", symbol)
	}

	return &sheet, nil
}
