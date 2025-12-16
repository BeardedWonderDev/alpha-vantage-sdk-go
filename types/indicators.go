/*
// Package models provides types and functions for working with Alpha Vantage indicators data.
//
// This file contains types and functions representing the interactions and responses
// for technical indicators provided by the Alpha Vantage API.
// For more information about Alpha Vantage API, see https://www.alphavantage.co/documentation/.

Author: Mason Wheeler
*/

package types

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type IndicatorParams struct {
	Function   string
	Symbol     string
	Interval   string
	TimePeriod int
	SeriesType string
	Month      string
	OutputSize string
}

type IndicatorResponse struct {
	MetaData        TimeSeriesMetaData `json:"Meta Data"`
	IndicatorValues []IndicatorValue   `json:"-"`
}

type IndicatorValue struct {
	Timestamp time.Time          `json:"-"`
	Values    map[string]float64 `json:"-"`
}

func UnmarshalIndicatorJSON(i *IndicatorResponse, data []byte, indicatorName string) error {

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Metadata extraction
	metaData, ok := raw["Meta Data"].(map[string]interface{})
	if ok {
		i.MetaData = extractMetaData(metaData)
	}

	// Construct the expected key name
	expectedKey := "Technical Analysis: " + indicatorName

	// Extracting the indicator values
	if tsData, exists := raw[expectedKey].(map[string]interface{}); exists {
		for k, v := range tsData {
			timestamp, err := time.Parse("2006-01-02 15:04", k)
			if err != nil {
				return err
			}

			indicatorData, ok := v.(map[string]interface{})
			if !ok {
				return fmt.Errorf("expected map for each timestamp data")
			}

			// New changes to extract multiple values
			valueMap := make(map[string]float64)

			for name, rawValue := range indicatorData {
				switch value := rawValue.(type) {
				case string:
					value = strings.TrimSpace(value)
					if value == "" || isNAString(value) {
						valueMap[name] = 0
						continue
					}
					floatValue, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return err
					}
					valueMap[name] = floatValue
				case float64:
					valueMap[name] = value
				case json.Number:
					floatValue, err := value.Float64()
					if err != nil {
						return err
					}
					valueMap[name] = floatValue
				}
			}

			i.IndicatorValues = append(i.IndicatorValues, IndicatorValue{
				Timestamp: timestamp,
				Values:    valueMap,
			})
		}
	}

	// Sorting based on timestamps
	sort.SliceStable(i.IndicatorValues, func(a, b int) bool {
		return i.IndicatorValues[a].Timestamp.Before(i.IndicatorValues[b].Timestamp)
	})

	return nil
}

func extractMetaData(rawData map[string]interface{}) TimeSeriesMetaData {
	var metaData TimeSeriesMetaData

	for key, value := range rawData {
		switch key {
		case "1: Symbol":
			if v, ok := value.(string); ok {
				metaData.Symbol = v
			}
		case "2: Indicator":
			if v, ok := value.(string); ok {
				metaData.Information = v
			}
		case "3: Last Refreshed":
			if v, ok := value.(string); ok {
				metaData.LastRefreshed = v
			}
		case "4: Interval":
			if v, ok := value.(string); ok {
				metaData.Interval = v
			}
		case "5: Time Period":
			switch v := value.(type) {
			case float64:
				metaData.TimePeriod = v
			case json.Number:
				if f, err := v.Float64(); err == nil {
					metaData.TimePeriod = f
				}
			case string:
				v = strings.TrimSpace(v)
				if v != "" && !isNAString(v) {
					if f, err := strconv.ParseFloat(v, 64); err == nil {
						metaData.TimePeriod = f
					}
				}
			}
		case "6: Series Type":
			if v, ok := value.(string); ok {
				metaData.SeriesType = v
			}
		case "7: Time Zone":
			if v, ok := value.(string); ok {
				metaData.TimeZone = v
			}
		}
	}
	return metaData
}

func (i IndicatorResponse) String() string {
	var sb strings.Builder

	// Print metadata
	sb.WriteString(i.MetaData.Information + "\n")
	sb.WriteString(fmt.Sprintf("Symbol: %s\n", i.MetaData.Symbol))
	sb.WriteString(fmt.Sprintf("Last Refreshed: %s\n", i.MetaData.LastRefreshed))
	sb.WriteString(fmt.Sprintf("Interval: %s\n", i.MetaData.Interval))
	sb.WriteString(fmt.Sprintf("Output Size: %s\n", i.MetaData.OutputSize))
	sb.WriteString(fmt.Sprintf("Time Zone: %s\n", i.MetaData.TimeZone))
	sb.WriteString("\n")

	// Define headers dynamically
	headers := []string{"Time"}
	if len(i.IndicatorValues) > 0 {
		for k := range i.IndicatorValues[0].Values {
			headers = append(headers, k)
		}
	}

	// Print headers
	sb.WriteString(fmt.Sprintf("%-24s", headers[0])) // Set width for "Time"
	for _, header := range headers[1:] {
		sb.WriteString(fmt.Sprintf("%-15s", header))
	}
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("%-24s", strings.Repeat("=", 24))) // Set width for "Time"
	sb.WriteString(strings.Repeat("=", 15*(len(headers)-1)))
	sb.WriteString("\n")

	// Loop through the Indicator slice
	for _, v := range i.IndicatorValues {
		timeStr := v.Timestamp.Format("2006-01-02 15:04:05")
		sb.WriteString(fmt.Sprintf("%-24s", timeStr)) // Set width for "Time"
		for _, header := range headers[1:] {          // Skip "Time"
			sb.WriteString(fmt.Sprintf("%15.2f", v.Values[header]))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
