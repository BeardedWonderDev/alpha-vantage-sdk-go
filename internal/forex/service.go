package forex

import itypes "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/types"

type ForexService struct {
	client itypes.Client
}

func NewForexService(client itypes.Client) *ForexService {
	return &ForexService{client: client}
}

