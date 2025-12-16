package alphainteligence

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

func (c *AlphaInteligenceService) AnalyticsSlidingWindow(params types.AnalyticsSlidingWindowParams) (*types.AnalyticsSlidingWindowResponse, error) {
	symbols := strings.TrimSpace(params.Symbols)
	interval := strings.TrimSpace(params.Interval)
	calculations := strings.TrimSpace(params.Calculations)
	ohlc := strings.TrimSpace(params.Ohlc)

	if symbols == "" {
		return nil, fmt.Errorf("symbols are required")
	}
	if interval == "" {
		return nil, fmt.Errorf("interval is required")
	}
	if params.WindowSize == 0 {
		return nil, fmt.Errorf("window size is required")
	}
	if calculations == "" {
		return nil, fmt.Errorf("calculations are required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbols", symbols)
	queryParams.Add("interval", interval)
	queryParams.Add("window_size", fmt.Sprintf("%d", params.WindowSize))
	queryParams.Add("calculations", calculations)

	for _, r := range params.Range {
		r = strings.TrimSpace(r)
		if r != "" {
			queryParams.Add("range", r)
		}
	}

	if ohlc != "" {
		queryParams.Add("ohlc", ohlc)
	}

	data, err := c.client.Do("ANALYTICS_SLIDING_WINDOW", queryParams)
	if err != nil {
		return nil, err
	}

	var resp types.AnalyticsSlidingWindowResponse
	if err := types.UnmarshalLenient(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *AlphaInteligenceService) AnalyticsFixedWindow(params types.AnalyticsFixedWindowParams) (*types.AnalyticsFixedWindowResponse, error) {
	symbols := strings.TrimSpace(params.Symbols)
	interval := strings.TrimSpace(params.Interval)
	calculations := strings.TrimSpace(params.Calculations)
	ohlc := strings.TrimSpace(params.Ohlc)

	if symbols == "" {
		return nil, fmt.Errorf("symbols are required")
	}
	if interval == "" {
		return nil, fmt.Errorf("interval is required")
	}
	if calculations == "" {
		return nil, fmt.Errorf("calculations are required")
	}

	queryParams := url.Values{}
	queryParams.Add("symbols", symbols)
	queryParams.Add("interval", interval)
	queryParams.Add("calculations", calculations)

	for _, r := range params.Range {
		r = strings.TrimSpace(r)
		if r != "" {
			queryParams.Add("range", r)
		}
	}

	if ohlc != "" {
		queryParams.Add("ohlc", ohlc)
	}

	data, err := c.client.Do("ANALYTICS_FIXED_WINDOW", queryParams)
	if err != nil {
		return nil, err
	}

	var resp types.AnalyticsFixedWindowResponse
	if err := types.UnmarshalLenient(data, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
