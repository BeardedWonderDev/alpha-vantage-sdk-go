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

// GetSymbolSearchData retrieves raw symbol search results (JSON or CSV) based on the provided parameters.
// The endpoint requires function=SYMBOL_SEARCH and keywords.
func (c *Client) GetSymbolSearchData(params models.SymbolSearchParams) ([]byte, error) {
	keywords := strings.TrimSpace(params.Keywords)
	if keywords == "" {
		return nil, fmt.Errorf("keywords is required")
	}

	queryParams := url.Values{}
	queryParams.Add("function", "SYMBOL_SEARCH")
	queryParams.Add("keywords", keywords)

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

// GetSymbolSearch retrieves the best-matching symbols and market information based on keywords.
// This method only supports datatype=json. For CSV output, use GetSymbolSearchData.
func (c *Client) GetSymbolSearch(params models.SymbolSearchParams) (*models.SymbolSearchResponse, error) {
	if strings.EqualFold(strings.TrimSpace(params.DataType), "csv") {
		return nil, fmt.Errorf("datatype csv is not supported for GetSymbolSearch; use GetSymbolSearchData")
	}

	data, err := c.GetSymbolSearchData(params)
	if err != nil {
		return nil, err
	}

	var search models.SymbolSearchResponse
	if err := json.Unmarshal(data, &search); err != nil {
		return nil, err
	}

	return &search, nil
}
