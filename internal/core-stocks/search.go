package corestocks

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

// SymbolSearch retrieves the best-matching symbols and market information based on keywords.
func (c *CoreStucksService) SymbolSearch(keywords string) (*types.SymbolSearchResponse, error) {
	keywords = strings.TrimSpace(keywords)
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
	if err := types.UnmarshalLenient(data, &search); err != nil {
		return nil, err
	}

	return &search, nil
}
