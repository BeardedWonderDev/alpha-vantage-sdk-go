package fundamentaldata

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

// ETFProfile retrieves ETF profile and holdings for the given symbol.
func (c *FundamentalDataService) ETFProfile(symbol string) (*types.ETFProfile, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("ETF_PROFILE", queryParams)
	if err != nil {
		return nil, err
	}

	var profile types.ETFProfile
	if err := json.Unmarshal(data, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}
