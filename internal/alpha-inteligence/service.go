package alphainteligence

import itypes "github.com/BeardedWonderDev/alpha-vantage-sdk-go/internal/types"

type AlphaInteligenceService struct {
	client itypes.Client
}

func NewAlphaInteligenceService(client itypes.Client) *AlphaInteligenceService {
	return &AlphaInteligenceService{client: client}
}
