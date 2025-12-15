package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/models"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestGetSymbolSearch_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture, err := os.ReadFile("../../models/testdata/symbol_search_sample.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.URL.Host != "www.alphavantage.co" || req.URL.Path != "/query" {
			return nil, fmt.Errorf("unexpected request url: %s", req.URL.String())
		}

		q := req.URL.Query()
		if q.Get("function") != "SYMBOL_SEARCH" {
			return nil, fmt.Errorf("expected function SYMBOL_SEARCH, got %q", q.Get("function"))
		}
		if q.Get("keywords") != "microsoft" {
			return nil, fmt.Errorf("expected keywords microsoft, got %q", q.Get("keywords"))
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
	resp, err := c.GetSymbolSearch(models.SymbolSearchParams{Keywords: "microsoft"})
	if err != nil {
		t.Fatalf("GetSymbolSearch returned error: %v", err)
	}

	if len(resp.BestMatches) != 10 {
		t.Fatalf("expected 10 matches, got %d", len(resp.BestMatches))
	}
	if resp.BestMatches[0].Symbol != "BA" {
		t.Fatalf("expected first symbol BA, got %s", resp.BestMatches[0].Symbol)
	}
}

func TestGetSymbolSearch_RejectsCSV(t *testing.T) {
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("unexpected network call: %s", req.URL.String())
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	c := NewClient("test-key")
	_, err := c.GetSymbolSearch(models.SymbolSearchParams{
		Keywords: "microsoft",
		DataType: "csv",
	})
	if err == nil {
		t.Fatalf("expected error for csv datatype, got nil")
	}
}

func TestGetSymbolSearchData_AllowsCSV(t *testing.T) {
	csvBody := []byte("symbol,name\nBA,Boeing Company\n")

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
	data, err := c.GetSymbolSearchData(models.SymbolSearchParams{
		Keywords: "microsoft",
		DataType: "CSV",
	})
	if err != nil {
		t.Fatalf("GetSymbolSearchData returned error: %v", err)
	}

	if !bytes.Equal(data, csvBody) {
		t.Fatalf("unexpected csv body: %q", string(data))
	}
}

func TestGetSymbolSearchData_ValidatesInputs(t *testing.T) {
	c := NewClient("test-key")

	if _, err := c.GetSymbolSearchData(models.SymbolSearchParams{Keywords: ""}); err == nil {
		t.Fatalf("expected error for empty keywords, got nil")
	}

	if _, err := c.GetSymbolSearchData(models.SymbolSearchParams{Keywords: "microsoft", DataType: "xml"}); err == nil {
		t.Fatalf("expected error for invalid datatype, got nil")
	}
}
