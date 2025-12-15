package fundamentaldata

import (
	"encoding/json"
	"net/url"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

// Dividends retrieves historical and declared dividends for a symbol.
func (c *FundamentalDataService) Dividends(symbol string) (*types.DividendsResponse, error) {
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
