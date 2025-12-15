package av

import (
	"net/http"

	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

func NewClient(apiKey string) types.Client {
	return internal.NewClient(apiKey, nil)
}

func NewClientWithHTTPClient(apiKey string, httpClient *http.Client) types.Client {
	return internal.NewClient(apiKey, httpClient)
}
