package corestocks

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

// Intraday retrieves intraday data based on the provided parameters.
// It returns a TimeSeriesIntraday and an error if there is any.
func (c *CoreStucksService) Intraday(params types.TimeSeriesParams) (types.TimeSeriesIntraday, error) {
	if strings.TrimSpace(params.Interval) == "" {
		return types.TimeSeriesIntraday{}, fmt.Errorf("interval is required")
	}

	data, err := c.getTimeSeriesData("TIME_SERIES_INTRADAY", params)
	if err != nil {
		return types.TimeSeriesIntraday{}, err
	}

	var intradayData types.TimeSeriesIntraday
	err = json.Unmarshal(data, &intradayData)
	if err != nil {
		return types.TimeSeriesIntraday{}, err
	}

	return intradayData, nil
}

// Daily retrieves daily data based on the provided parameters.
// It returns a TimeSeriesDaily and an error if there is any.
func (c *CoreStucksService) Daily(params types.TimeSeriesParams) (types.TimeSeriesDaily, error) {
	data, err := c.getTimeSeriesData("TIME_SERIES_DAILY", params)
	if err != nil {
		return types.TimeSeriesDaily{}, err
	}

	var dailyData types.TimeSeriesDaily
	err = json.Unmarshal(data, &dailyData)
	if err != nil {
		return types.TimeSeriesDaily{}, err
	}

	return dailyData, nil
}

// DailyAdjusted retrieves daily adjusted data based on the provided parameters.
// It returns a TimeSeriesDailyAdjusted and an error if there is any.
func (c *CoreStucksService) DailyAdjusted(params types.TimeSeriesParams) (types.TimeSeriesDailyAdjusted, error) {
	data, err := c.getTimeSeriesData("TIME_SERIES_DAILY_ADJUSTED", params)
	if err != nil {
		return types.TimeSeriesDailyAdjusted{}, err
	}

	var dailyAdjustedData types.TimeSeriesDailyAdjusted
	err = json.Unmarshal(data, &dailyAdjustedData)
	if err != nil {
		return types.TimeSeriesDailyAdjusted{}, err
	}
	return dailyAdjustedData, nil
}

// Weekly retrieves weekly data based on the provided parameters.
// It returns a TimeSeriesWeekly and an error if there is any.
func (c *CoreStucksService) Weekly(params types.TimeSeriesParams) (types.TimeSeriesWeekly, error) {
	data, err := c.getTimeSeriesData("TIME_SERIES_WEEKLY", params)
	if err != nil {
		return types.TimeSeriesWeekly{}, err
	}

	var weeklyData types.TimeSeriesWeekly
	err = json.Unmarshal(data, &weeklyData)
	if err != nil {
		return types.TimeSeriesWeekly{}, err
	}
	return weeklyData, nil
}

// WeeklyAdjusted retrieves weekly adjusted data based on the provided parameters.
// It returns a TimeSeriesWeekly and an error if there is any.
func (c *CoreStucksService) WeeklyAdjusted(params types.TimeSeriesParams) (types.TimeSeriesWeekly, error) {
	data, err := c.getTimeSeriesData("TIME_SERIES_WEEKLY_ADJUSTED", params)
	if err != nil {
		return types.TimeSeriesWeekly{}, err
	}

	var weeklyAdjustedData types.TimeSeriesWeekly
	err = json.Unmarshal(data, &weeklyAdjustedData)
	if err != nil {
		return types.TimeSeriesWeekly{}, err
	}
	return weeklyAdjustedData, nil
}

// Monthly retrieves monthly data based on the provided parameters.
// It returns a TimeSeriesMonthly and an error if there is any.
func (c *CoreStucksService) Monthly(params types.TimeSeriesParams) (types.TimeSeriesMonthly, error) {
	data, err := c.getTimeSeriesData("TIME_SERIES_MONTHLY", params)
	if err != nil {
		return types.TimeSeriesMonthly{}, err
	}

	var monthlyData types.TimeSeriesMonthly
	err = json.Unmarshal(data, &monthlyData)
	if err != nil {
		return types.TimeSeriesMonthly{}, err
	}
	return monthlyData, nil
}

// MonthlyAdjusted retrieves monthly adjusted data based on the provided parameters.
// It returns a TimeSeriesMonthlyAdjusted and an error if there is any.
func (c *CoreStucksService) MonthlyAdjusted(params types.TimeSeriesParams) (types.TimeSeriesMonthlyAdjusted, error) {
	data, err := c.getTimeSeriesData("TIME_SERIES_MONTHLY_ADJUSTED", params)
	if err != nil {
		return types.TimeSeriesMonthlyAdjusted{}, err
	}

	var monthlyAdjustedData types.TimeSeriesMonthlyAdjusted
	err = json.Unmarshal(data, &monthlyAdjustedData)
	if err != nil {
		return types.TimeSeriesMonthlyAdjusted{}, err
	}
	return monthlyAdjustedData, nil
}
