package corestocks

import (
	"encoding/json"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

// Quote retrieves the quote endpoint based on the provided parameters.
// It returns a Quote and an error if there is any.
func (c *CoreStucksService) Quote(params types.TimeSeriesParams) (types.Quote, error) {
	data, err := c.getTimeSeriesData("GLOBAL_QUOTE", params)
	if err != nil {
		return types.Quote{}, err
	}

	var quote types.Quote
	err = json.Unmarshal(data, &quote)
	if err != nil {
		return types.Quote{}, err
	}
	return quote, nil
}
