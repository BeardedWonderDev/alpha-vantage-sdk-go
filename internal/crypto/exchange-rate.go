package crypto

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

// ExchangeRate retrieves the real-time exchange rate between a digital currency and another currency.
func (c *CryptoService) ExchangeRate(params types.CryptoExchangeRateParams) (*types.CurrencyExchangeRateResponse, error) {
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
	if err := types.UnmarshalLenient(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
