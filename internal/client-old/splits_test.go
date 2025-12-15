package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/models"
)

func TestGetSplits_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture, err := os.ReadFile("../models/testdata/splits_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.URL.Host != "www.alphavantage.co" || req.URL.Path != "/query" {
			return nil, fmt.Errorf("unexpected request url: %s", req.URL.String())
		}

		q := req.URL.Query()
		if q.Get("function") != "SPLITS" {
			return nil, fmt.Errorf("expected function SPLITS, got %q", q.Get("function"))
		}
		if q.Get("symbol") != "IBM" {
			return nil, fmt.Errorf("expected symbol IBM, got %q", q.Get("symbol"))
		}
		if q.Get("apikey") != "test-key" {
			return nil, fmt.Errorf("expected apikey test-key, got %q", q.Get("apikey"))
		}
		if q.Get("datatype") != "" {
			return nil, fmt.Errorf("expected no datatype param, got %q", q.Get("datatype"))
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(fixture)),
			Header:     make(http.Header),
			Request:    req,
		}, nil
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	c := NewClient("test-key")
	resp, err := c.GetSplits(models.SplitsParams{Symbol: "IBM"})
	if err != nil {
		t.Fatalf("GetSplits returned error: %v", err)
	}

	if resp.Symbol != "IBM" {
		t.Fatalf("expected IBM, got %s", resp.Symbol)
	}
	if len(resp.Data) != 2 {
		t.Fatalf("expected 2 split records, got %d", len(resp.Data))
	}
}

func TestGetSplits_RejectsCSV(t *testing.T) {
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("unexpected network call: %s", req.URL.String())
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	c := NewClient("test-key")
	_, err := c.GetSplits(models.SplitsParams{
		Symbol:   "IBM",
		DataType: "csv",
	})
	if err == nil {
		t.Fatalf("expected error for csv datatype, got nil")
	}
}

func TestGetSplitsData_AllowsCSV(t *testing.T) {
	csvBody := []byte("effective_date,split_factor\n2021-11-04,1.0460\n")

	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		q := req.URL.Query()
		if q.Get("datatype") != "csv" {
			return nil, fmt.Errorf("expected datatype csv, got %q", q.Get("datatype"))
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(csvBody)),
			Header:     make(http.Header),
			Request:    req,
		}, nil
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	c := NewClient("test-key")
	data, err := c.GetSplitsData(models.SplitsParams{
		Symbol:   "IBM",
		DataType: "CSV",
	})
	if err != nil {
		t.Fatalf("GetSplitsData returned error: %v", err)
	}

	if !bytes.Equal(data, csvBody) {
		t.Fatalf("unexpected csv body: %q", string(data))
	}
}

func TestGetSplitsData_ValidatesInputs(t *testing.T) {
	c := NewClient("test-key")

	if _, err := c.GetSplitsData(models.SplitsParams{Symbol: ""}); err == nil {
		t.Fatalf("expected error for empty symbol, got nil")
	}

	if _, err := c.GetSplitsData(models.SplitsParams{Symbol: "IBM", DataType: "xml"}); err == nil {
		t.Fatalf("expected error for invalid datatype, got nil")
	}
}
