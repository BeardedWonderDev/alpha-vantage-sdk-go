package forex

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

// ExchangeRate retrieves real-time exchange rates for physical currencies.
func (c *ForexService) ExchangeRate(params types.ForexExchangeRateParams) (*types.CurrencyExchangeRateResponse, error) {
	from := strings.TrimSpace(params.FromCurrency)
	to := strings.TrimSpace(params.ToCurrency)
	if from == "" {
		return nil, fmt.Errorf("from currency is required")
	}
	if to == "" {
		return nil, fmt.Errorf("to currency is required")
	}

	queryParams := url.Values{}
	queryParams.Add("from_currency", from)
	queryParams.Add("to_currency", to)

	data, err := c.client.Do("CURRENCY_EXCHANGE_RATE", queryParams)
	if err != nil {
		return nil, err
	}

	var resp types.CurrencyExchangeRateResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
