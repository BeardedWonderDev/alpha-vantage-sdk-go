package corestocks

import (
	"net/url"

	itypes "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/types"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

type CoreStucksService struct {
	client itypes.Client
}

func NewCoreStocksService(client itypes.Client) *CoreStucksService {
	return &CoreStucksService{client: client}
}

// getTimeSeriesData retrieves time series data based on the provided parameters.
func (c *CoreStucksService) getTimeSeriesData(function string, params types.TimeSeriesParams) ([]byte, error) {
	queryParams := url.Values{}
	queryParams.Add("symbol", params.Symbol)
	queryParams.Add("interval", params.Interval)

	if monthStr, ok := params.Month.(string); ok {
		queryParams.Add("month", monthStr)
	} else if monthPtr, ok := params.Month.(*string); ok {
		queryParams.Add("month", *monthPtr)
	}

	if outputStr, ok := params.OutputSize.(string); ok {
		queryParams.Add("outputsize", outputStr)
	} else if outputPtr, ok := params.OutputSize.(*string); ok {
		queryParams.Add("outputsize", *outputPtr)
	}

	return c.client.Do(function, queryParams)
}
