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

// GetSplitsData retrieves raw split history data (JSON or CSV) based on the provided parameters.
// The endpoint requires function=SPLITS and a stock symbol.
func (c *Client) GetSplitsData(params models.SplitsParams) ([]byte, error) {
	symbol := strings.TrimSpace(params.Symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("function", "SPLITS")
	queryParams.Add("symbol", symbol)

	if strings.TrimSpace(params.DataType) != "" {
		dataType := strings.ToLower(strings.TrimSpace(params.DataType))
		switch dataType {
		case "json", "csv":
			queryParams.Add("datatype", dataType)
		default:
			return nil, fmt.Errorf("datatype must be \"json\" or \"csv\"")
		}
	}
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

	return data, nil
}

// GetSplits retrieves historical split events for a symbol.
// This method only supports datatype=json. For CSV output, use GetSplitsData.
func (c *Client) GetSplits(params models.SplitsParams) (*models.SplitsResponse, error) {
	if strings.EqualFold(strings.TrimSpace(params.DataType), "csv") {
		return nil, fmt.Errorf("datatype csv is not supported for GetSplits; use GetSplitsData")
	}

	data, err := c.GetSplitsData(params)
	if err != nil {
		return nil, err
	}

	var splits models.SplitsResponse
	if err := json.Unmarshal(data, &splits); err != nil {
		return nil, err
	}

	return &splits, nil
}
