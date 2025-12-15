package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CryptoExchangeRateParams struct {
	FromCurrency string
	ToCurrency   string
}

type CryptoIntradayParams struct {
	Symbol     string
	Market     string
	Interval   string
	OutputSize string
}

type CryptoDailyParams struct {
	Symbol string
	Market string
}

type CryptoWeeklyParams struct {
	Symbol string
	Market string
}

type CryptoMonthlyParams struct {
	Symbol string
	Market string
}

type CryptoSeriesResponse struct {
	MetaData      CryptoMetaData
	TimeSeries    []CryptoTimeSeriesData
	IntervalLabel string
}

type CryptoMetaData struct {
	Information         string
	DigitalCurrencyCode string
	DigitalCurrencyName string
	MarketCode          string
	MarketName          string
	LastRefreshed       string
	TimeZone            string
}

type CryptoTimeSeriesData struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
	MarketCap float64
}

func UnmarshalCryptoJSON(c *CryptoSeriesResponse, data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Metadata extraction
	metaData, ok := raw["Meta Data"].(map[string]interface{})
	if ok {
		c.MetaData = extractCryptoMetaData(metaData)
	}

	for tsKey, tsData := range raw {
		if strings.HasPrefix(tsKey, "Time Series") {
			c.IntervalLabel = tsKey
			timeSeriesMap, ok := tsData.(map[string]interface{})
			if !ok {
				return fmt.Errorf("expected map for time series data")
			}

			for date, values := range timeSeriesMap {
				timestamp, err := time.Parse("2006-01-02", date)
				if err != nil {
					return err
				}

				valuesMap, ok := values.(map[string]interface{})
				if !ok {
					return fmt.Errorf("expected map for timestamp data")
				}

				open, _ := strconv.ParseFloat(asString(valuesMap["1a. open (USD)"]), 64)
				high, _ := strconv.ParseFloat(asString(valuesMap["2a. high (USD)"]), 64)
				low, _ := strconv.ParseFloat(asString(valuesMap["3a. low (USD)"]), 64)
				closeVal, _ := strconv.ParseFloat(asString(valuesMap["4a. close (USD)"]), 64)
				volume, _ := strconv.ParseFloat(asString(valuesMap["5. volume"]), 64)
				marketCap, _ := strconv.ParseFloat(asString(valuesMap["6. market cap (USD)"]), 64)

				c.TimeSeries = append(c.TimeSeries, CryptoTimeSeriesData{
					Timestamp: timestamp,
					Open:      open,
					High:      high,
					Low:       low,
					Close:     closeVal,
					Volume:    volume,
					MarketCap: marketCap,
				})
			}
		}
	}

	// Sorting based on timestamps
	sort.SliceStable(c.TimeSeries, func(a, b int) bool {
		return c.TimeSeries[a].Timestamp.Before(c.TimeSeries[b].Timestamp)
	})

	return nil
}

func extractCryptoMetaData(rawData map[string]interface{}) CryptoMetaData {
	var metaData CryptoMetaData

	for key, value := range rawData {
		switch key {
		case "1. Information":
			metaData.Information = asString(value)
		case "2. Digital Currency Code":
			metaData.DigitalCurrencyCode = asString(value)
		case "3. Digital Currency Name":
			metaData.DigitalCurrencyName = asString(value)
		case "4. Market Code":
			metaData.MarketCode = asString(value)
		case "5. Market Name":
			metaData.MarketName = asString(value)
		case "6. Last Refreshed":
			metaData.LastRefreshed = asString(value)
		case "7. Time Zone":
			metaData.TimeZone = asString(value)
		}
	}
	return metaData
}

func asString(v interface{}) string {
	if v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}

