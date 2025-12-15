package client

import "testing"

func TestDetectAPIMessageInformation(t *testing.T) {
	sample := []byte(`{"Information": "Thank you for using Alpha Vantage! This is a premium endpoint. You may subscribe to any of the premium plans at https://www.alphavantage.co/premium/ to instantly unlock all premium endpoints"}`)

	if err := detectAPIMessage(sample); err == nil {
		t.Fatalf("expected detectAPIMessage to return an error for Information payload")
	}
}

func TestDetectAPIMessagePassThrough(t *testing.T) {
	sample := []byte(`{"foo": "bar"}`)
	if err := detectAPIMessage(sample); err != nil {
		t.Fatalf("expected no error for non-message payload, got %v", err)
	}
}
