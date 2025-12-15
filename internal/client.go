package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	alphainteligence "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/alpha-inteligence"
	corestocks "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/core-stocks"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/crypto"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/forex"
	fundamentaldata "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/fundamental-data"
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
	return alphainteligence.NewAlphaInteligenceService(c)
}

func (c Client) FundamentalData() types.FundamentalData {
	return fundamentaldata.NewFundamentalDataService(c)
}

func (c Client) Forex() types.Forex {
	return forex.NewForexService(c)
}

func (c Client) Crypto() types.Crypto {
	return crypto.NewCryptoService(c)
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

	for key, values := range params {
		for _, v := range values {
			if strings.TrimSpace(v) != "" {
				query.Add(key, v)
			}
		}
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
