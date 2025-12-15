package crypto

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

func (c *CryptoService) Intraday(params types.CryptoIntradayParams) (*types.CryptoSeriesResponse, error) {
	symbol := strings.TrimSpace(params.Symbol)
	market := strings.TrimSpace(params.Market)
	interval := strings.TrimSpace(params.Interval)
	outputSize := strings.TrimSpace(params.OutputSize)

	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if market == "" {
		return nil, fmt.Errorf("market is required")
	}
	if interval == "" {
		return nil, fmt.Errorf("interval is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)
	queryParams.Add("market", market)
	queryParams.Add("interval", interval)
	if outputSize != "" {
		queryParams.Add("outputsize", outputSize)
	}

	data, err := c.client.Do("CRYPTO_INTRADAY", queryParams)
	if err != nil {
		return nil, err
	}

	cryptoData := &types.CryptoSeriesResponse{}
	if err := types.UnmarshalCryptoJSON(cryptoData, data); err != nil {
		return nil, err
	}

	return cryptoData, nil
}

func (c *CryptoService) Daily(params types.CryptoDailyParams) (*types.CryptoSeriesResponse, error) {
	symbol := strings.TrimSpace(params.Symbol)
	market := strings.TrimSpace(params.Market)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if market == "" {
		return nil, fmt.Errorf("market is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)
	queryParams.Add("market", market)

	data, err := c.client.Do("DIGITAL_CURRENCY_DAILY", queryParams)
	if err != nil {
		return nil, err
	}

	cryptoData := &types.CryptoSeriesResponse{}
	if err := types.UnmarshalCryptoJSON(cryptoData, data); err != nil {
		return nil, err
	}

	return cryptoData, nil
}

func (c *CryptoService) Weekly(params types.CryptoWeeklyParams) (*types.CryptoSeriesResponse, error) {
	symbol := strings.TrimSpace(params.Symbol)
	market := strings.TrimSpace(params.Market)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if market == "" {
		return nil, fmt.Errorf("market is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)
	queryParams.Add("market", market)

	data, err := c.client.Do("DIGITAL_CURRENCY_WEEKLY", queryParams)
	if err != nil {
		return nil, err
	}

	cryptoData := &types.CryptoSeriesResponse{}
	if err := types.UnmarshalCryptoJSON(cryptoData, data); err != nil {
		return nil, err
	}

	return cryptoData, nil
}

func (c *CryptoService) Monthly(params types.CryptoMonthlyParams) (*types.CryptoSeriesResponse, error) {
	symbol := strings.TrimSpace(params.Symbol)
	market := strings.TrimSpace(params.Market)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if market == "" {
		return nil, fmt.Errorf("market is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)
	queryParams.Add("market", market)

	data, err := c.client.Do("DIGITAL_CURRENCY_MONTHLY", queryParams)
	if err != nil {
		return nil, err
	}

	cryptoData := &types.CryptoSeriesResponse{}
	if err := types.UnmarshalCryptoJSON(cryptoData, data); err != nil {
		return nil, err
	}

	return cryptoData, nil
}
