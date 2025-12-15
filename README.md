# Alpha Vantage Go Wrapper

![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Version: v1.0.0](https://img.shields.io/badge/version-v1.0.0-blue)

A lightweight, dependency-free Go client for the [Alpha Vantage](https://www.alphavantage.co/) REST API. It exposes typed request/response structs for equities, crypto, technical indicators, and fundamentals while keeping a consistent, ergonomic interface.

** NOTE: ** Client only supports JSON responses at this time

## Table of Contents
- [Quick Start](#quick-start)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage Examples](#usage-examples)
- [Endpoint Coverage](#endpoint-coverage)
  - [Alpha Intelligence](#alpha-intelligence)
  - [Fundamental Data](#fundamental-data)
- [Output Format](#output-format)
- [Development](#development)
- [License](#license)
- [Contact](#contact)

## Quick Start

```bash
export ALPHAVANTAGE_API_KEY=your_api_key
go test ./...   # optional: verifies your environment
```

Embed in your app:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/client"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/models"
)

func main() {
	cli := client.NewClient(os.Getenv("ALPHAVANTAGE_API_KEY"))

	ts, err := cli.GetIntraday(models.TimeSeriesParams{
		Symbol:     "MSFT",
		Interval:   "5min",
		OutputSize: "compact",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ts) // pretty String() output
}
```

## Installation

- Go **1.21+** required.
- Add the module:

```bash
go get github.com/masonJamesWheeler/alpha-vantage-go-wrapper
```

Import packages:

```go
import (
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/client"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/models"
)
```

## Configuration

- **API key**: recommended via `ALPHAVANTAGE_API_KEY` env var; you can also pass the raw string to `client.NewClient`.
- **Data format**: most params accept `DataType` (`json` or `csv`); defaults to JSON when omitted.
- **Rate limits**: the client surfaces Alpha Vantage informational/error messages (e.g., throttling) as Go errors; it does not auto-retry.

## Usage Examples

Each endpoint uses a parameter struct from `models` and a method on `Client`.

### Time Series
```go
daily, err := cli.GetDailyAdjusted(models.TimeSeriesParams{
	Symbol:     "AAPL",
	OutputSize: "full", // or "compact"
})
```

### Cryptocurrency
```go
btc, err := cli.GetCryptoDaily(models.CryptoParams{
	Symbol:  "BTC",
	Market:  "USD",
	Interval:"1min",
})
```

### Technical Indicators
```go
bb, err := cli.GetBBANDS(models.IndicatorParams{
	Symbol:     "MSFT",
	Interval:   "15min",
	TimePeriod: 20,
	SeriesType: "close",
})
```

### Fundamentals & Analytics
```go
overview, _ := cli.GetCompanyOverview(models.CompanyOverviewParams{Symbol: "IBM"})
etf, _ := cli.GetETFProfile(models.ETFProfileParams{Symbol: "QQQ"})
divs, _ := cli.GetDividends(models.DividendsParams{Symbol: "IBM"})

analytics, _ := cli.GetAnalyticsSlidingWindow(models.AnalyticsSlidingWindowParams{
	Symbols:      "AAPL,IBM",
	Range:        []string{"2month"},
	Interval:     "DAILY",
	WindowSize:   20,
	Calculations: "MEAN,STDDEV(annualized=true)",
})

fixed, _ := cli.GetAnalyticsFixedWindow(models.AnalyticsFixedWindowParams{
	Symbols:      "IBM,AAPL,MSFT",
	Range:        []string{"2023-07-03", "2023-08-31"},
	Interval:     "DAILY",
	Calculations: "MEAN,STDDEV,CORRELATION",
})
```

### Symbol Search
```go
search, err := cli.GetSymbolSearch(models.SymbolSearchParams{Keywords: "microsoft"})
if err != nil {
	log.Fatal(err)
}

for _, match := range search.BestMatches {
	fmt.Printf("%s - %s (%s) score=%0.4f\n", match.Symbol, match.Name, match.Region, match.MatchScore)
}

csvBytes, _ := cli.GetSymbolSearchData(models.SymbolSearchParams{
	Keywords: "microsoft",
	DataType: "csv",
})
fmt.Println(string(csvBytes))
```

## Endpoint Coverage

### Time Series (equities): 

- Intraday
- Daily
- Daily Adjusted
- Weekly
- Weekly Adjusted
- Monthly
- Monthly Adjusted
- Global Quote

### Cryptocurrencies: 

- Intraday
- Daily
- Weekly
- Monthly
- Exchange Rates

### Technical Indicators: 

- SMA/EMA/WMA/DEMA/TEMA/TRIMA/KAMA/MAMA/VWAP/T3
- MACD/MACDEXT
- TOCH/STOCHF
- RSI/STOCHRSI/WILLR
- ADX/ADXR
- AROON/AROONOSC
- BBANDS
- AD/ADOSC
- OBV
- CCI/CMO
- MIDPOINT/MIDPRICE
- SAR
- TRANGE/ATR/NATR
- ROC/ROCR
- MOM/BOP/APO/PPO
- MFI/TRIX/ULTOSC
- DX/MINUS_DI/PLUS_DI/MINUS_DM/PLUS_DM
- HT_* family

### Alpha Intelligence

- News & Sentiments `[planned]`
- Earnings Call Transcript `[planned]`
- Top Gainers & Losers `[planned]`
- Insider Transactions `[planned]`
- Analytics (Fixed Window)
- Analytics (Sliding Window)

### Fundamental Data

- Company Overview
- ETF Profile & Holdings
- Corporate Action - Dividends
- Corporate Action - Splits
- Income Statement
- Balance Sheet
- Cash Flow
- Shares Outstanding `[planned]`
- Earnings History `[planned]`
- Earnings Estimates `[planned]`
- Listing & Delisting Status `[planned]`
- Earnings Calendar `[planned]`
- IPO Calendar `[planned]`

Each endpoint has a dedicated params struct in `models` and a matching `Client` method. Planned items will follow the same pattern when added.

## Output Format

- All response types implement `String()` for terminal-friendly tables with deterministic ordering (maps normalized to slices).
- For structured access, use the fields on the returned structs (e.g., `TimeSeriesDaily.Metadata`, `TimeSeriesDaily.Data`).
- Errors from Alpha Vantage (notes, rate limits, premium notices) are returned as Go errors.

## Development

- Keep the library dependency-free.
- Run `go test ./...` before submitting changes.
- Format and vet touched files: `gofmt -w` and `go vet ./...`.

## License

MIT License. See `LICENSE` for details.

## Contact

- Issues & features: GitHub issues on this repo.
- Maintainer: masonJamesWheeler (GitHub).
