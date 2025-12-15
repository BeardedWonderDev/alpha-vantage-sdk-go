package fundamentaldata

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

// CompanyOverview retrieves the Alpha Vantage company overview for the given symbol.
// The endpoint requires function=OVERVIEW and a stock symbol.
func (c *FundamentalDataService) CompanyOverview(params types.CompanyOverviewParams) (*types.CompanyOverview, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", params.Symbol)

	data, err := c.client.Do("OVERVIEW", queryParams)
	if err != nil {
		return nil, err
	}

	var overview types.CompanyOverview
	if err := json.Unmarshal(data, &overview); err != nil {
		return nil, err
	}

	if overview.Symbol == "" {
		return nil, fmt.Errorf("failed to retrieve company overview for symbol %s", params.Symbol)
	}

	return &overview, nil
}
