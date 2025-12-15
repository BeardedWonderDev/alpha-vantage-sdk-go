package fundamentaldata

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

// Dividends retrieves historical and declared dividends for a symbol.
func (c *FundamentalDataService) Dividends(symbol string) (*types.DividendsResponse, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("DIVIDENDS", queryParams)
	if err != nil {
		return nil, err
	}

	var divs types.DividendsResponse
	if err := json.Unmarshal(data, &divs); err != nil {
		return nil, err
	}

	return &divs, nil
}

// Splits retrieves historical split events for a symbol.
func (c *FundamentalDataService) Splits(symbol string) (*types.SplitsResponse, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("SPLITS", queryParams)
	if err != nil {
		return nil, err
	}

	var splits types.SplitsResponse
	if err := json.Unmarshal(data, &splits); err != nil {
		return nil, err
	}

	return &splits, nil
}
