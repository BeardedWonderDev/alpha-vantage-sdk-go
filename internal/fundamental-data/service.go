package fundamentaldata

import itypes "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/types"

type FundamentalDataService struct {
	client itypes.Client
}

func NewFundamentalDataService(client itypes.Client) *FundamentalDataService {
	return &FundamentalDataService{
		client: client,
	}
}
