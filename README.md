<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a id="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
I'm using markdown "reference style" links for readability.
Reference links are enclosed in brackets [ ] instead of parentheses ( ).
See the bottom of this document for the declaration of the reference variables.
-->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT HEADER -->
<br />
<div align="center">
  <h3 align="center">alpha-vantage-sdk-go</h3>

  <p align="center">
    A lightweight, dependency-free Go SDK for the Alpha Vantage REST API with typed requests/responses and domain-oriented services.
    <br />
    <a href="https://github.com/BeardedWonderDev/alpha-vantage-sdk-go"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/BeardedWonderDev/alpha-vantage-sdk-go/issues/new?labels=bug">Report Bug</a>
    &middot;
    <a href="https://github.com/BeardedWonderDev/alpha-vantage-sdk-go/issues/new?labels=enhancement">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

`alpha-vantage-sdk-go` is a Go SDK for the [Alpha Vantage](https://www.alphavantage.co/) REST API. The package is organized by domain services (e.g. `CoreStocks`, `Crypto`, `Forex`) and returns strongly typed response structs.

Key characteristics:

- Dependency-free (standard library only)
- JSON-only (the client enforces `datatype=json`)
- Domain-oriented services from a single client
- Alpha Vantage informational/error payloads are surfaced as Go errors

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

- [Go](https://go.dev/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

- Go 1.21+
- An Alpha Vantage API key (get one at https://www.alphavantage.co/support/#api-key)

### Installation

```sh
go get github.com/BeardedWonderDev/alpha-vantage-sdk-go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

Create a client and call domain services:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/av"
	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

func main() {
	cli := av.NewClient(os.Getenv("ALPHAVANTAGE_API_KEY"))

	// Core Stocks: Intraday time series
	intraday, err := cli.CoreStocks().Intraday(types.TimeSeriesParams{
		Symbol:     "MSFT",
		Interval:   "5min",
		OutputSize: "compact",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(intraday)

	// Core Stocks: Quote
	quote, err := cli.CoreStocks().Quote("MSFT")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(quote)

	// Technical Indicators: BBANDS
	bbands, err := cli.TechnicalIndicators().BBANDS(types.IndicatorParams{
		Symbol:     "MSFT",
		Interval:   "15min",
		TimePeriod: 20,
		SeriesType: "close",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bbands)
}
```

### Custom HTTP Client

```go
httpClient := &http.Client{Timeout: 10 * time.Second}
cli := av.NewClientWithHTTPClient(apiKey, httpClient)
```

### Additional Examples

```go
// Symbol Search
search, err := cli.CoreStocks().SymbolSearch("microsoft")

// Crypto: Daily series
cryptoDaily, err := cli.Crypto().Daily(types.CryptoDailyParams{Symbol: "BTC", Market: "USD"})

// Forex: Exchange rate
fx, err := cli.Forex().ExchangeRate(types.ForexExchangeRateParams{FromCurrency: "USD", ToCurrency: "EUR"})

// Fundamental Data
overview, err := cli.FundamentalData().CompanyOverview("IBM")

// Alpha Inteligence: Analytics
fixed, err := cli.AlphaInteligence().AnalyticsFixedWindow(types.AnalyticsFixedWindowParams{
	Symbols:      "IBM,AAPL,MSFT",
	Range:        []string{"2023-07-03", "2023-08-31"},
	Interval:     "DAILY",
	Calculations: "MEAN,STDDEV,CORRELATION",
})
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->
## Roadmap

- Expand endpoint coverage across all Alpha Vantage categories
- Add more fixtures and tests for service-level behavior

See the [open issues][issues-url] for a list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are welcome. Please open an issue to discuss the change, then submit a PR.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Project Link: https://github.com/BeardedWonderDev/alpha-vantage-sdk-go

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

- [Alpha Vantage](https://www.alphavantage.co/documentation/)
- [Best-README-Template](https://github.com/othneildrew/Best-README-Template)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/BeardedWonderDev/alpha-vantage-sdk-go.svg?style=for-the-badge
[contributors-url]: https://github.com/BeardedWonderDev/alpha-vantage-sdk-go/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/BeardedWonderDev/alpha-vantage-sdk-go.svg?style=for-the-badge
[forks-url]: https://github.com/BeardedWonderDev/alpha-vantage-sdk-go/network/members
[stars-shield]: https://img.shields.io/github/stars/BeardedWonderDev/alpha-vantage-sdk-go.svg?style=for-the-badge
[stars-url]: https://github.com/BeardedWonderDev/alpha-vantage-sdk-go/stargazers
[issues-shield]: https://img.shields.io/github/issues/BeardedWonderDev/alpha-vantage-sdk-go.svg?style=for-the-badge
[issues-url]: https://github.com/BeardedWonderDev/alpha-vantage-sdk-go/issues
[license-shield]: https://img.shields.io/github/license/BeardedWonderDev/alpha-vantage-sdk-go.svg?style=for-the-badge
[license-url]: https://github.com/BeardedWonderDev/alpha-vantage-sdk-go/blob/main/LICENSE
