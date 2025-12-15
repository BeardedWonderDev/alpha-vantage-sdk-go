package av

import (
	"net/http"

	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/internal"
	"github.com/BeardedWonderDev/alpha-vantage-sdk-go/types"
)

func NewClient(apiKey string) types.Client {
	return internal.NewClient(apiKey, nil)
}

func NewClientWithHTTPClient(apiKey string, httpClient *http.Client) types.Client {
	return internal.NewClient(apiKey, httpClient)
}
