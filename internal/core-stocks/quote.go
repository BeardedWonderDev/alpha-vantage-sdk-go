package corestocks

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

// Quote retrieves the quote endpoint based on the provided parameters.
// It returns a Quote and an error if there is any.
func (c *CoreStucksService) Quote(symbol string) (types.Quote, error) {
	symbol = strings.TrimSpace(symbol)
	if symbol == "" {
		return types.Quote{}, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	data, err := c.client.Do("GLOBAL_QUOTE", queryParams)
	if err != nil {
		return types.Quote{}, err
	}

	var quote types.Quote
	if err := json.Unmarshal(data, &quote); err != nil {
		return types.Quote{}, err
	}

	return quote, nil
}
