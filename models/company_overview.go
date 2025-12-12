package models

import (
	"fmt"
	"strings"
)

// CompanyOverviewParams defines the required parameters for the Alpha Vantage
// company overview (function=OVERVIEW) endpoint.
type CompanyOverviewParams struct {
	Symbol   string
	DataType string
}

// CompanyOverview models the response returned by the Alpha Vantage Company
// Overview endpoint. Alpha Vantage returns most numeric values as strings, so
// numeric-looking fields are parsed into typed numbers using the json ",string"
// tag pattern used elsewhere in the models package.
type CompanyOverview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	CIK                        string  `json:"CIK"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	OfficialSite               string  `json:"OfficialSite"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"`
	MarketCapitalization       int64   `json:"MarketCapitalization,string"`
	EBITDA                     int64   `json:"EBITDA,string"`
	PERatio                    float64 `json:"PERatio,string"`
	PEGRatio                   float64 `json:"PEGRatio,string"`
	BookValue                  float64 `json:"BookValue,string"`
	DividendPerShare           float64 `json:"DividendPerShare,string"`
	DividendYield              float64 `json:"DividendYield,string"`
	EPS                        float64 `json:"EPS,string"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 int64   `json:"RevenueTTM,string"`
	GrossProfitTTM             int64   `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         float64 `json:"AnalystTargetPrice,string"`
	AnalystRatingStrongBuy     int     `json:"AnalystRatingStrongBuy,string"`
	AnalystRatingBuy           int     `json:"AnalystRatingBuy,string"`
	AnalystRatingHold          int     `json:"AnalystRatingHold,string"`
	AnalystRatingSell          int     `json:"AnalystRatingSell,string"`
	AnalystRatingStrongSell    int     `json:"AnalystRatingStrongSell,string"`
	TrailingPE                 float64 `json:"TrailingPE,string"`
	ForwardPE                  float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                float64 `json:"EVToRevenue,string"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,string"`
	Beta                       float64 `json:"Beta,string"`
	Week52High                 float64 `json:"52WeekHigh,string"`
	Week52Low                  float64 `json:"52WeekLow,string"`
	MovingAverage50Day         float64 `json:"50DayMovingAverage,string"`
	MovingAverage200Day        float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int64   `json:"SharesOutstanding,string"`
	SharesFloat                int64   `json:"SharesFloat,string"`
	PercentInsiders            float64 `json:"PercentInsiders,string"`
	PercentInstitutions        float64 `json:"PercentInstitutions,string"`
	DividendDate               string  `json:"DividendDate"`
	ExDividendDate             string  `json:"ExDividendDate"`
}

// String returns a succinct, human-readable summary of key overview metrics.
func (o CompanyOverview) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s (%s)\n", o.Name, o.Symbol))
	sb.WriteString(fmt.Sprintf("Exchange: %s | Sector: %s | Industry: %s\n", o.Exchange, o.Sector, o.Industry))
	sb.WriteString(fmt.Sprintf("Market Cap: %d %s | EPS: %.2f | P/E: %.2f | PEG: %.2f\n", o.MarketCapitalization, o.Currency, o.EPS, o.PERatio, o.PEGRatio))
	sb.WriteString(fmt.Sprintf("Dividend/Share: %.2f (Yield: %.4f) | Latest Quarter: %s\n", o.DividendPerShare, o.DividendYield, o.LatestQuarter))
	sb.WriteString(fmt.Sprintf("52W High/Low: %.2f / %.2f | 50D MA: %.2f | 200D MA: %.2f\n", o.Week52High, o.Week52Low, o.MovingAverage50Day, o.MovingAverage200Day))

	return sb.String()
}
