package corestocks

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

// SymbolSearch retrieves the best-matching symbols and market information based on keywords.
func (c *CoreStucksService) SymbolSearch(params types.SymbolSearchParams) (*types.SymbolSearchResponse, error) {
	keywords := strings.TrimSpace(params.Keywords)
	if keywords == "" {
		return nil, fmt.Errorf("keywords is required")
	}

	queryParams := url.Values{}
	queryParams.Add("keywords", keywords)

	data, err := c.client.Do("SYMBOL_SEARCH", queryParams)
	if err != nil {
		return nil, err
	}

	var search types.SymbolSearchResponse
	if err := json.Unmarshal(data, &search); err != nil {
		return nil, err
	}

	return &search, nil
}
