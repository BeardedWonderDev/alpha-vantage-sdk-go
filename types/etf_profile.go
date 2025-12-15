package types

import (
	"fmt"
	"strings"
)

// ETFProfileParams defines the parameters for the ETF profile endpoint.
type ETFProfileParams struct {
	Symbol string
}

// ETFProfile represents the response for the ETF profile & holdings endpoint.
type ETFProfile struct {
	NetAssets         int64               `json:"net_assets,string"`
	NetExpenseRatio   float64             `json:"net_expense_ratio,string"`
	PortfolioTurnover string              `json:"portfolio_turnover"`
	DividendYield     float64             `json:"dividend_yield,string"`
	InceptionDate     string              `json:"inception_date"`
	Leveraged         string              `json:"leveraged"`
	Sectors           []ETFProfileSector  `json:"sectors"`
	Holdings          []ETFProfileHolding `json:"holdings"`
}

// ETFProfileSector represents sector allocation entries.
type ETFProfileSector struct {
	Sector string  `json:"sector"`
	Weight float64 `json:"weight,string"`
}

// ETFProfileHolding represents a single holding entry.
type ETFProfileHolding struct {
	Symbol      string  `json:"symbol"`
	Description string  `json:"description"`
	Weight      float64 `json:"weight,string"`
}

// String renders a concise human-readable summary of the ETF profile.
func (p ETFProfile) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("ETF Profile\nNet Assets: %d\nExpense Ratio: %.4f\nDividend Yield: %.4f\nInception Date: %s\nLeveraged: %s\n", p.NetAssets, p.NetExpenseRatio, p.DividendYield, p.InceptionDate, p.Leveraged))

	// Sectors
	if len(p.Sectors) > 0 {
		sb.WriteString("\nSectors (weight):\n")
		for _, s := range p.Sectors {
			sb.WriteString(fmt.Sprintf(" - %s: %.3f\n", s.Sector, s.Weight))
		}
	}

	// Holdings (top 10 shown if longer)
	if len(p.Holdings) > 0 {
		sb.WriteString("\nHoldings (weight):\n")
		limit := len(p.Holdings)
		if limit > 10 {
			limit = 10
		}
		for i := 0; i < limit; i++ {
			h := p.Holdings[i]
			sb.WriteString(fmt.Sprintf(" - %s (%s): %.4f\n", h.Symbol, h.Description, h.Weight))
		}
		if len(p.Holdings) > limit {
			sb.WriteString(fmt.Sprintf(" ...and %d more\n", len(p.Holdings)-limit))
		}
	}

	return sb.String()
}
