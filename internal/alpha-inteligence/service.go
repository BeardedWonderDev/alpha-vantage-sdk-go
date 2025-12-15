package alphainteligence

import itypes "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/types"

type AlphaInteligenceService struct {
	client itypes.Client
}

func NewAlphaInteligenceService(client itypes.Client) *AlphaInteligenceService {
	return &AlphaInteligenceService{client: client}
}

