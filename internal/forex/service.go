package forex

import itypes "github.com/BeardedWonderDev/alpha-vantage-sdk-go/internal/types"

type ForexService struct {
	client itypes.Client
}

func NewForexService(client itypes.Client) *ForexService {
	return &ForexService{client: client}
}
