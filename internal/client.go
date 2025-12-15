package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	corestocks "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/core-stocks"
	technicalindicators "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/technical-indicators"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

const alphaVantageURL = "https://www.alphavantage.co/query"

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string, httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return Client{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (c Client) CoreStocks() types.CoreStocks {
	return corestocks.NewCoreStocksService(c)
}

func (c Client) OptionsData() types.OptionsData {
	return nil
}

func (c Client) AlphaInteligence() types.AlphaInteligence {
	return nil
}

func (c Client) FundamentalData() types.FundamentalData {
	return nil
}

func (c Client) Forex() types.Forex {
	return nil
}

func (c Client) Crypto() types.Crypto {
	return nil
}

func (c Client) Commodities() types.Commodities {
	return nil
}

func (c Client) EconomicIndicators() types.EconomicIndicators {
	return nil
}

func (c Client) TechnicalIndicators() types.TechnicalIndicators {
	return technicalindicators.NewTechnicalIndicatorsService(c)
}

func (c Client) Do(function string, params url.Values) ([]byte, error) {
	query := url.Values{}
	query.Add("function", function)
	query.Add("datatype", "json")

	for key := range params {
		query.Add(key, params.Get(key))
	}

	query.Add("apikey", c.apiKey)

	resp, err := c.httpClient.Get(alphaVantageURL + "?" + query.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := detectAPIMessage(data); err != nil {
		return nil, err
	}

	return data, nil
}

// detectAPIMessage inspects a raw Alpha Vantage response for top-level
// informational or error messages (e.g., rate limits, premium endpoint notices)
// and converts them into Go errors for callers.
func detectAPIMessage(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		// If the payload isn't a JSON object, let the caller's unmarshal handle it.
		return nil
	}

	for _, key := range []string{"Information", "Note", "Error Message"} {
		if v, ok := raw[key]; ok {
			if msg, ok := v.(string); ok && strings.TrimSpace(msg) != "" {
				return fmt.Errorf("alpha vantage %s: %s", strings.ToLower(key), msg)
			}
		}
	}

	return nil
}
