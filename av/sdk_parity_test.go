package av_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/av"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestCoreStocks_Quote_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture := []byte(`{
  "Global Quote": {
    "01. symbol": "IBM",
    "02. open": "100.0",
    "03. high": "110.0",
    "04. low": "90.0",
    "05. price": "105.0",
    "06. volume": "123",
    "07. latest trading day": "2025-12-12",
    "08. previous close": "99.0",
    "09. change": "6.0",
    "10. change percent": "6.06%"
  }
}`)

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "GLOBAL_QUOTE" {
				return nil, fmt.Errorf("expected function GLOBAL_QUOTE, got %q", q.Get("function"))
			}
			if q.Get("symbol") != "IBM" {
				return nil, fmt.Errorf("expected symbol IBM, got %q", q.Get("symbol"))
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}
			if q.Get("apikey") != "test-key" {
				return nil, fmt.Errorf("expected apikey test-key, got %q", q.Get("apikey"))
			}
			if q.Get("interval") != "" {
				return nil, fmt.Errorf("did not expect interval param, got %q", q.Get("interval"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	quote, err := cli.CoreStocks().Quote("IBM")
	if err != nil {
		t.Fatalf("Quote returned error: %v", err)
	}
	if quote.Symbol != "IBM" {
		t.Fatalf("expected IBM, got %q", quote.Symbol)
	}
}

func TestForex_ExchangeRate_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture := []byte(`{
  "Realtime Currency Exchange Rate": {
    "1. From_Currency Code": "USD",
    "2. From_Currency Name": "United States Dollar",
    "3. To_Currency Code": "EUR",
    "4. To_Currency Name": "Euro",
    "5. Exchange Rate": "0.9000",
    "6. Last Refreshed": "2025-12-12 16:00:00",
    "7. Time Zone": "UTC",
    "8. Bid Price": "0.8990",
    "9. Ask Price": "0.9010"
  }
}`)

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "CURRENCY_EXCHANGE_RATE" {
				return nil, fmt.Errorf("expected function CURRENCY_EXCHANGE_RATE, got %q", q.Get("function"))
			}
			if q.Get("from_currency") != "USD" {
				return nil, fmt.Errorf("expected from_currency USD, got %q", q.Get("from_currency"))
			}
			if q.Get("to_currency") != "EUR" {
				return nil, fmt.Errorf("expected to_currency EUR, got %q", q.Get("to_currency"))
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}
			if q.Get("apikey") != "test-key" {
				return nil, fmt.Errorf("expected apikey test-key, got %q", q.Get("apikey"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	rate, err := cli.Forex().ExchangeRate(types.ForexExchangeRateParams{FromCurrency: "USD", ToCurrency: "EUR"})
	if err != nil {
		t.Fatalf("ExchangeRate returned error: %v", err)
	}
	if rate.ExchangeRateInfo.FromCurrencyCode != "USD" {
		t.Fatalf("expected USD, got %q", rate.ExchangeRateInfo.FromCurrencyCode)
	}
	if rate.ExchangeRateInfo.ToCurrencyCode != "EUR" {
		t.Fatalf("expected EUR, got %q", rate.ExchangeRateInfo.ToCurrencyCode)
	}
}

func TestCrypto_Daily_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture := []byte(`{
  "Meta Data": {
    "1. Information": "Daily Prices and Volumes for Digital Currency",
    "2. Digital Currency Code": "BTC",
    "3. Digital Currency Name": "Bitcoin",
    "4. Market Code": "USD",
    "5. Market Name": "United States Dollar",
    "6. Last Refreshed": "2025-12-12",
    "7. Time Zone": "UTC"
  },
  "Time Series (Digital Currency Daily)": {
    "2025-12-12": {
      "1a. open (USD)": "100.0",
      "2a. high (USD)": "110.0",
      "3a. low (USD)": "90.0",
      "4a. close (USD)": "105.0",
      "5. volume": "1.0",
      "6. market cap (USD)": "2.0"
    }
  }
}`)

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "DIGITAL_CURRENCY_DAILY" {
				return nil, fmt.Errorf("expected function DIGITAL_CURRENCY_DAILY, got %q", q.Get("function"))
			}
			if q.Get("symbol") != "BTC" {
				return nil, fmt.Errorf("expected symbol BTC, got %q", q.Get("symbol"))
			}
			if q.Get("market") != "USD" {
				return nil, fmt.Errorf("expected market USD, got %q", q.Get("market"))
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}
			if q.Get("apikey") != "test-key" {
				return nil, fmt.Errorf("expected apikey test-key, got %q", q.Get("apikey"))
			}
			if q.Get("interval") != "" {
				return nil, fmt.Errorf("did not expect interval param, got %q", q.Get("interval"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	resp, err := cli.Crypto().Daily(types.CryptoDailyParams{Symbol: "BTC", Market: "USD"})
	if err != nil {
		t.Fatalf("Daily returned error: %v", err)
	}
	if resp.MetaData.DigitalCurrencyCode != "BTC" {
		t.Fatalf("expected BTC, got %q", resp.MetaData.DigitalCurrencyCode)
	}
	if len(resp.TimeSeries) != 1 {
		t.Fatalf("expected 1 time series entry, got %d", len(resp.TimeSeries))
	}
}

func TestAlphaInteligence_AnalyticsFixedWindow_UsesRangeParams(t *testing.T) {
	fixture, err := os.ReadFile("../models/testdata/analytics_fixed_window.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "ANALYTICS_FIXED_WINDOW" {
				return nil, fmt.Errorf("expected function ANALYTICS_FIXED_WINDOW, got %q", q.Get("function"))
			}
			ranges := q["range"]
			if len(ranges) != 2 || ranges[0] != "2023-07-03" || ranges[1] != "2023-08-31" {
				return nil, fmt.Errorf("expected range values [2023-07-03 2023-08-31], got %v", ranges)
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	_, err = cli.AlphaInteligence().AnalyticsFixedWindow(types.AnalyticsFixedWindowParams{
		Symbols:      "IBM,AAPL,MSFT",
		Range:        []string{"2023-07-03", "2023-08-31"},
		Interval:     "DAILY",
		Calculations: "MEAN,STDDEV,CORRELATION",
	})
	if err != nil {
		t.Fatalf("AnalyticsFixedWindow returned error: %v", err)
	}
}

func TestFundamentalData_ETFProfile_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture, err := os.ReadFile("../models/testdata/etf_profile_QQQ.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "ETF_PROFILE" {
				return nil, fmt.Errorf("expected function ETF_PROFILE, got %q", q.Get("function"))
			}
			if q.Get("symbol") != "QQQ" {
				return nil, fmt.Errorf("expected symbol QQQ, got %q", q.Get("symbol"))
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	profile, err := cli.FundamentalData().ETFProfile("QQQ")
	if err != nil {
		t.Fatalf("ETFProfile returned error: %v", err)
	}
	if profile.NetAssets == 0 {
		t.Fatalf("expected non-zero net assets")
	}
}

func TestFundamentalData_Dividends_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture, err := os.ReadFile("../models/testdata/dividends_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "DIVIDENDS" {
				return nil, fmt.Errorf("expected function DIVIDENDS, got %q", q.Get("function"))
			}
			if q.Get("symbol") != "IBM" {
				return nil, fmt.Errorf("expected symbol IBM, got %q", q.Get("symbol"))
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	resp, err := cli.FundamentalData().Dividends("IBM")
	if err != nil {
		t.Fatalf("Dividends returned error: %v", err)
	}
	if resp.Symbol != "IBM" {
		t.Fatalf("expected IBM, got %q", resp.Symbol)
	}
}

func TestFundamentalData_Splits_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture, err := os.ReadFile("../models/testdata/splits_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "SPLITS" {
				return nil, fmt.Errorf("expected function SPLITS, got %q", q.Get("function"))
			}
			if q.Get("symbol") != "IBM" {
				return nil, fmt.Errorf("expected symbol IBM, got %q", q.Get("symbol"))
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	resp, err := cli.FundamentalData().Splits("IBM")
	if err != nil {
		t.Fatalf("Splits returned error: %v", err)
	}
	if resp.Symbol != "IBM" {
		t.Fatalf("expected IBM, got %q", resp.Symbol)
	}
}

func TestTechnicalIndicators_SMA_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture := []byte(`{
  "Meta Data": {
    "1: Symbol": "IBM",
    "2: Indicator": "Simple Moving Average (SMA)",
    "3: Last Refreshed": "2025-12-12",
    "4: Interval": "daily",
    "5: Time Period": 20,
    "6: Series Type": "close",
    "7: Time Zone": "UTC"
  },
  "Technical Analysis: SMA": {
    "2025-12-12 16:00": {
      "SMA": "123.45"
    }
  }
}`)

	httpClient := &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			q := req.URL.Query()
			if q.Get("function") != "SMA" {
				return nil, fmt.Errorf("expected function SMA, got %q", q.Get("function"))
			}
			if q.Get("symbol") != "IBM" {
				return nil, fmt.Errorf("expected symbol IBM, got %q", q.Get("symbol"))
			}
			if q.Get("interval") != "daily" {
				return nil, fmt.Errorf("expected interval daily, got %q", q.Get("interval"))
			}
			if q.Get("time_period") != "20" {
				return nil, fmt.Errorf("expected time_period 20, got %q", q.Get("time_period"))
			}
			if q.Get("series_type") != "close" {
				return nil, fmt.Errorf("expected series_type close, got %q", q.Get("series_type"))
			}
			if q.Get("datatype") != "json" {
				return nil, fmt.Errorf("expected datatype json, got %q", q.Get("datatype"))
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(fixture)),
				Header:     make(http.Header),
				Request:    req,
			}, nil
		}),
	}

	cli := av.NewClientWithHTTPClient("test-key", httpClient)
	resp, err := cli.TechnicalIndicators().SMA(types.IndicatorParams{
		Symbol:     "IBM",
		Interval:   "daily",
		TimePeriod: 20,
		SeriesType: "close",
	})
	if err != nil {
		t.Fatalf("SMA returned error: %v", err)
	}
	if len(resp.IndicatorValues) != 1 {
		t.Fatalf("expected 1 indicator value, got %d", len(resp.IndicatorValues))
	}
}
