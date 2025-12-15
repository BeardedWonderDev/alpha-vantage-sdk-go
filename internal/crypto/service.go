package crypto

import itypes "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/types"

type CryptoService struct {
	client itypes.Client
}

func NewCryptoService(client itypes.Client) *CryptoService {
	return &CryptoService{client: client}
}

