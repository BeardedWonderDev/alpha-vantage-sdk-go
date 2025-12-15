package crypto

import itypes "github.com/BeardedWonderDev/alpha-vantage-sdk-go/internal/types"

type CryptoService struct {
	client itypes.Client
}

func NewCryptoService(client itypes.Client) *CryptoService {
	return &CryptoService{client: client}
}
