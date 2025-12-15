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

func TestGetCashFlow_SendsExpectedQueryAndParsesResponse(t *testing.T) {
	fixture, err := os.ReadFile("../models/testdata/cash_flow_IBM.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.URL.Host != "www.alphavantage.co" || req.URL.Path != "/query" {
			return nil, fmt.Errorf("unexpected request url: %s", req.URL.String())
		}

		q := req.URL.Query()
		if q.Get("function") != "CASH_FLOW" {
			return nil, fmt.Errorf("expected function CASH_FLOW, got %q", q.Get("function"))
		}
		if q.Get("symbol") != "IBM" {
			return nil, fmt.Errorf("expected symbol IBM, got %q", q.Get("symbol"))
		}
		if q.Get("apikey") != "test-key" {
			return nil, fmt.Errorf("expected apikey test-key, got %q", q.Get("apikey"))
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
	resp, err := c.GetCashFlow(models.CashFlowParams{Symbol: "IBM"})
	if err != nil {
		t.Fatalf("GetCashFlow returned error: %v", err)
	}

	if resp.Symbol != "IBM" {
		t.Fatalf("expected IBM, got %s", resp.Symbol)
	}
	if len(resp.AnnualReports) != 2 {
		t.Fatalf("expected 2 annual reports, got %d", len(resp.AnnualReports))
	}
	if len(resp.QuarterlyReports) != 2 {
		t.Fatalf("expected 2 quarterly reports, got %d", len(resp.QuarterlyReports))
	}
}

func TestGetCashFlow_ValidatesInputs(t *testing.T) {
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("unexpected network call: %s", req.URL.String())
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	c := NewClient("test-key")
	if _, err := c.GetCashFlow(models.CashFlowParams{Symbol: ""}); err == nil {
		t.Fatalf("expected error for empty symbol, got nil")
	}
}
