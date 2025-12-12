package models

import (
	"fmt"
	"strings"
)

// DividendsParams defines parameters for the dividends endpoint.
type DividendsParams struct {
	Symbol   string
	DataType string
}

// DividendsResponse models the dividends API response.
type DividendsResponse struct {
	Symbol string           `json:"symbol"`
	Data   []DividendRecord `json:"data"`
}

// DividendRecord represents an individual dividend event.
type DividendRecord struct {
	ExDividendDate  string  `json:"ex_dividend_date"`
	DeclarationDate string  `json:"declaration_date"`
	RecordDate      string  `json:"record_date"`
	PaymentDate     string  `json:"payment_date"`
	Amount          float64 `json:"amount,string"`
}

// String renders a concise summary of the dividends schedule.
func (r DividendsResponse) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Dividends for %s\n", r.Symbol))

	limit := len(r.Data)
	if limit > 10 {
		limit = 10
	}

	for i := 0; i < limit; i++ {
		d := r.Data[i]
		sb.WriteString(fmt.Sprintf("%s: %.2f (pay %s)\n", d.ExDividendDate, d.Amount, d.PaymentDate))
	}

	if len(r.Data) > limit {
		sb.WriteString(fmt.Sprintf("...and %d more\n", len(r.Data)-limit))
	}

	return sb.String()
}
