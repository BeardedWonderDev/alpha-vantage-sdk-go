package corestocks

import (
	"fmt"
	"net/url"
	"strings"

	itypes "github.com/BeardedWonderDev/alpha-vantage-sdk-go/internal/types"
	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

type CoreStucksService struct {
	client itypes.Client
}

func NewCoreStocksService(client itypes.Client) *CoreStucksService {
	return &CoreStucksService{client: client}
}

// getTimeSeriesData retrieves time series data based on the provided parameters.
func (c *CoreStucksService) getTimeSeriesData(function string, params types.TimeSeriesParams) ([]byte, error) {
	symbol := strings.TrimSpace(params.Symbol)
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbol", symbol)

	interval := strings.TrimSpace(params.Interval)
	if interval != "" {
		queryParams.Add("interval", interval)
	}

	if monthStr, ok := params.Month.(string); ok {
		monthStr = strings.TrimSpace(monthStr)
		if monthStr != "" {
			queryParams.Add("month", monthStr)
		}
	} else if monthPtr, ok := params.Month.(*string); ok {
		if monthPtr != nil {
			month := strings.TrimSpace(*monthPtr)
			if month != "" {
				queryParams.Add("month", month)
			}
		}
	}

	if outputStr, ok := params.OutputSize.(string); ok {
		outputStr = strings.TrimSpace(outputStr)
		if outputStr != "" {
			queryParams.Add("outputsize", outputStr)
		}
	} else if outputPtr, ok := params.OutputSize.(*string); ok {
		if outputPtr != nil {
			outputSize := strings.TrimSpace(*outputPtr)
			if outputSize != "" {
				queryParams.Add("outputsize", outputSize)
			}
		}
	}

	return c.client.Do(function, queryParams)
}
