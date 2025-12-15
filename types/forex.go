package types

// ForexExchangeRateParams defines the request parameters for the CURRENCY_EXCHANGE_RATE endpoint.
type ForexExchangeRateParams struct {
	FromCurrency string
	ToCurrency   string
}

// CurrencyExchangeRateResponse models the response returned by the
// CURRENCY_EXCHANGE_RATE endpoint. This response shape is used by both
// the Forex and Crypto exchange rate endpoints.
type CurrencyExchangeRateResponse struct {
	ExchangeRateInfo ExchangeRateInfo `json:"Realtime Currency Exchange Rate"`
}

type ExchangeRateInfo struct {
	FromCurrencyCode string `json:"1. From_Currency Code"`
	FromCurrencyName string `json:"2. From_Currency Name"`
	ToCurrencyCode   string `json:"3. To_Currency Code"`
	ToCurrencyName   string `json:"4. To_Currency Name"`
	ExchangeRate     string `json:"5. Exchange Rate"`
	LastRefreshed    string `json:"6. Last Refreshed"`
	TimeZone         string `json:"7. Time Zone"`
	BidPrice         string `json:"8. Bid Price"`
	AskPrice         string `json:"9. Ask Price"`
}

